package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mthcsta/star-wars-api-test/apis"
)

type PlanetService struct{}

// GetFilmsIdByPlanetName returns slice of films where the planet appears
func (p *PlanetService) GetFilmsIdByPlanetName(name string) ([]int, error) {
	json := apis.SearchPlanet(name)
	var films []int

	switch {
	case json.Count == 1:
		fmt.Println(json.Results[0].Films)
		for _, film := range json.Results[0].Films {
			splitted := strings.Split(film, "/")
			splittedIndex := splitted[len(splitted)-2]
			id, err := strconv.Atoi(splittedIndex)
			if err != nil {
				log.Panic(err)
			}
			films = append(films, id)
		}
		break
	case json.Count >= 1:
		err := errors.New("More than one planet contains this name. No planet was selected from API")
		return films, err
		break
	default:
		err := errors.New("Planet Not Found")
		return films, err
		break
	}

	return films, nil
}
