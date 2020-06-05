package gocfg

import "github.com/caioeverest/gocfg/reader"

func convertSubStruct(rawInput interface{}) (output reader.FileContent, converted bool) {
	var (
		key   string
		input map[interface{}]interface{}
	)

	if output, converted = rawInput.(reader.FileContent); converted {
		return
	}

	output = make(reader.FileContent)
	input, converted = rawInput.(map[interface{}]interface{})
	for rawKey, value := range input {
		if key, converted = rawKey.(string); !converted {
			return
		}
		output[key] = value
	}
	return
}
