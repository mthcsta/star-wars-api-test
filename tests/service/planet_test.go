package service

import (
	"testing"

	"github.com/mthcsta/star-wars-api-test/service"
)

var (
	planetService service.PlanetService = service.PlanetService{}
)

func TestGetFilmsIdByPlanetName(t *testing.T) {
	var err error
	_, err = planetService.GetFilmsIdByPlanetName("Alderaan")
	if err != nil {
		t.Error(err)
	}
	_, err = planetService.GetFilmsIdByPlanetName("a")
	if err == nil {
		t.Error(err)
	}
	_, err = planetService.GetFilmsIdByPlanetName("sajdsakas")
	if err == nil {
		t.Error(err)
	}
}
