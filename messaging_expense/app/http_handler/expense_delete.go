package http_handler

import (
	"errors"
	"fmt"
	"messaging_expense/app"
	"messaging_expense/app/model"
	"net/http"
)

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	if r.Method != http.MethodDelete {
		ResponseMethodNotAllowed(w, errors.New(fmt.Sprintf("function doesn't support %s method", r.Method)))
		return
	}

	if err := requireField(r.URL.Query(), "user_id", "since"); err != nil {
		ResponseError(w, errors.New("fields are missing"))
		return
	}

	userId := r.URL.Query()["user_id"]
	id := r.URL.Query()["id"]

	if err := app.Db.Where("user_id = ? AND id = ?", userId, id).Delete(&model.Expense{}).Error; err != nil {
		ResponseError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
