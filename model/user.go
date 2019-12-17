package model

type Sex int

const (
	USER_SEX_MAN = 0
	USER_SEX_WOMAN = 1
)

type User struct {
	Name string `json:"name" db:"name"`
	Age int `json:"age" db:"age"`
	Sex Sex `json:"sex" db:"sex"`
	Id int `json:"id" db:"id"`
}