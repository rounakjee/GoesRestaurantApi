package Psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type PSqlDB struct {
	*gorm.DB
}

func NewPSqlDB(username string,password string, host string, schema string) *PSqlDB {

	DB, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, schema))
	if err != nil {
		panic("Can't connect to db" + err.Error())
	}
	return &PSqlDB{
		DB,
	}
}
