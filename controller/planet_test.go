package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mthcsta/star-wars-api-test/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	planetController PlanetController = PlanetController{}
	planets                           = []model.AddPlanet{
		{
			Name:    "Tatooine",
			Climate: "arid",
			Terrain: "desert",
		},
		{
			Name:    "Alderaan",
			Climate: "temperate",
			Terrain: "grasslands, mountains",
		},
	}
)

func TestInsert(t *testing.T) {
	route := gin.Default()
	route.POST("/", planetController.Insert)
	jsonValue, err := json.Marshal(planets[0])
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAll(t *testing.T) {
	route := gin.Default()
	route.GET("/", planetController.GetAll)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	response, _ := ioutil.ReadAll(w.Body)
	var results []model.Planet
	err := json.Unmarshal(response, &results)
	assert.Equal(t, http.StatusOK, w.Code)
	if err != nil {
		t.Error(err)
	}
}

func TestRemove(t *testing.T) {
	route := gin.Default()
	route.DELETE("/:id", planetController.Remove)

	// test value bad format
	req, _ := http.NewRequest("DELETE", "/123", nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// test value nil
	req, _ = http.NewRequest("DELETE", "/123", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// test value OK
	objectId := primitive.NewObjectID().Hex()
	req, _ = http.NewRequest("DELETE", "/"+objectId, nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

}
