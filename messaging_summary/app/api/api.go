package api

import (
	"messaging_summary/app"
	"net/http"
)

const (
	AuthToken = "X-Authorization-Token"
)

func AddAuth(r *http.Request) *http.Request {
	r.Header.Set(AuthToken, app.Config.AuthToken)
	return r
}

func CallREST(r *http.Request) (*http.Response, error){
	client := http.Client{}
	return client.Do(r)
}
