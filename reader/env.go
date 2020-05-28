package reader

type Env struct{}

func (Env) Open(string) (FileContent, error) {
	return nil, nil
}
