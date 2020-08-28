package reader

type Type int8

type FileContent map[string]interface{}

type Reader interface {
	Open(filePath string) (FileContent, error)
}
