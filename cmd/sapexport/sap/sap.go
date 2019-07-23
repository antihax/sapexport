package sap

import (
	"errors"
	"strconv"
	"strings"
	"os"
	"github.com/sap/gorfc/gorfc"
)

// RFC is an interface to a SAP system via RFC
type RFC struct {
	conn *gorfc.Connection
}

// Connect to the SAP ECC/S4 system
func (s *RFC) Connect(g gorfc.ConnectionParameter) error {
	if g.Passwd == "" {
		g.Passwd = os.Getenv("SAPRFC_PASS")
	}
	c, err := gorfc.ConnectionFromParams(g)
	s.conn = c
	return err
}

// Disconnect from the SAP ECC/S4 system
func (s *RFC) Disconnect() error {
	return s.conn.Close()
}

// UsersOfRole calls RFC ESS_USERS_OF_ROLE_GET to list users with a specific role
func (s *RFC) UsersOfRole(roleName string) (map[string]interface{}, error) {
	params := map[string]interface{}{
		"ROLE": roleName,
	}
	r, err := s.conn.Call("ESS_USERS_OF_ROLE_GET", params)
	if err != nil {
		return nil, err
	}

	return r, err
}

// ReadTable calls RFC /BODS/RFC_READ_TABLE2 with an optional OPTIONS text containing
// an ABAP where clause
func (s *RFC) ReadTable(tableName string, where string) ([]map[string]string, error) {
	options := make([]map[string]interface{}, 0)

	if where != "" {
		options = append(options, map[string]interface{}{
			"TEXT": where,
		})
	}

	params := map[string]interface{}{
		"QUERY_TABLE": tableName,
		"OPTIONS":     options,
	}
	r, err := s.conn.Call("/BODS/RFC_READ_TABLE2", params)
	if err != nil {
		return nil, err
	}

	fields, ok := r["FIELDS"].([]interface{})
	if !ok {
		return nil, errors.New("no fields")
	}

	for _, f := range fields {
		v := f.(map[string]interface{})
		l, _ := strconv.Atoi(v["LENGTH"].(string))
		o, _ := strconv.Atoi(v["OFFSET"].(string))
		v["len"] = l
		v["offset"] = o
	}
	rows := make([]map[string]string, 0)
	for k := range r {
		if !strings.Contains(k, "TBLOUT") {
			continue
		}
		data := r[k].([]interface{})

		for _, d := range data {
			result := make(map[string]string)
			for _, f := range fields {
				fv := f.(map[string]interface{})
				dv := d.((map[string]interface{}))

				fieldName, ok := fv["FIELDNAME"].(string)
				if !ok {
					return nil, errors.New("missing field name")
				}
				offset, ok := fv["offset"].(int)
				if !ok {
					return nil, errors.New("missing field name")
				}
				length, ok := fv["len"].(int)
				if !ok {
					return nil, errors.New("missing field name")
				}
				fieldData, ok := dv["WA"].(string)
				if !ok {
					return nil, errors.New("missing field data")
				}
				end := offset + length
				start := offset
				if len(fieldData) < start {
					start = len(fieldData)
				}
				if len(fieldData) < end {
					end = len(fieldData)
				}
				result[strings.TrimSpace(fieldName)] = strings.TrimSpace(fieldData[start:end])

			}
			rows = append(rows, result)
		}
	}

	return rows, nil
}
