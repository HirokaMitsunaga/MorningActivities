package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	GetCommentsHandler(c echo.Context) error
	GetAllComments(c echo.Context) error
	GetCommentById(c echo.Context) error
	// GetCommentsByTimelineId(c echo.Context, timelineId uint) error
	CreateComment(c echo.Context) error
	UpdateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}

type commentController struct {
	cu usecase.ICommentUsecase
}

func NewCommentController(cu usecase.ICommentUsecase) ICommentController {
	return &commentController{cu}
}

func (cc *commentController) GetCommentsHandler(c echo.Context) error {
	timelineIdParam := c.QueryParam("timelineID")
	if timelineIdParam != "" {
		return cc.GetCommentsByTimelineId(c)
	}
	return cc.GetAllComments(c)
}

func (cc *commentController) GetAllComments(c echo.Context) error {
	//JWTのuser_idを取り込む
	//JWTをデコードして値を取ってくる
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	commentRes, err := cc.cu.GetAllComments(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentRes)
}

func (cc *commentController) GetCommentById(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// userId := claims["user_id"]
	id := c.Param("commentID")
	//Atoiを使ってstring型→int型へ変換する
	commentId, _ := strconv.Atoi(id)
	commentRes, err := cc.cu.GetCommentById(uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentRes)
}

func (cc *commentController) GetCommentsByTimelineId(c echo.Context) error {
	timelineIdParam := c.QueryParam("timelineID")
	timelineId, err := strconv.Atoi(timelineIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid timelineID")
	}
	commentRes, err := cc.cu.GetCommentsByTimelineId(uint(timelineId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentRes)
}

func (cc *commentController) CreateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	comment.UserId = uint(userId.(float64))
	commentRes, err := cc.cu.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, commentRes)
}
func (cc *commentController) UpdateComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("commentID")
	commentId, _ := strconv.Atoi(id)

	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	comment.UserId = uint(userId.(float64))
	comment.ID = uint(commentId)
	commentRes, err := cc.cu.UpdateComment(comment, uint(userId.(float64)), uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentRes)
}

func (cc *commentController) DeleteComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("commentID")
	commentId, _ := strconv.Atoi(id)
	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	comment.ID = uint(commentId)
	comment.UserId = uint(userId.(float64))

	if err := cc.cu.DeleteComment(&comment, uint(userId.(float64))); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
