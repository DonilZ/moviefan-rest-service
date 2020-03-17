package ui

import (
	"net"
	"net/http"
	"webApp/src/model"

	"github.com/gin-gonic/gin"
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
		v1.GET("/funcs", getFuncsHandler(m))
		v1.PUT("/funcs/:funcName", getFuncHandler(m))

		v1.GET("/users", getUsersHandler(m))
		v1.GET("/users/:name", getUserHandler(m))
		v1.POST("/users", registerHandler(m))

		v1.GET("/userFilms", getUserFilmsHandler(m))
		v1.POST("/userFilms", addUserFilmHandler(m))
		v1.DELETE("/userFilms", deleteUserFilmHandler(m))

		v1.GET("/films", getFilmsHandler(m))
		v1.GET("/films/:id", getFilmHandler(m))
	}

	//Attach the router to http.Server and start listening HTTP requests
	router.RunListener(listener)
}
