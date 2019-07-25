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

	var cmdUserProfiles = &cobra.Command{
		Use:   "userprofiles [user]",
		Short: "List profiles for a user",
		Long:  `userprofiles will list profiles for a user.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			params := map[string]interface{}{
				"IT_VALUES": []map[string]interface{}{
					map[string]interface{}{
						"IT_VALUES": args[0],
					}},
			}
			rows, err := s.Call("SUSR_GET_PROFILES_OF_USER_RFC", params)
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
	var cmdFindUsers = &cobra.Command{
		Use:   "findusers [user]",
		Short: "List profiles for a user",
		Long:  `authobjusers will list profiles for a user.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			params := map[string]interface{}{
				"IT_USER": []map[string]interface{}{map[string]interface{}{
					"SIGN": "I",
					"LOW":  args[0],
				}},
			}
			rows, err := s.Call("SUSR_SUIM_API_RSUSR002", params)
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
	var cmdProfileUsers = &cobra.Command{
		Use:   "profileusers [profile]",
		Short: "List users with profile",
		Long:  `profileusers will list users with a given profile.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s := sap.RFC{}

			err := s.Connect(abapSystem)
			if err != nil {
				log.Fatalln(err)
			}

			params := map[string]interface{}{
				"PROFILES": []map[string]interface{}{map[string]interface{}{
					"PROFILE": args[0],
				}},
			}
			rows, err := s.Call("SUSR_GET_USERS_WITH_PROFS_RFC", params)
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

	cmdTable.Flags().StringVarP(&where, "where", "w", "", "ABAP WHERE clause to filter the table")

	rootCmd.AddCommand(cmdTable)
	rootCmd.AddCommand(cmdProfileUsers)
	rootCmd.AddCommand(cmdUserProfiles)
	rootCmd.AddCommand(cmdFindUsers)
}
