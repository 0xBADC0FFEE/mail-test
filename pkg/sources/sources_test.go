package sources

import (
	"bytes"
	"testing"
)

func TestUrlReader(t *testing.T) {
	u := "https://golang.org"
	s := "Go"
	expect := 9

	ur := NewUrlReader(u)

	data, err := ur.Get()
	if err != nil {
		t.Errorf("Can't read url: %v", err)
	}

	count := bytes.Count(data, []byte(s))

	t.Logf("Count for %s: %d", u, count)

	if count != expect {
		t.Errorf("Wrong count for %s: expected %d got %d", u, expect, count)
	}

}

func TestFileReader(t *testing.T) {
	u := "/etc/passwd"
	s := "Go"
	expect := 0

	ur := NewFileReader(u)

	data, err := ur.Get()
	if err != nil {
		t.Errorf("Can't read url: %v", err)
	}

	count := bytes.Count(data, []byte(s))

	t.Logf("Count for %s: %d", u, count)

	if count != expect {
		t.Errorf("Wrong count for %s: expected %d got %d", u, expect, count)
	}

}
