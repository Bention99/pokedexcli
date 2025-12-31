package api

import (
	"fmt"
	"io"
	"net/http"
	"errors"
	"github.com/Bention99/pokedexcli/internal/pokecache"
)

type locationAreaList struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []locationArea  `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(c pokecache.Cache, url string) (locationAreaList, error) {
	val, ok := c.Get(url)
	if ok {
		v, err := unmarshal[locationAreaList](val)
		if err != nil {
			return locationAreaList{}, errors.New("Error Unmarshaling body")
		}
		return v, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return locationAreaList{}, errors.New("Error: Calling API failed\n")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return locationAreaList{}, errors.New("Error: Bad response\n")
	}
	c.Add(url, body)
	apiResponse, err := unmarshal[locationAreaList](body)
	if err != nil {
		return locationAreaList{}, errors.New("Error: Unmarshaling body\n")
	}
	return apiResponse, nil
}