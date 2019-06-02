package http_handler

import (
	"errors"
	"messaging_reset/app"
	"net/http"
)

const AuthToken = "X-Authorization-Token"

func CheckAuth(r *http.Request) error {
	v := r.Header.Get(AuthToken)
	if v != app.Config.AuthToken {
		return errors.New("unauthorized request")
	}

	return nil
}
