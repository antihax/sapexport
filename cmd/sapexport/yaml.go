package sapexport

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func readYamlFile(path string) (map[string]interface{}, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return readYaml(yamlFile)
}

func readYaml(data []byte) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	err := yaml.Unmarshal(data, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
