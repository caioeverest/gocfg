package gocfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	var testObject = struct {
		SomeKey    string `cfg:"some_key"`
		AnotherKey int    `cfg:"another_key"`
		SubObject  struct {
			SomeThing     string `cfg:"something"`
			SomethingElse int    `cfg:"something_else"`
		} `cfg:"sub_object"`
		LastKey           string `cfg:"last_key"`
		LastKeyWithNumber uint   `cfg:"last_key_with_number"`
	}{}

	err := Load(&testObject, YAML, "test.yml")
	assert.Nil(t, err)
	fmt.Printf("%+v", testObject)
}
