package http_handler

import (
	"errors"
	"net/url"
)

func requireField(vs url.Values, fs... string) error {
	for _, f := range fs {
		if v := vs[f]; v == nil {
			return errors.New("fields are missing")
		}
	}

	return nil
}
