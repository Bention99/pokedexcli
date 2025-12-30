package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)
	
type LocationArea struct {
	ID                    int                    `json:"id"`
	Name                  string                 `json:"name"`
	GameIndex             int                    `json:"game_index"`
	Location              NamedAPIResource       `json:"location"`
	Names                 []LocalizedName        `json:"names"`
	EncounterMethodRates  []EncounterMethodRate  `json:"encounter_method_rates"`
	PokemonEncounters     []PokemonEncounter     `json:"pokemon_encounters"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocalizedName struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource       `json:"encounter_method"`
	VersionDetails  []EncounterVersionRate `json:"version_details"`
}

type EncounterVersionRate struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource      `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

type PokemonVersionDetail struct {
	MaxChance       int                     `json:"max_chance"`
	Version         NamedAPIResource        `json:"version"`
	EncounterDetails []EncounterDetail      `json:"encounter_details"`
}

type EncounterDetail struct {
	Chance          int                  `json:"chance"`
	MinLevel        int                  `json:"min_level"`
	MaxLevel        int                  `json:"max_level"`
	Method          NamedAPIResource     `json:"method"`
	ConditionValues []NamedAPIResource   `json:"condition_values"`
}

func GetPokemonInLocation(url, location string) (LocationArea, error) {
	res, err := http.Get(url + "/" + location)
	if err != nil {
		return LocationArea{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationArea{}, err
	}
	if err != nil {
		return LocationArea{}, err
	}
	response, err := unmarshal(body)
	if err != nil {
		fmt.Println("Error Unmarshaling body")
		return LocationArea{}, err
	}
	return response, nil
}

func unmarshal(b []byte) (LocationArea, error) {
	var response LocationArea
	err := json.Unmarshal(b, &response)
	if err != nil {
		return LocationArea{}, err
	}
	return response, nil 
}