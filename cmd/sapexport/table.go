package sapexport

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/sapexport/cmd/sapexport/sap"
	"github.com/spf13/cobra"
)

func init() {
	var where string
	var cmdTable = &cobra.Command{
		Use:   "table [SAP Table]",
		Short: "Extract a table to JSON",
		Long:  `table will extract a table from the SAP system and return a JSON representation.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			rows, err := s.ReadTable(args[0], where)
			if err != nil {
				log.Fatalln(err)
			}

			j, err := json.MarshalIndent(rows, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%s\n", j)
		},
	}

	cmdTable.Flags().StringVarP(&where, "where", "w", "", "ABAP WHERE clause to filter the table")

	rootCmd.AddCommand(cmdTable)
}
