package gocfg

import (
	"testing"

	"github.com/caioeverest/gocfg/reader"
	"github.com/stretchr/testify/assert"
)

var (
	inputOk = reader.FileContent{
		"some_pure_key":                      "exemple",
		"key_with_interpolation":             "{GOPATH}",
		"key_with_interpolation_and_default": "{TEST:default}",
		"flat_key_with_number":               58,
	}
)

func TestInterpolate_TestWithSuccess(t *testing.T) {
	inttOut := interpolate(inputOk)
	print(inttOut)
	assert.NotNil(t, inttOut)
}
