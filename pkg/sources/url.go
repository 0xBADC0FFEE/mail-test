package sources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type UrlReader struct {
	Path string
}

func (ur *UrlReader) Get() ([]byte, error) {
	path := ur.Path

	uri, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %s: %v", path, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, fmt.Errorf("http get %s failed: %v", path, err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get not ok, status: %v", resp.StatusCode, respBody)
	}

	return respBody, nil
}
