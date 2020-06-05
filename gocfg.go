package gocfg

import (
	"fmt"

	"github.com/caioeverest/gocfg/reader"
)

const (
	ENV = iota
	YAML
	JSON
	TOML
)

//It start's the gocfg process, setting the structure used as reference and the type of the configuration:
//- Environment variables is gocfg.ENV or 0
//- Yaml file is gocfg.YAML or 1
//if no file name is informed the function will search on the current path for a application.yml
func Load(s interface{}, loadType reader.Type, files ...string) (err error) {
	var (
		r                   reader.Reader
		extension           string
		locations           []string
		fileRawContent      reader.FileContent
		interpolatedContent reader.FileContent
	)

	r, extension = selector(loadType)
	locations = filename(extension, files)
	if fileRawContent, err = r.Open(locations[0]); err != nil {
		return
	}
	interpolatedContent = interpolate(fileRawContent)

	return fill(interpolatedContent, s)
}

func selector(loadType reader.Type) (selected reader.Reader, extension string) {
	switch loadType {
	case ENV:
		selected = reader.Env{}
		extension = ""
	case YAML:
		selected = reader.Yaml{}
		extension = "yml"
	//case JSON:
	//	selected = Json{}
	//	extension = "json"
	//case TOML:
	//	selected = Toml{}
	//	extension = "toml"
	default:
		selected = reader.Env{}
		extension = ""
	}
	return
}

func filename(extension string, filesLocation []string) (output []string) {
	const base = "application"
	if len(filesLocation) == 0 {
		output = append(output, fmt.Sprintf("%s.%s", base, extension))
	} else {
		for _, path := range filesLocation {
			output = append(output, path)
		}
	}
	return
}
