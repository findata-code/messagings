package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messaging_delete/app"
	"messaging_delete/app/model"
	"net/http"
)

func GetExpenses(uid, since string) ([]model.Expense, error) {
	url := fmt.Sprintf("%s?user_id=%s&since=%s", app.Config.API.GetExpenses, uid, since)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = AddAuth(req)
	res, err := CallREST(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	expenses := make([]model.Expense, 0)
	err = json.Unmarshal(b, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func GetLastNExpense(uid string, n int) ([]model.Expense, error) {
	url := fmt.Sprintf("%s?user_id=%s&n=%d", app.Config.API.GetLastNExpenses, uid, n)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = AddAuth(req)
	res, err := CallREST(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	expenses := make([]model.Expense, 0)
	err = json.Unmarshal(b, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func DeleteExpense(uid string, id int) error {
	url := fmt.Sprintf("%s?user_id=%s&id=%d", app.Config.API.DeleteExpense, uid, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req = AddAuth(req)
	_, err = CallREST(req)

	return err
}
