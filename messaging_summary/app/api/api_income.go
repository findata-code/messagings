package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messaging_summary/app"
	"messaging_summary/app/model"
	"net/http"
)

func GetIncomes(uid, since string) ([]model.Income, error) {
	url := fmt.Sprintf("%s?user_id=%s&since=%s", app.Config.API.GetIncomes, uid, since)
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
