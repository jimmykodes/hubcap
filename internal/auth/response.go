package auth

import (
	"errors"
	"net/url"
)

type Response struct {
	Error string
	Code  string
	State string
}

func (r Response) Validate(sessionState string) error {
	if r.Error != "" {
		if r.Error == "access_denied" {
			// user canceled just return to home screen
			return ErrCanceled
		}
		return errors.New(r.Error)
	}
	if r.State == "" {
		return ErrNoState
	}
	if r.State != sessionState {
		return ErrMismatchedState
	}
	if r.Code == "" {
		return ErrNoCode
	}
	return nil
}

func NewResponse(query url.Values) *Response {
	return &Response{
		Error: query.Get("error"),
		Code:  query.Get("code"),
		State: query.Get("state"),
	}
}
