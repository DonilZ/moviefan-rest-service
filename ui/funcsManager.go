package ui

import (
	"fmt"
	"net/http"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"
)

var funcs map[string]func(*model.Model, *gin.Context)

func defaultFuncs(m *model.Model) {
	funcs = make(map[string]func(*model.Model, *gin.Context))

	funcs["login"] = login
	funcs["logout"] = logout
}

func login(m *model.Model, c *gin.Context) {

	var knockingUser model.User

	if !tryBindJSON(c, &knockingUser) {
		return
	}

	enteredLogin := knockingUser.Login
	enteredPassword := knockingUser.Password

	desiredUser, err := m.GetUserByLogin(enteredLogin)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			jsonResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	if desiredUser.Login == "" {
		c.JSON(http.StatusUnauthorized,
			jsonResponse(http.StatusUnauthorized,
				fmt.Sprintf("User with login %s not found!", enteredLogin)))
		return
	}

	if desiredUser.Password != enteredPassword {
		c.JSON(http.StatusUnauthorized,
			jsonResponse(http.StatusUnauthorized,
				"Wrong password!"))
		return
	}

	setSession(desiredUser.Login, c)

	c.JSON(http.StatusOK,
		jsonResponse(http.StatusOK,
			fmt.Sprintf("User %s logged in successfully!", enteredLogin)))
}

func logout(m *model.Model, c *gin.Context) {
	clearSession(c)
	c.JSON(http.StatusOK, jsonResponse(http.StatusOK, "Success logout"))
}

// GetFuncs godoc
// @Summary Retrieves all functions
// @Produce json
// @Success 200 array string
// @Router /funcs [get]
func GetFuncs(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		var funcNames []string

		for key := range funcs {
			funcNames = append(funcNames, key)
		}

		c.JSON(http.StatusOK, &funcNames)
	}
}

// GetFunc godoc
// @Summary Call function based on given funcName
// @Param funcName path string true "Function name"
// @Success 200 {object} model.DefaultResponse "Function successfully called"
// @Failure 404 {object} model.DefaultResponse "Function with specified funcName not found"
// @Router /funcs/{funcName} [put]
func GetFunc(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		funcName := c.Param("funcName")

		if funcs[funcName] == nil {
			c.JSON(http.StatusNotFound,
				jsonResponse(http.StatusNotFound, "Function not found"))
			return
		}

		funcs[funcName](m, c)
	}
}
