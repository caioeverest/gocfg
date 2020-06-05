package gocfg

import (
	"reflect"
	"strconv"

	"github.com/caioeverest/gocfg/reader"
)

func fill(content reader.FileContent, inputAddr interface{}) (err error) {
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
			sub, converted := convertSubStruct(rawVal)
			if !converted {
				return FailToParseSubObject(key)
			}
			if err = fill(sub, fieldVal.Addr().Interface()); err != nil {
				return
			}
		} else {
			if fieldStruct.Type != reflect.TypeOf(rawVal) {
				if reflect.TypeOf(rawVal).Kind() == reflect.String {
					value, err := interpolationConverter(fieldStruct.Type.Kind(), rawVal)
					if err != nil {
						return TypeMismatch(key, reflect.TypeOf(rawVal), fieldStruct.Type)
					}
					fieldVal.Set(value)
				} else {
					return TypeMismatch(key, reflect.TypeOf(rawVal), fieldStruct.Type)
				}
			} else {
				fieldVal.Set(reflect.ValueOf(rawVal))
			}
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

func interpolationConverter(kind reflect.Kind, rawvalue interface{}) (out reflect.Value, err error) {
	v := rawvalue.(string)

	switch kind {
	case reflect.Bool:
		var tmp bool
		tmp, err = strconv.ParseBool(v)
		out = reflect.ValueOf(tmp)
	case reflect.Int:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(tmp)
	case reflect.Int8:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int8(tmp))
	case reflect.Int16:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int16(tmp))
	case reflect.Int32:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int32(tmp))
	case reflect.Int64:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(int64(tmp))
	case reflect.Uint:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint(tmp))
	case reflect.Uint8:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint8(tmp))
	case reflect.Uint16:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint16(tmp))
	case reflect.Uint32:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint32(tmp))
	case reflect.Uint64:
		var tmp int
		tmp, err = strconv.Atoi(v)
		out = reflect.ValueOf(uint64(tmp))
	case reflect.Float32:
		var tmp float64
		tmp, err = strconv.ParseFloat(v, 32)
		out = reflect.ValueOf(float32(tmp))
	case reflect.Float64:
		var tmp float64
		tmp, err = strconv.ParseFloat(v, 64)
		out = reflect.ValueOf(tmp)
	default:
		out = reflect.ValueOf(v)
	}
	return
}
