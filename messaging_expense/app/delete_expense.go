package app

import (
	"errors"
	"fmt"
	"net/http"
)

func DeleteExpese(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	if r.Method != http.MethodDelete {
		ResponseError(w, errors.New(fmt.Sprintf("function doesn't support %s method", r.Method)))
		return
	}

	userId := r.URL.Query()["user_id"]

	var ex Expense
	if err := db.Where("user_id = ?", userId).Last(&ex).Error; err != nil {
		ResponseError(w, err)
		return
	}

	if err := db.Delete(&ex).Error; err != nil {
		ResponseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
