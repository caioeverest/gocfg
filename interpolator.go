package gocfg

import (
	"os"
	"regexp"
	"strings"

	"github.com/caioeverest/gocfg/reader"
)

const (
	isolatorRegex = "{(.+?)}"
	divisor       = ":"
)

func interpolate(content reader.FileContent) (output reader.FileContent) {
	output = make(reader.FileContent)
	for key, valueRaw := range content {
		if value, ok := valueRaw.(string); ok {
			output[key] = exportValue(value)
		} else {
			output[key] = valueRaw
		}
	}
	return
}

func exportValue(vraw string) string {
	var isolator = regexp.MustCompile(isolatorRegex)
	if isolator.MatchString(vraw) {
		v := removeBraces(vraw)
		key, def := splitKeyAndDefault(v)
		if value, ext := os.LookupEnv(key); ext {
			return value
		}
		return def
	}
	return vraw
}

func removeBraces(value string) string {
	var braces = regexp.MustCompile("[{}]")
	return braces.ReplaceAllString(value, "")
}

func splitKeyAndDefault(value string) (key, def string) {
	ref := strings.Split(value, divisor)
	if len(ref) > 1 {
		return ref[0], ref[1]
	}
	return ref[0], ""
}
