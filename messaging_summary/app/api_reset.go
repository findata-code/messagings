package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetLatestReset (uid string) (Reset, error){
	url := fmt.Sprintf("%s?user_id=%s", config.API.GetLatestReset, uid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Reset{}, err
	}

	req = AddAuth(req)
	res, err := CallREST(req)
	if err != nil {
		return Reset{}, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Reset{}, err
	}

	var reset Reset
	err = json.Unmarshal(b, &reset)
	if err != nil {
		return Reset{}, err
	}

	return reset, nil
}
