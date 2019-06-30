package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messaging_delete/app"
	"messaging_delete/app/model"
	"net/http"
)

func GetLastNIncome(uid string, n int) ([]model.Income, error) {
	url := fmt.Sprintf("%s?user_id=%s&n=%d", app.Config.API.GetLastNIncomes, uid, n)
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

	incomes := make([]model.Income, 0)
	err = json.Unmarshal(b, &incomes)
	if err != nil {
		return nil, err
	}

	return incomes, nil
}

func DeleteIncome(uid string, id int) error {
	url := fmt.Sprintf("%s?user_id=%s&id=%d", app.Config.API.DeleteIncome, uid, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req = AddAuth(req)
	_, err = CallREST(req)

	return err
}
