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

func (f *FilmController) GetAll(ctx *gin.Context) {
	search := ctx.Query("search")
	films := f.filmDAO.GetAll(ctx, search)
	ctx.JSON(http.StatusOK, films)
}
