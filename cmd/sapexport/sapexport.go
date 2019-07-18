package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/antihax/sapexport/cmd/sapexport/sap"
	"github.com/sap/gorfc/gorfc"
	"github.com/spf13/cobra"
)

func main() {
	var (
		abapSystem gorfc.ConnectionParameter
		where      string
	)

	var cmdRoleUsers = &cobra.Command{
		Use:   "roleusers [role]",
		Short: "Export a list of users in a role to JSON",
		Long:  `roleusers will extract list of users with a specific role from the SAP system and return a JSON representation.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}
			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			rows, err := s.UsersOfRole(args[0])
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

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdTable)

	cmdTable.Flags().StringVarP(&where, "where", "w", "", "ABAP WHERE clause to filter the table")

	rootCmd.AddCommand(cmdRoleUsers)

	rootCmd.PersistentFlags().StringVarP(&abapSystem.User, "user", "u", "SAP*", "RFC Username")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Passwd, "pass", "p", "PASS", "RFC Password")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Lang, "language", "l", "EN", "System Language")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Client, "client", "c", "001", "System Client")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Ashost, "address", "a", "localhost", "System Address")
	rootCmd.PersistentFlags().StringVarP(&abapSystem.Saprouter, "router", "r", "", "Router")

	rootCmd.Execute()
}
