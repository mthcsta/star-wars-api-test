package apis

import (
	"io"
	"log"
	"net/http"
)

type API struct {
	BaseUrl string
}

func (api *API) HandleRequest(path string) []byte {
	url := api.BaseUrl + path
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
