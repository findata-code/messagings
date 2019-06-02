package app

import "net/http"

func AddAuth(r *http.Request) *http.Request {
	r.Header.Set(AuthToken, config.AuthToken)
	return r
}

func CallREST(r *http.Request) (*http.Response, error){
	client := http.Client{}
	return client.Do(r)
}
