package model

import (

	"myProject/datastore/postgres"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

const (
	queryInsert   = "INSERT INTO userdata(username, email, password , avatar) VALUES($1, $2, $3, $4);"
	queryData = "SELECT * from userdata WHERE email = $1 AND password = $2"

)

func (s *User) Create() error {
	_, err := postgres.Db.Exec(queryInsert, s.Username, s.Email, s.Password , s.Avatar)
	return err
}

func (a *User) Check(email string) error {
	const queryCheck = "Select * from userdata where email = $1;"
	err := postgres.Db.QueryRow(queryCheck,email).Scan(&a.Id,&a.Username,&a.Email,&a.Password,&a.Avatar)
	return err
}

