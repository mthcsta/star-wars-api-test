package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mthcsta/star-wars-api-test/model"
	"github.com/stretchr/testify/assert"
)

var (
	filmController FilmController = FilmController{}
)

func TestFilmGetAll(t *testing.T) {
	route := gin.Default()
	route.GET("/", filmController.GetAll)

	// test get all films records
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	response, _ := ioutil.ReadAll(w.Body)
	var results []model.Film
	err := json.Unmarshal(response, &results)
	assert.Equal(t, http.StatusOK, w.Code)
	if err != nil {
		t.Error(err)
	}

	if len(results) == 0 {
		t.Skip("Not found results, finish tests here.")
	}

	// test get all films records
	substr := "a"
	req, _ = http.NewRequest("GET", "/?search="+substr, nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	response, _ = ioutil.ReadAll(w.Body)
	var results2 []model.Film
	err = json.Unmarshal(response, &results2)
	assert.Equal(t, http.StatusOK, w.Code)
	if err != nil {
		t.Error(err)
	}
	letterExists := false
	for _, result := range results2 {
		if strings.Contains(result.Title, substr) {
			letterExists = true
			break
		}
	}
	if letterExists == false {
		t.Error("Not returned results with letter")
	}

}
