package reader

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Yaml struct{}

func (Yaml) Open(filepath string) (output FileContent, err error) {
	var rawFile []byte
	rawFile, err = ioutil.ReadFile(filepath)
	err = yaml.Unmarshal(rawFile, &output)
	return
}
