package comphouse

import (
	"errors"
	"net/http"
)

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrNotFound         = errors.New("not found")
	ErrTooManyRequests  = errors.New("too many requests")
	ErrUnexpectedStatus = errors.New("unexpected status")
)

func statusCodeToError(status int) error {
	errors := map[int]error{
		http.StatusUnauthorized:    ErrUnauthorized,
		http.StatusNotFound:        ErrNotFound,
		http.StatusTooManyRequests: ErrTooManyRequests,
	}

	if err, ok := errors[status]; ok {
		return err
	}

	if status < 200 || status >= 300 {
		return ErrUnexpectedStatus
	}

	return nil
}
