package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetExpenses(uid, since string) ([]Expense, error) {
	url := fmt.Sprintf("%s?user_id=%s&since=%s", config.API.GetExpenses, uid, since)
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
	fmt.Println(url)
	incomes := make([]Expense, 0)
	err = json.Unmarshal(b, &incomes)
	if err != nil {
		return nil, err
	}

	return incomes, nil
}

