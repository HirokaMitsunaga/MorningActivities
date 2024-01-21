package main

import (
	"MorningActivities-API/controller"
	"MorningActivities-API/db"
	"MorningActivities-API/repository"
	"MorningActivities-API/router"
	"MorningActivities-API/usecase"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	//エラーが出たらLoggerがエラー出力して、強制終了する
	e.Logger.Fatal(e.Start(":8080"))
}