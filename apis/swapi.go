package apis

import (
	"encoding/json"
	"net/url"
	"strconv"
)

var (
	api API = API{
		BaseUrl: "https://swapi.dev/api/",
	}
)

type PlanetResponse struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}

type PlanetPayloadResponse struct {
	Count   int              `json:"count"`
	Results []PlanetResponse `json:"results"`
}

type FilmPayloadResponse struct {
	Title       string `json:"title"`
	EpisodeId   int    `json:"episode_id"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
}

func GetPlanets() PlanetPayloadResponse {
	planetsByte := api.HandleRequest("/planets/")
	var payload PlanetPayloadResponse
	json.Unmarshal(planetsByte, &payload)
	return payload
}

func SearchPlanet(search string) PlanetPayloadResponse {
	planetsByte := api.HandleRequest("/planets/?search=" + url.QueryEscape(search))
	var payload PlanetPayloadResponse
	json.Unmarshal(planetsByte, &payload)
	return payload
}

func GetFilmById(id int) FilmPayloadResponse {
	filmByte := api.HandleRequest("/films/" + strconv.Itoa(id))
	var payload FilmPayloadResponse
	json.Unmarshal(filmByte, &payload)
	return payload
}
