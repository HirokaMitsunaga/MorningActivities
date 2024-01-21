package main

import (
	"MorningActivities-API/db"
	"MorningActivities-API/model"
	"fmt"
)

func main() {
	dbConn:= db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}