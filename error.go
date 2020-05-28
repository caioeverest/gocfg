package gocfg

import (
	"fmt"
	"reflect"
)

type Error string

const (
	requiredKeyNotFoundTmpl = "key %s wasn`t found on"
	typeMismatch            = "key %s has type %s, but it expects %s"
	failToParseSubObject    = "the structure %s encounter problems to be converted into a map"
)

func (e Error) Error() string { return string(e) }

func RequiredKeyNotFound(key string) error {
	return Error(fmt.Sprintf(requiredKeyNotFoundTmpl, key))
}

func TypeMismatch(key string, found, expected reflect.Type) error {
	return Error(fmt.Sprintf(typeMismatch, key, found, expected))
}

func FailToParseSubObject(key string) error {
	return Error(fmt.Sprintf(failToParseSubObject, key))
}
