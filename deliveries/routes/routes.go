package routes

import (
	"github.com/labstack/echo/v4"

	ps "github.com/dakasakti/postingan/deliveries/controllers/post"
	pt "github.com/dakasakti/postingan/deliveries/controllers/posttype"
	uc "github.com/dakasakti/postingan/deliveries/controllers/user"
	"github.com/dakasakti/postingan/deliveries/middlewares"
)

func UserPath(e *echo.Echo, uc uc.UserController) {
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)
}

func PostTypePath(e *echo.Echo, pt pt.PostTypeController) {
	e.POST("types", pt.Register, middlewares.JWTSign())
	e.GET("types", pt.GetAll)
	e.GET("types/:id", pt.GetById)
	e.PUT("types/:id", pt.UpdateById, middlewares.JWTSign())
	e.DELETE("types/:id", pt.DeleteById, middlewares.JWTSign())
}

func PostPath(e *echo.Echo, ps ps.PostController) {
	e.POST("/posts", ps.Register, middlewares.JWTSign())
	e.GET("/posts", ps.GetAll)
	e.GET("/posts/:id", ps.GetById)
	e.PUT("/posts/:id", ps.UpdateById, middlewares.JWTSign())
	e.DELETE("/posts/:id", ps.DeleteById, middlewares.JWTSign())
}
