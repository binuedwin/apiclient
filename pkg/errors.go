package pkg

import (
	"errors"
	"fmt"
)

type Errors struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Errors) ToError() error {
	if len(e.Errors) > 0 {
		var Error string
		for _, err := range e.Errors {
			Error += fmt.Sprintf("%s: %s;", err.Code, err.Message)
		}

		return errors.New(Error[:len(Error)-1])
	}

	return nil
}
