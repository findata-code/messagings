package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Reset struct {
	gorm.Model
	UserId      string
	FullMessage string
	Timestamp   time.Time
	UnixNano    string
}
