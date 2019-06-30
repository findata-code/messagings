package utils

import (
	"strconv"
	"strings"
)

const (
	LatestExpense = -1
)

func GetEntityIdentifier(s string) (int64, error) {
	if strings.HasSuffix(s, "ล่าสุด") {
		return -1, nil
	}

	ss := strings.Split(s, " ")
	return strconv.ParseInt(ss[len(ss) - 1], 10, 64)
}