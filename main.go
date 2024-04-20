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
	likeRepository := repository.NewLikeRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	timelineUsecase := usecase.NewTimelineUsecase(timelineRepository)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	timelineController := controller.NewTimelineController(timelineUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	e := router.NewRouter(userController, taskController, timelineController, likeController)
	//エラーが出たらLoggerがエラー出力して、強制終了する
	e.Logger.Fatal(e.Start(":8080"))
}
