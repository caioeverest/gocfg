package gocfg

import (
	"fmt"
	"testing"

	"github.com/caioeverest/gocfg/reader"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	First              string `cfg:"required"`
	Second             string
	SomeBooleanExample bool
	SomePointerInt     *int
	SomePointerString  *string
	SomeIntExample     int
	SubStruct          SubStructExample
}

type SubStructExample struct {
	SubFirst  float64
	SubSecond bool
}

var (
	inputOkMapper = reader.FileContent{
		"First":              "exemple",
		"Second":             "some value",
		"SomeBooleanExample": true,
		"SomePointerInt":     3123,
		"SomeIntExample":     58,
		"SubStruct": reader.FileContent{
			"SubFirst":  1280.8,
			"SubSecond": false,
		},
	}

	inputWithoutRequiredFieldMapper = reader.FileContent{
		"Second":             "some value",
		"SomeBooleanExample": true,
		"SomePointerInt":     3123,
		"SomeIntExample":     58,
		"SubStruct": reader.FileContent{
			"SubFirst":  1280.8,
			"SubSecond": false,
		},
	}
)

func TestFill(t *testing.T) {
	var c = Example{}
	err := fill(inputOkMapper, &c)
	assert.Nil(t, err)
	fmt.Printf("%+v", c)
}

func TestFill_ShouldErrorWhenRequiredFieldIsMissing(t *testing.T) {
	var c = Example{}
	err := fill(inputWithoutRequiredFieldMapper, &c)
	assert.NotNil(t, err)
}
