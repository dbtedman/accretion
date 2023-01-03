package git

import (
	"errors"
	"regexp"
)

func ParseCloneHTTP(value string) (CloneHTTP, error) {
	if err := validateCloneHTTPValue(value); err != nil {
		return nil, err
	}

	return cloneHTTP{
		value: value,
	}, nil
}

func validateCloneHTTPValue(value string) error {
	if !regexp.MustCompile("https?://.+\\.git$").MatchString(value) {
		return errors.New("value must match expected pattern")
	}
	return nil
}

type CloneHTTP interface {
	String() string
}

type cloneHTTP struct {
	value string
}

// ensure cloneHTTP implements CloneHTTP interface
var _ CloneHTTP = cloneHTTP{}

func (my cloneHTTP) String() string {
	return my.value
}
