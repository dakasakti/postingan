package main

import (
	"fmt"

	"github.com/dakasakti/postingan/config"

	"github.com/dakasakti/postingan/deliveries/middlewares"
	"github.com/dakasakti/postingan/deliveries/routes"

	pm "github.com/dakasakti/postingan/repositories/post"
	tm "github.com/dakasakti/postingan/repositories/posttype"
	um "github.com/dakasakti/postingan/repositories/user"

	ps "github.com/dakasakti/postingan/services/post"
	ts "github.com/dakasakti/postingan/services/posttype"
	us "github.com/dakasakti/postingan/services/user"
	"github.com/dakasakti/postingan/services/validation"

	pc "github.com/dakasakti/postingan/deliveries/controllers/post"
	tc "github.com/dakasakti/postingan/deliveries/controllers/posttype"
	uc "github.com/dakasakti/postingan/deliveries/controllers/user"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	middlewares.General(server)

	// database connection
	db := config.InitDB(*config.GetConfig())
	config.AutoMigrate(db)

	// models
	userModel := um.NewUserModel(db)
	typeModel := tm.NewPostTypeModel(db)
	postModel := pm.NewPostModel(db)

	// services
	validate := validation.NewValidation()
	userService := us.NewUserService(userModel)
	typeService := ts.NewPostTypeService(typeModel)
	postService := ps.NewPostService(postModel)

	// controllers
	userController := uc.NewUserController(userService, validate)
	typeController := tc.NewPostTypeController(typeService, validate)
	postController := pc.NewPostController(postService, validate)

	// routes
	routes.UserPath(server, userController)
	routes.PostTypePath(server, typeController)
	routes.PostPath(server, postController)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", config.GetConfig().Port)))
}
