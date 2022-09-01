package service

import (
	"testing"
)

var (
	planetService PlanetService = PlanetService{}
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
