package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/router"
	"go-api/usecase"
	"go-api/validator"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// ログのフォーマットを設定
	log.Formatter = &logrus.JSONFormatter{}
	// ログレベルの設定
	log.Level = logrus.DebugLevel
	// ログの出力先を設定
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}
func main() {
	db := db.NewDB()
	userValidator := validator.NewuserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	timelineRepository := repository.NewTimelineRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	likeRepository := repository.NewLikeRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	timelineUsecase := usecase.NewTimelineUsecase(timelineRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	likeUsecase := usecase.NewLikeUsecase(likeRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	timelineController := controller.NewTimelineController(timelineUsecase)
	commentController := controller.NewCommentController(commentUsecase)
	likeController := controller.NewLikeController(likeUsecase)
	e := router.NewRouter(userController, taskController, timelineController, likeController, commentController)
	//エラーが出たらLoggerがエラー出力して、強制終了する
	e.Logger.Fatal(e.Start(":8080"))
}
