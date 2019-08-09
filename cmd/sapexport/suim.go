package sapexport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/antihax/sapexport/cmd/sapexport/sap"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var file string
var format string

var stdin bool

func getParameters() (map[string]interface{}, error) {
	if file != "" {
		return readYamlFile(file)
	} else if stdin {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		return readYaml(data)
	}
	return nil, nil
}

func outputResult(v interface{}) error {
	switch format {
	case "yaml":
		j, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", j)

	default:
		j, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", j)

	}
	return nil
}

func init() {

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

			params, err := getParameters()
			if err != nil {
				log.Fatalln(err)
			}

			result, err := s.Call(args[0], params)
			if err != nil {
				log.Fatalln(err)
			}

			if err := outputResult(result); err != nil {
				log.Fatalln(err)
			}
		},
	}

	rootCmd.AddCommand(cmdRun)
	cmdRun.Flags().StringVarP(&file, "file", "f", "", "YAML file containing query parameters")
	cmdRun.Flags().StringVarP(&format, "format", "F", "", "Output as yaml or json")
	cmdRun.Flags().BoolVarP(&stdin, "stdin", "S", false, "Read YAML parameters from STDIN")
}
