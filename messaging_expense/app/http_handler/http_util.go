package http_handler

import (
	"fmt"
	"net/http"
)

func ResponseError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, err.Error())
}

func ResponseUnauthorized(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, err.Error())
}

func ResponseMethodNotAllowed(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, err.Error())
}
