package gocfg

import (
	"reflect"

	"github.com/caioeverest/gocfg/reader"
)

func Fill(content reader.FileContent, inputAddr interface{}) (err error) {
	var (
		strValue = reflect.Indirect(reflect.ValueOf(inputAddr))
		strType  = strValue.Type()
	)

	for i := 0; i < strValue.NumField(); i++ {
		var (
			rawVal        interface{}
			fieldVal      = strValue.Field(i)
			fieldStruct   = strType.Field(i)
			key, required = getParams(fieldStruct)
		)

		if rawVal, err = getValue(content, key, required); err != nil {
			return
		}

		if rawVal == nil {
			continue
		} else if fieldVal.Kind() == reflect.Struct {
			sub, converted := tryConversion(rawVal)
			if !converted {
				return FailToParseSubObject(key)
			}
			if err = Fill(sub, fieldVal.Addr().Interface()); err != nil {
				return
			}
		} else {
			if fieldStruct.Type != reflect.TypeOf(rawVal) {
				return TypeMismatch(key, reflect.TypeOf(rawVal), fieldStruct.Type)
			}

			fieldVal.Set(reflect.ValueOf(rawVal))
		}
	}

	return
}

func getValue(content reader.FileContent, key string, required bool) (interface{}, error) {
	value, ext := content[key]
	if !ext && required {
		return nil, RequiredKeyNotFound(key)
	}
	return value, nil
}

func getParams(f reflect.StructField) (alias string, required bool) {
	const (
		aliasTag    = "cfg"
		requiredTag = "required"
	)
	alias, extAlias := f.Tag.Lookup(aliasTag)
	_, required = f.Tag.Lookup(requiredTag)

	if !extAlias {
		alias = f.Name
	}
	return
}

func tryConversion(rawInput interface{}) (output reader.FileContent, converted bool) {
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
