package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mthcsta/star-wars-api-test/dao"
	"github.com/mthcsta/star-wars-api-test/service"
)

type FilmController struct {
	filmDAO     dao.FilmDAO
	filmService service.FilmService
}

// GetAll godoc
// @Summary Get all films recorded
// @Description Get all films recorded
// @Tags film
// @Accept  json
// @Produce  json
// @Param search query string false "Search film by title"
// @Success 200 {object} model.Film
// @Router /films [get]
func (f *FilmController) GetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	films := f.filmDAO.GetAll(ctx, search)
	ctx.JSON(http.StatusOK, films)
}
