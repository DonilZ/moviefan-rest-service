package ui

import (
	"net"
	"net/http"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Config ...
type Config struct {
	Assets http.FileSystem
}

//Start ...
func Start(cfg Config, m *model.Model, listener net.Listener) {

	/* Creates a gin router with default middleware:
	 * logger and recovery (crash-free) middleware */
	router := gin.Default()

	//Initialize the pseudo-collection of functions and add default functions (login, logout)
	defaultFuncs(m)

	//Create a group that unites all our api methods
	v1 := router.Group("api/v1")
	{
		v1.GET("/funcs", GetFuncs(m))
		v1.PUT("/funcs/:funcName", GetFunc(m))

		v1.GET("/users", GetUsers(m))
		v1.GET("/users/:userName", GetUser(m))
		v1.POST("/users", RegisterUser(m))

		v1.GET("/users/:userName/films", GetUserFilms(m))
		v1.POST("/users/:userName/films", AddUserFilm(m))
		v1.DELETE("/users/:userName/films", DeleteUserFilm(m))

		v1.GET("/films", GetFilms(m))
		v1.GET("/films/:id", GetFilm(m))

		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//Attach the router to http.Server and start listening HTTP requests
	router.RunListener(listener)
}
