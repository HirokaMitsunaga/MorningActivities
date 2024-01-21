package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/router"
	"go-api/usecase"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController,taskController)
	//エラーが出たらLoggerがエラー出力して、強制終了する
	e.Logger.Fatal(e.Start(":8080"))
}