package sources

type Source interface {
	Get() ([]byte, error)
}

func NewUrlReader(path string) Source {
	return &UrlReader{Path: path}
}

func NewFileReader(path string) Source {
	return &FileCounter{Path: path}
}
