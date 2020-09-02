package gocfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/caioeverest/gocfg/reader"
)

const (
	aliasTag    = "cfg"
	requiredTag = "required"
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
		}

		switch fieldVal.Kind() {
		case reflect.Struct:
			var (
				sub       reader.FileContent
				converted bool
			)

			if sub, converted = convertSubStruct(rawVal); !converted {
				return failToParseSubObject(key)
			}

			if err = fill(sub, fieldVal.Addr().Interface()); err != nil {
				return
			}
		case reflect.Ptr:
			var (
				err            error
				value          reflect.Value
				fieldStruct    = reflect.Indirect(reflect.ValueOf(inputAddr)).Type().Field(i)
				blockValueTrue = reflect.ValueOf(inputAddr)
				fieldValElem   = blockValueTrue.Elem().Field(i)
			)

			if value, err = interpolationConverter(fieldStruct.Type.Elem().Kind(), rawVal); err != nil {
				return err
			}

			initializer := reflect.New(fieldStruct.Type.Elem())
			initializer.Elem().Set(value)
			fieldValElem.Set(initializer)
		default:
			value, err := interpolationConverter(fieldStruct.Type.Kind(), rawVal)
			if err != nil {
				return typeMismatch(key, reflect.TypeOf(rawVal), fieldStruct.Type)
			}

			fieldVal.Set(value)
		}
	}

	return
}

func getValue(content reader.FileContent, key string, required bool) (interface{}, error) {
	value, ext := content[key]
	if !ext && required {
		return nil, requiredKeyNotFound(key)
	}
	return value, nil
}

func getParams(f reflect.StructField) (alias string, required bool) {
	var (
		tmp []string
		extAlias bool
		content string
	)

	if content, extAlias = f.Tag.Lookup(aliasTag); !extAlias {
		return f.Name, false
	}

	required = strings.Contains(content, requiredTag)
	tmp = strings.Split(content, ",")

	if required && len(tmp) == 1 {
		return f.Name, required
	}

	return tmp[0], required
}

func interpolationConverter(kind reflect.Kind, rawvalue interface{}) (out reflect.Value, err error) {
	v := fmt.Sprintf("%v", rawvalue)

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
