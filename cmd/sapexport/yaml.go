package sapexport

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func readYaml(path string) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
