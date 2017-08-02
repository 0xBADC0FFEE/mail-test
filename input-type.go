package main

type InputType string

func (it *InputType) Set(v string) error {
	*it = InputType(v)
	return nil
}

func (it InputType) String() string {
	return string(it)
}

const (
	URL  InputType = "url"
	FILE InputType = "file"
)
