package main


import (
	"fmt"
	"io"
	"net/http"
)

func getLocationAreas(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" 
	if c.previousURL != nil {
		url = url + *c.previousURL + "/"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return err
	}
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", body)
	return nil
}