package gocfg

import (
	"fmt"

	"github.com/caioeverest/gocfg/reader"
)

//It start's the Gonfig process, setting the structure used as reference and the type of the configuration:
//- Environment variables is gonfig.ENV or 0
//- Yaml file is gonfig.YAML or 1
//if no file name is informed the function will search on the current path for a application.yml
func Load(s interface{}, loadType reader.Type, files ...string) (err error) {
	var (
		r                   reader.Reader
		extension           string
		locations           []string
		fileRawContent      reader.FileContent
		interpolatedContent reader.FileContent
	)

	r, extension = reader.Selector(loadType)
	locations = filename(extension, files)
	if fileRawContent, err = r.Open(locations[0]); err != nil {
		return
	}
	interpolatedContent = interpolate(fileRawContent)

	return Fill(interpolatedContent, s)
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
