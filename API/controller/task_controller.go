package controller

import (
	"MorningActivities-API/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)
 
type ITaskController interface{
	GetAllTask(c echo.Context)error
	GetTaskById(c echo.Context)error
	CreateTask(c echo.Context)error
	UpdateTask(c echo.Context)error
	DeleteTask(c echo.Context)error
}

type taskController struct{
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase)ITaskController{
	return &taskController{tu}
}

func(tc *taskController)GetAllTasks(c echo.Context) error{
	// taskRes :=[]model.TaskResponse{} //GetAllTasksでTaskResponse{}が生成されるため不要
	user := c.Get("user").(*jwt.Token)
	claims :=user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	
	taskRes, err :=tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
} 

func(tc *taskController)GetTaskById(c echo.Context)error{
	user := c.Get("user").(*jwt.Token)
	claims :=user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)),uint(taskId))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}