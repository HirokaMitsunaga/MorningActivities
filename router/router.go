package router

import (
	"go-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, tlc controller.ITimelineController, lc controller.ILikeController, cc controller.ICommentController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	tl := e.Group("/timelines")
	tl.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	tl.GET("", tlc.GetAllTimelines)
	tl.GET("/:timelineID", tlc.GetTimelineById)
	tl.POST("", tlc.CreateTimeline)
	tl.PUT("/:timelineID", tlc.UpdateTimeline)
	tl.DELETE("/:timelineID", tlc.DeleteTimeline)

	//likeの設定
	l := e.Group("/likes")
	l.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	l.POST("", lc.CreateLike)
	l.DELETE("/:likeID", lc.DeleteLike)
	l.POST("/toggle", lc.ToggleLike)

	com := e.Group("/comments")
	com.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	com.GET("", cc.GetCommentsHandler)
	com.GET("/:commentID", cc.GetCommentById)
	com.POST("", cc.CreateComment)
	com.PUT("/:commentID", cc.UpdateComment)
	com.DELETE("/:commentID", cc.DeleteComment)
	return e

}
