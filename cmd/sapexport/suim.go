package sapexport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/antihax/sapexport/cmd/sapexport/sap"
	"github.com/spf13/cobra"
)

func init() {
	var file string
	var stdin bool

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
				params, err = readYamlFile(file)
				if err != nil {
					log.Fatalln(err)
				}
			} else if stdin {
				data, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Fatalln(err)
				}
				params, err = readYaml(data)
				if err != nil {
					log.Fatalln(err)
				}
			}

			result, err := s.Call(args[0], params)
			if err != nil {
				log.Fatalln(err)
			}

			j, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%s\n", j)
		},
	}

	rootCmd.AddCommand(cmdRun)
	cmdRun.Flags().StringVarP(&file, "file", "f", "", "YAML file containing query parameters")
	cmdRun.Flags().BoolVarP(&stdin, "stdin", "S", false, "Read YAML parameters from STDIN")

}
