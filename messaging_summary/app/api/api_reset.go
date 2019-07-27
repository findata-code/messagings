package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"messaging_summary/app"
	"messaging_summary/app/model"
	"net/http"
)

func GetLatestReset(uid string) (model.Reset, error) {
	url := fmt.Sprintf("%s?user_id=%s", app.Config.API.GetLatestReset, uid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Reset{}, err
	}

	req = AddAuth(req)
	res, err := CallREST(req)
	if err != nil {
		return model.Reset{}, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return model.Reset{}, err
	}

	var reset model.Reset
	err = json.Unmarshal(b, &reset)
	if err != nil {
		return model.Reset{}, err
	}

	return reset, nil
}
