package ui

import (
	"net/http"
	"strconv"

	"github.com/DonilZ/moviefan-rest-service/model"
	"github.com/gin-gonic/gin"
)

// GetFilms godoc
// @Summary Retrieves all films added by users
// @Produce json
// @Success 200 array model.Film
// @Failure 500 {object} model.DefaultResponse "Database error"
// @Router /films [get]
func GetFilms(m *model.Model) gin.HandlerFunc {
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

// GetFilm godoc
// @Summary Retrieves film based on given ID
// @Produce json
// @Param id path integer true "Film ID"
// @Success 200 {object} model.Film
// @Failure 400 {object} model.DefaultResponse "Invalid film ID"
// @Failure 404 {object} model.DefaultResponse "Film with specified ID not found"
// @Failure 500 {object} model.DefaultResponse "Database error"
// @Router /films/{id} [get]
func GetFilm(m *model.Model) gin.HandlerFunc {
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
