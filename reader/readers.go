package reader

type Type int

type FileContent map[string]interface{}

type Reader interface {
	Open(filePath string) (FileContent, error)
}
