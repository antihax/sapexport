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

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "run function module",
		Long:  `run a function module over RFC.`,
		Args:  cobra.MinimumNArgs(1),
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
			}

			result, err := s.Call(args[0], params)
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

	rootCmd.AddCommand(cmdRun)
	cmdRun.Flags().StringVarP(&file, "file", "f", "", "YAML file containing query parameters")

}
