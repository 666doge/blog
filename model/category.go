package model

import (
	"time"
)

type Category struct {
	Id int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	CreateTime time.Time `json:"createTime" db:"create_time"`
	UpdateTime time.Time  `json:"updateTime" db:"update_time"`
}