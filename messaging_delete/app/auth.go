package app

import (
	"errors"
	"net/http"
)

const AuthToken = "X-Authorization-Token"

func CheckAuth(r *http.Request) error {
	v := r.Header.Get(AuthToken)
	if v != config.AuthToken {
		return errors.New("unauthorized request")
	}

	return nil
}
