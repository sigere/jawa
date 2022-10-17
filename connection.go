package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	*gorm.DB
}

func getConnection() Connection {
	dsn := "root:root@tcp(127.0.0.1:4306)/jawa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return Connection{db}
}
