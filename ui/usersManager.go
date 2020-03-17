package ui

import (
	"fmt"
	"net/http"
	"webApp/src/model"

	"github.com/gin-gonic/gin"
)

func registerHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newUser model.User

		if !tryBindJSON(c, &newUser) {
			return
		}

		enteredFirstName := newUser.FirstName
		enteredEmail := newUser.Email
		enteredLogin := newUser.Login
		enteredPassword := newUser.Password

		if isEmptyField(c, &enteredFirstName, "firstname") ||
			isEmptyField(c, &enteredEmail, "e-mail") ||
			isEmptyField(c, &enteredLogin, "login") ||
			isEmptyField(c, &enteredPassword, "password") {
			return
		}

		allUsers, err := m.Users()

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		if isTheEnteredNewUserLoginOrEmailExist(c, allUsers, &enteredLogin, &enteredEmail) {
			return
		}

		if err = m.AddUser(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK,
			jsonResponse(http.StatusOK, "Registration completed successfully"))
	}
}

func isTheEnteredNewUserLoginOrEmailExist(c *gin.Context, allUsers []*model.User, enteredLogin, enteredEmail *string) bool {

	enteredLoginValue := *enteredLogin
	enteredEmailValue := *enteredEmail

	for _, existingUser := range allUsers {

		if existingUser.Login == enteredLoginValue {
			c.JSON(http.StatusConflict,
				jsonResponse(http.StatusConflict,
					fmt.Sprintf("User with %s login is already registered", enteredLoginValue)))
			return true
		}

		if existingUser.Email == enteredEmailValue {
			c.JSON(http.StatusConflict,
				jsonResponse(http.StatusConflict,
					fmt.Sprintf("User with %s e-mail is already registered", enteredEmailValue)))
			return true
		}
	}

	return false
}

func getUsersHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		userInfos := make([]*model.UserInfo, 0)
		users, err := m.Users()

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		for _, user := range users {
			userInfos = append(userInfos,
				&model.UserInfo{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Login:     user.Login,
					Email:     user.Email})
		}

		c.JSON(http.StatusOK, &userInfos)
	}
}

func getUserHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogin := c.Param("name")

		desiredUser, err := m.GetUserByLogin(userLogin)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		if desiredUser.Login == "" {
			c.JSON(http.StatusNotFound,
				jsonResponse(http.StatusNotFound, "User not found"))
			return
		}

		userInfo := model.UserInfo{
			FirstName: desiredUser.FirstName,
			LastName:  desiredUser.LastName,
			Login:     desiredUser.Login,
			Email:     desiredUser.Email}

		c.JSON(http.StatusOK, &userInfo)
	}
}
