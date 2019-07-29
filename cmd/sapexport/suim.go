package sapexport

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/sapexport/cmd/sapexport/sap"
	"github.com/spf13/cobra"
)

func init() {
	var file string
	/*
		    - SUSR_SUIM_API_RSUSR002: Users by complex selection criteria
		    - SUSR_SUIM_API_RSUSR008_009_NEW: Users with critical combinations of auhthorizations
		    - SUSR_SUIM_API_RSUSR020: Profiles by complex selection criteria
		    - SUSR_SUIM_API_RSUSR050_AUTH: Authorization Comparison
		    - SUSR_SUIM_API_RSUSR050_PROF: Profile Comparison
		    - SUSR_SUIM_API_RSUSR050_ROLE: Role Comparison
		    - SUSR_SUIM_API_RSUSR050_USER: User Comparison
		    - SUSR_SUIM_API_RSUSR070: Roles by complex search criteria
		    - SUSR_SUIM_API_RSUSR100N: Change documents for users
			- SUSR_SUIM_API_RSUSR200: Users according to logon date and password change
	*/

	var cmdChangeDocs = &cobra.Command{
		Use:   "changedocs [user]",
		Short: "Search changedocs",
		Long:  `changedocs will list change documents that match the parameters.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}
			var params map[string]interface{}
			if file != "" {
				params, err = readYaml(file)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				params = map[string]interface{}{
					"IT_USER": []map[string]interface{}{
						map[string]interface{}{
							"SIGN": "I",
							"LOW":  args[0],
						}},
					"IV_PASS": true,
					"IV_ROLE": true,
					"IV_PROF": true,
				}
			}

			result, err := s.Call("SUSR_SUIM_API_RSUSR100N", params)
			if err != nil {
				log.Fatalln(err)
			}

			users, ok := result["ET_USERS"]
			if !ok {
				arr := result["RETURN"].([]interface{})
				ret := arr[0].(map[string]interface{})
				log.Fatalf("%s\n", ret["MESSAGE"])
			}

			j, err := json.MarshalIndent(users, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%s\n", j)
		},
	}

	rootCmd.AddCommand(cmdChangeDocs)
	cmdChangeDocs.Flags().StringVarP(&file, "file", "f", "", "YAML file containing query parameters")

	var cmdUsers = &cobra.Command{
		Use:   "users [user]",
		Short: "Search users",
		Long:  `users will list users that match the parameters.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			var params map[string]interface{}
			if file != "" {
				params, err = readYaml(file)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				params = map[string]interface{}{
					"IT_USER": []map[string]interface{}{map[string]interface{}{
						"SIGN": "I",
						"LOW":  args[0],
					}},
				}
			}

			fmt.Printf("%#v\n", params)
			result, err := s.Call("SUSR_SUIM_API_RSUSR002", params)
			if err != nil {
				log.Fatalln(err)
			}

			users, ok := result["ET_USERS"]
			if !ok {
				log.Fatalf("Found no users: %v\n", result["RETURN"])
			}

			j, err := json.MarshalIndent(users, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%s\n", j)
		},
	}

	rootCmd.AddCommand(cmdUsers)
	cmdUsers.Flags().StringVarP(&file, "file", "f", "", "YAML file containing query parameters")

}
