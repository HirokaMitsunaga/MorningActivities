package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITimelineController interface {
	GetAllTimelines(c echo.Context) error
	GetTimelineById(c echo.Context) error
	CreateTimeline(c echo.Context) error
	UpdateTimeline(c echo.Context) error
	DeleteTimeline(c echo.Context) error
}

type timelineController struct {
	tlu usecase.ITimelineUsecase
}

func NewTimelineController(tlu usecase.ITimelineUsecase) ITimelineController {
	return &timelineController{tlu}
}

func (tlc *timelineController) GetAllTimelines(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// userId := claims["user_id"]
	taskRes, err := tlc.tlu.GetAllTimelines()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tlc *timelineController) GetTimelineById(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// userId := claims["user_id"]
	id := c.Param("timelineID")
	//Atoiを使ってstring型→int型へ変換する
	timelineId, _ := strconv.Atoi(id)
	timelineRes, err := tlc.tlu.GetTimelineById(uint(timelineId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, timelineRes)
}

func (tlc *timelineController) CreateTimeline(c echo.Context) error {

	timeline := model.Timeline{}
	if err := c.Bind(&timeline); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	timelineRes, err := tlc.tlu.CreateTimeline(timeline)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, timelineRes)
}
func (tlc *timelineController) UpdateTimeline(c echo.Context) error {
	id := c.Param("timelineID")
	timelineID, _ := strconv.Atoi(id)

	timeline := model.Timeline{}
	if err := c.Bind(&timeline); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	timelineRes, err := tlc.tlu.UpdateTimeline(timeline, uint(timelineID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, timelineRes)
}

func (tlc *timelineController) DeleteTimeline(c echo.Context) error {
	id := c.Param("timelineID")
	timelineID, _ := strconv.Atoi(id)

	if err := tlc.tlu.DeleteTimeline(uint(timelineID)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
