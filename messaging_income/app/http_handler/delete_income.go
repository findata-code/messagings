package http_handler

import (
	"errors"
	"fmt"
	"messaging_income/app"
	"messaging_income/app/model"
	"net/http"
)

func DeleteIncome(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	if r.Method != http.MethodDelete {
		ResponseMethodUnsupport(w, errors.New(fmt.Sprintf("function doesn't support %s method", r.Method)))
		return
	}

	userId := r.URL.Query()["user_id"]
	id := r.URL.Query()["id"]

	if userId == nil {
		ResponseError(w, errors.New("fields are missing"))
		return
	}

	if id == nil {
		ResponseError(w, errors.New("fields are missing"))
		return
	}

	if err := app.Db.Where("user_id = ? AND id = ?", userId, id).Delete(&model.Income{}).Error; err != nil {
		ResponseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
