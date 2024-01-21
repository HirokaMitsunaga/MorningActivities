package main

import (
	"fmt"
	"go-api/API/db"
	"go-api/API/model"
)

func main() {
	dbConn:= db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}