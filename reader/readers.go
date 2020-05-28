package reader

type Type int

type FileContent map[string]interface{}

type Reader interface {
	Open(filePath string) (FileContent, error)
}

const (
	ENV = iota
	YAML
	JSON
	TOML
)

func Selector(loadType Type) (selected Reader, extension string) {
	switch loadType {
	case ENV:
		selected = Env{}
		extension = ""
	case YAML:
		selected = Yaml{}
		extension = "yml"
	//case JSON:
	//	selected = Json{}
	//	extension = "json"
	//case TOML:
	//	selected = Toml{}
	//	extension = "toml"
	default:
		selected = Env{}
		extension = ""
	}
	return
}
