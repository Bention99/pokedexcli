package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
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
		v, err := unmarshalBody(val)
		if err != nil {
			fmt.Println("Error Unmarshaling body")
			return locationAreaList{}, err
		}
		return v, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return locationAreaList{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return locationAreaList{}, err
	}
	if err != nil {
		return locationAreaList{}, err
	}
	c.Add(url, body)
	apiResponse, err := unmarshalBody(body)
	if err != nil {
		fmt.Println("Error Unmarshaling body")
		return locationAreaList{}, err
	}
	return apiResponse, nil
}

func unmarshalBody(b []byte) (locationAreaList, error) {
	var apiResponse locationAreaList
	err := json.Unmarshal(b, &apiResponse)
	if err != nil {
		return locationAreaList{}, err
	}
	return apiResponse, nil 
}