package error

import "encoding/json"

type DefaultError struct {
	HttpCode int
	Err      error
}

func (de DefaultError) Error() string {
	return de.Err.Error()
}

type UnauthorizeError struct {
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
}

func (ue UnauthorizeError) Error() string {
	out, _ := json.Marshal(ue)
	return string(out)
}

func NewUnauthorizeError() error {
	return UnauthorizeError{}
}

type NotFoundError struct {
	StatusMessage string `json:"status_message"`
	StatusCode    int    `json:"status_code"`
}

func (ne NotFoundError) Error() string {
	out, _ := json.Marshal(ne)
	return string(out)
}

func NewNotFoundError() error {
	return NotFoundError{}
}
