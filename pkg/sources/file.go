package sources

import (
	"fmt"
	"io/ioutil"
)

type FileCounter struct {
	Path string
}

func (fileCounter *FileCounter) Get() ([]byte, error) {
	path := fileCounter.Path

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file %s: %v", path, err)
	}

	return file, nil
}
