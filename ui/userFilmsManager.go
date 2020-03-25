package ui

import (
	"net/http"
	"time"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"
)

func getUserFilmsHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, isAuthorized := isTheUserAuthorized(c)
		if !isAuthorized {
			return
		}

		enteredUserName := c.Param("name")

		currentUser, err := m.GetUserByLogin(enteredUserName)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		userFilms, err := m.UserFilms(currentUser.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, &userFilms)
	}
}

func addUserFilmHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUserName, isAuthorized := isTheUserAuthorized(c)
		if !isAuthorized {
			return
		}

		enteredUserName := c.Param("name")

		if enteredUserName != currentUserName {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusForbidden, "Not enough rights"))
			return
		}

		currentUser, _ := m.GetUserByLogin(currentUserName)

		var newUserFilm model.Film
		if !tryBindJSON(c, &newUserFilm) {
			return
		}

		if isEmptyField(c, &newUserFilm.Name, "FilmName") {
			return
		}

		if newUserFilm.Year < 1895 || newUserFilm.Year > time.Now().Year() {
			c.JSON(http.StatusBadRequest,
				jsonResponse(http.StatusBadRequest,
					"Incorrect year entered!"))
			return
		}

		existingFilmID, err := m.GetFilmIDByNameAndYear(&newUserFilm.Name, newUserFilm.Year)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		IsTheNewFilmOnTheListOfFilms := existingFilmID > -1

		if !IsTheNewFilmOnTheListOfFilms {
			m.AddFilm(&newUserFilm)

			addedFilmID, err := m.GetFilmIDByNameAndYear(&newUserFilm.Name, newUserFilm.Year)

			if err != nil {
				c.JSON(http.StatusInternalServerError,
					jsonResponse(http.StatusInternalServerError, err.Error()))
				return
			}

			m.AddUserFilm(currentUser.ID, addedFilmID)

		} else {
			userFilms, err := m.UserFilms(currentUser.ID)

			if err != nil {
				c.JSON(http.StatusInternalServerError,
					jsonResponse(http.StatusInternalServerError, err.Error()))
				return
			}

			for _, userFilm := range userFilms {
				if userFilm.ID == existingFilmID {
					c.JSON(http.StatusConflict,
						jsonResponse(http.StatusConflict, "Film already added"))
					return
				}
			}

			m.AddUserFilm(currentUser.ID, existingFilmID)
		}

		c.JSON(http.StatusOK,
			jsonResponse(http.StatusOK, "Film successfully added"))
	}
}

func deleteUserFilmHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		currentUserName, isAuthorized := isTheUserAuthorized(c)
		if !isAuthorized {
			return
		}

		enteredUserName := c.Param("name")

		if enteredUserName != currentUserName {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusForbidden, "Not enough rights"))
			return
		}

		currentUser, err := m.GetUserByLogin(currentUserName)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		var removableFilm model.Film

		if !tryBindJSON(c, &removableFilm) {
			return
		}

		userFilms, err := m.UserFilms(currentUser.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		for _, userFilm := range userFilms {
			if userFilm.ID == removableFilm.ID {
				m.DeleteUserFilm(currentUser.ID, removableFilm.ID)
				c.JSON(http.StatusOK,
					jsonResponse(http.StatusOK, "Film successfully deleted"))
				return
			}
		}

		c.JSON(http.StatusNotFound,
			jsonResponse(http.StatusNotFound, "Removable film not found"))
	}
}
