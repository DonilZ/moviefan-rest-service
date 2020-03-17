package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(userName string, c *gin.Context) string {
	value := map[string]string{
		"name": userName,
	}

	if encodedValue, err := cookieHandler.Encode("session", value); err == nil {

		c.SetCookie("session", encodedValue, 3600, "/", "localhost",
			http.SameSiteLaxMode, false, true)

		return encodedValue
	}

	return ""
}

func getCurrentSessionUserName(c *gin.Context) (userName string) {

	if cookie, err := c.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)

		if err = cookieHandler.Decode("session", cookie, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}

	return userName
}

func clearSession(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", http.SameSiteLaxMode, false, true)
}

func isTheUserAuthorized(c *gin.Context) (string, bool) {
	currentUserName := getCurrentSessionUserName(c)

	if currentUserName == "" {
		c.JSON(http.StatusUnauthorized,
			jsonResponse(http.StatusUnauthorized, "You must be logged in to perform this action"))
		return currentUserName, false
	}

	return currentUserName, true
}
