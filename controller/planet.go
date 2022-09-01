package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mthcsta/star-wars-api-test/dao"
	"github.com/mthcsta/star-wars-api-test/httputil"
	"github.com/mthcsta/star-wars-api-test/model"
	"github.com/mthcsta/star-wars-api-test/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetController struct {
	planetDAO     dao.PlanetDAO
	planetService service.PlanetService
	filmService   service.FilmService
}

// Insert godoc
// @Summary Insert a new planet
// @Description Insert a new planet
// @Tags planet
// @Accept  json
// @Produce  json
// Param user body models.Planet true "Add Movie"
// Failure 400 {object} httputil.Error
// @Success 201 {object} model.Planet.Id "a"
// @Router /planets [post]
func (p *PlanetController) Insert(ctx *gin.Context) {
	var planet model.AddPlanet
	var warning string
	err := ctx.BindJSON(&planet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Error{
			Status:  http.StatusBadRequest,
			Message: "Invalid send payload",
		})
		return
	}

	// if planet name has already been registered
	if p.planetDAO.Exists(ctx, planet.Name) {
		ctx.JSON(http.StatusConflict, httputil.Error{
			Status:  http.StatusConflict,
			Message: "This planet has already been registered",
		})
		return
	}

	films, err := p.planetService.GetFilmsIdByPlanetName(planet.Name)

	if err != nil {
		warning = err.Error()
	} else {
		planet.Films = p.filmService.GetFilmsObjectIDById(films)
	}

	result := p.planetDAO.InsertOne(ctx, planet)

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"InsertedID": result.InsertedID,
		"warning":    warning,
	})
}

// GetAll godoc
// @Summary Get all planets recorded
// @Description Get all planets recorded
// @Tags planet
// @Accept  json
// @Produce  json
// @Failure 404
// @Success 200 {object} model.Planet
// @Router /planets [get]
func (p *PlanetController) GetAll(ctx *gin.Context) {
	var filterPlanet model.FilterPlanet
	err := ctx.ShouldBindQuery(&filterPlanet)
	if err != nil {
		fmt.Println(err)
	}
	if ctx.Query("id") != "" {
		filterPlanet.Id, _ = primitive.ObjectIDFromHex(ctx.Query("id"))
	}
	planets := p.planetDAO.GetAll(ctx, filterPlanet)
	ctx.JSON(http.StatusOK, planets)
}

// Remove godoc
// @Summary Remove a planet by ObjectID
// @Description Remove a planet by ObjectID
// @Tags planet
// @Accept json
// @Produce json
// @Param id path string true "Planet Object ID to remove"
// @Failure 401 ObjectID invalido
// @Success 204 Registro removido
// @Router /planets/{id} [delete]
func (p *PlanetController) Remove(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, httputil.Error{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	p.planetDAO.Remove(ctx, objectId)
	ctx.AbortWithStatus(http.StatusNoContent)
}
