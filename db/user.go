package db

import (
	"blog/model"
)

func InsertUser(user *model.User) (userId int64, err error){
	sqlStr := `insert into
				user(name, sex, age)
			values(?,?,?)`
	result, err := DB.Exec(sqlStr, user.Name, user.Sex, user.Age)
	if err != nil {
		return
	}
	userId, err = result.LastInsertId()
	return
}

func GetUserById(userId int64) (user *model.User, err error) {
	sqlStr := `select 
					id, name, sex, age
				from 
					user
				where
					id = ?
			`
	user = &model.User{}	
	err = DB.Get(user, sqlStr, userId)
	return
}

func GetUserList() (userList []*model.User, err error){
	sqlStr := `select id, name, age, sex from user`
	err = DB.Select(&userList, sqlStr)
	return
}

func DeleteUserById(id int64) error {
	sqlStr := `delete from user where id = ?`
	// sqlStr := `update user set status = 0 where id = ?`
	_, err := DB.Exec(sqlStr, id)
	return err
}