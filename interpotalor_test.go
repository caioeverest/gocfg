package gocfg

import (
	"testing"

	"github.com/caioeverest/gocfg/reader"
	"github.com/stretchr/testify/assert"
)

var (
	inputOkInterpolator = reader.FileContent{
		"some_pure_key":                                "exemple",
		"key_with_interpolation":                       "{GOPATH}",
		"key_with_interpolation_and_default":           "{TEST:default}",
		"key_with_interpolation_and_default_as_number": "{TEST:89}",
		"flat_key_with_number":                         58,
	}
)

func TestInterpolate_TestWithSuccess(t *testing.T) {
	inttOut := interpolate(inputOkInterpolator)
	print(inttOut)
	assert.NotNil(t, inttOut)
}
