package ui

import (
	"net/http"
	"strconv"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"
)

func getFilmsHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		films, err := m.Films()

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, &films)
	}
}

func getFilmHandler(m *model.Model) gin.HandlerFunc {
	return func(c *gin.Context) {

		filmID, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				jsonResponse(http.StatusBadRequest, err.Error()))
			return
		}

		desiredFilm, err := m.GetFilmByID(filmID)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				jsonResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		if desiredFilm.Name == "" {
			c.JSON(
				http.StatusNotFound,
				jsonResponse(http.StatusNotFound, "Film not found"))
			return
		}

		c.JSON(http.StatusOK, &desiredFilm)
	}
}
