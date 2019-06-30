package http_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"messaging_income/app"
	"messaging_income/app/model"
	"net/http"
	"strconv"
)

func GetIncome(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	userId := r.URL.Query()["user_id"]
	since := r.URL.Query()["since"]

	incomes := make([]model.Income, 0)

	err := app.Db.Where("user_id = ? AND unix_nano >= ?", userId, since).Find(&incomes).Error
	if err != nil {
		ResponseError(w, err)
		return
	}

	b, err := json.Marshal(incomes)
	if err != nil {
		ResponseError(w, err)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func GetLastNIncomes(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	if r.Method != http.MethodGet {
		ResponseMethodUnsupport(w, errors.New(fmt.Sprintf("function doesn't support %s method", r.Method)))
		return
	}

	if err := requireField(r.URL.Query(), "user_id", "n"); err != nil {
		ResponseError(w, errors.New("fields are missing"))
		return
	}

	userId := r.URL.Query()["user_id"]
	an := r.URL.Query()["n"]

	n, err := strconv.ParseInt(an[0], 10, 64)
	if err != nil {
		ResponseError(w, err)
		return
	}

	incomes := make([]model.Income, 0)
	if err := app.Db.Where("user_id = ?", userId).Limit(n).Order("id desc").Find(&incomes).Error; err != nil {
		ResponseError(w, err)
		return
	}

	b, err := json.Marshal(incomes)
	if err != nil {
		ResponseError(w, err)
	}

	fmt.Fprintf(w, "%s", string(b))
}
