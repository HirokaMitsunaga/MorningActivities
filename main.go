package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/router"
	"go-api/usecase"
	"go-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewuserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	timelineRepository := repository.NewTimelineRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	timelineUsecase := usecase.NewTimelineUsecase(timelineRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	timelineController := controller.NewTimelineController(timelineUsecase)
	e := router.NewRouter(userController, taskController, timelineController)
	//エラーが出たらLoggerがエラー出力して、強制終了する
	e.Logger.Fatal(e.Start(":8080"))
}
