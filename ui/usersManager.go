package ui

import (
	"fmt"
	"net/http"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Accept json
// @Produce json
// @Success 200 {object} model.DefaultResponse "Registration completed successfully"
// @Failure 400 {object} model.DefaultResponse "Invalid register data"
// @Failure 409 {object} model.DefaultResponse "User with such data is already registered"
// @Failure 500 {object} model.DefaultResponse "Database error"
// @Router /users [post]
func RegisterUser(m *model.Model) gin.HandlerFunc {
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

// GetUsers godoc
// @Summary Retrieves all registered users
// @Produce json
// @Success 200 array model.UserInfo
// @Failure 500 {object} model.DefaultResponse "Database error"
// @Router /users [get]
func GetUsers(m *model.Model) gin.HandlerFunc {
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

// GetUser godoc
// @Summary Retrieves user based on given UserName (Login)
// @Produce json
// @Param userName path string true "UserName (Login)"
// @Success 200 {object} model.UserInfo
// @Failure 404 {object} model.DefaultResponse "User not found"
// @Failure 500 {object} model.DefaultResponse "Database error"
// @Router /users/{userName} [get]
func GetUser(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		userLogin := c.Param("userName")

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
