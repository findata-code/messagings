package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetExpense(w http.ResponseWriter, r *http.Request) {
	if err := CheckAuth(r); err != nil {
		ResponseUnauthorized(w, err)
		return
	}

	userId := r.URL.Query()["user_id"]
	since := r.URL.Query()["since"]

	expenses := make([]Expense, 0)
	err := db.Where("user_id = ? AND unix_nano >= ?", userId, since).Find(&expenses).Error
	if err != nil {
		ResponseError(w, err)
		return
	}

	b, err := json.Marshal(expenses)
	if err != nil {
		ResponseError(w, err)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

