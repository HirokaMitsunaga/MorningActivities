package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ILikeController interface {
	CreateLike(c echo.Context) error
	DeleteLike(c echo.Context) error
}

type likeController struct {
	lu usecase.ILikeUsecase
}

func NewLikeController(lu usecase.ILikeUsecase) ILikeController {
	return &likeController{lu}
}

func (lc *likeController) CreateLike(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64) // JWTからuserIdを取得
	like := model.Like{}
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	like.UserId = uint(userId) // UserIdをTimelineオブジェクトに設定
	likeRes, err := lc.lu.CreateLike(like)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, likeRes)
}

func (lc *likeController) DeleteLike(c echo.Context) error {
	user, _ := c.Get("user").(*jwt.Token)
	id := c.Param("likeID")
	likeID, _ := strconv.Atoi(id)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64) // JWTからuserIdを取得
	like := model.Like{}
	if err := c.Bind(&like); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	like.UserId = uint(userId) // UserIdをTimelineオブジェクトに設定
	like.ID = uint(likeID)
	err := lc.lu.DeleteLike(like)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
