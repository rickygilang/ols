package helpers

import (
	"encoding/json"
)

type ResponseError struct {
	HttpResponseCode int
	ErrorMessage     string
	ErrorDetail      error
}

func (re *ResponseError) Error() string {
	out, err := json.Marshal(re)
	if err != nil {
		panic(err)
	}
	return string(out)
}
