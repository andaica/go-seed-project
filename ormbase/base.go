package ormbase

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	ex: Connect("root", "password", "@tcp(127.0.0.1:3306)", "goexample")
*/
func Connect(user, password, address, dbname string) (db *gorm.DB, err error) {
	conf := user + ":" + password + address + "/" + dbname + "?charset=utf8&parseTime=True"
	db, err = gorm.Open("mysql", conf)
	return db, err
}
