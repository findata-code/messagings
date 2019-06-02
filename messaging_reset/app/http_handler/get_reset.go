package http_handler

import (
	"encoding/json"
	"fmt"
	"messaging_reset/app"
	"messaging_reset/app/model"
	"net/http"
)

func GetLatestReset (w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	userId := r.URL.Query()["user_id"]

	var re model.Reset

	err := app.Db.Where("user_id = ?", userId).Order("unix_nano desc").Limit(1).Find(&re).Error
	if err != nil {
		ResponseError(w, err)
		return
	}

	b, err := json.Marshal(re)
	if err != nil {
		ResponseError(w, err)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func ResponseError (w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, err.Error())
}

func ResponseUnauthorized(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, err.Error())
}
