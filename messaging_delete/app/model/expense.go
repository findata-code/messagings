package app

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserId      string
	Value       float64
	FullMessage string
	Timestamp   time.Time
	UnixNano    string
}
