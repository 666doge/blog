package db

import (
	"blog/model"
)

func GetCategoryList() (list []*model.Category, err error) {
	sqlStr := `select id, name, create_time, update_time from category`
	err = DB.Select(&list, sqlStr)
	return list, err
}
