package gocfg_test

import (
	"fmt"
	"testing"

	"github.com/caioeverest/gocfg"
	"github.com/caioeverest/gocfg/reader"
	"github.com/stretchr/testify/assert"
)

type Example struct {
	First              string
	Second             string
	SomeBooleanExample bool
	SomeIntExample     int
	SubStruct          SubStructExample
}

type SubStructExample struct {
	SubFirst  float64
	SubSecond bool
}

var (
	inputOk = reader.FileContent{
		"First":              "exemple",
		"Second":             "some value",
		"SomeBooleanExample": true,
		"SomeIntExample":     58,
		"SubStruct": reader.FileContent{
			"SubFirst":  1280.8,
			"SubSecond": false,
		},
	}
)

func TestFill(t *testing.T) {
	var c = Example{}
	err := gocfg.Fill(inputOk, &c)
	assert.Nil(t, err)
	fmt.Printf("%+v", c)
}
