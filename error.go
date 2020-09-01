package gocfg

import (
	"fmt"
	"reflect"
)

// Error alias helper
type Error string

const (
	requiredKeyNotFoundTmpl  = "key %s wasn`t found on"
	typeMismatchTmpl         = "key %s has type %s, but it expects %s"
	failToParseSubObjectTmpl = "the structure %s encounter problems to be converted into a map"
)

func (e Error) Error() string { return string(e) }

func requiredKeyNotFound(key string) error {
	return Error(fmt.Sprintf(requiredKeyNotFoundTmpl, key))
}

func typeMismatch(key string, found, expected reflect.Type) error {
	return Error(fmt.Sprintf(typeMismatchTmpl, key, found, expected))
}

func failToParseSubObject(key string) error {
	return Error(fmt.Sprintf(failToParseSubObjectTmpl, key))
}
