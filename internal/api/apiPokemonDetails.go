package api

import (
	"fmt"
	"io"
	"net/http"
	"errors"
)

type Pokemon struct {
	ID                     int               `json:"id"`
	Name                   string            `json:"name"`
	BaseExperience         int               `json:"base_experience"`
	Height                 int               `json:"height"`
	IsDefault              bool              `json:"is_default"`
	Order                  int               `json:"order"`
	Weight                 int               `json:"weight"`
	Abilities              []AbilityEntry    `json:"abilities"`
	Forms                  []NamedAPIR `json:"forms"`
	GameIndices            []GameIndex       `json:"game_indices"`
	HeldItems              []HeldItem        `json:"held_items"`
	LocationAreaEncounters string            `json:"location_area_encounters"`
	Moves                  []MoveEntry       `json:"moves"`
	Species                NamedAPIR  `json:"species"`
	Sprites                Sprites           `json:"sprites"`
	Cries                  Cries             `json:"cries"`
	Stats                  []StatEntry       `json:"stats"`
	Types                  []TypeEntry       `json:"types"`
	PastTypes              []PastType        `json:"past_types"`
	PastAbilities          []PastAbility     `json:"past_abilities"`
}

type NamedAPIR struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AbilityEntry struct {
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
	Ability  NamedAPIR  `json:"ability"`
}

type GameIndex struct {
	GameIndex int              `json:"game_index"`
	Version   NamedAPIR `json:"version"`
}

type HeldItem struct {
	Item           NamedAPIR      `json:"item"`
	VersionDetails []HeldItemVersionInfo `json:"version_details"`
}

type HeldItemVersionInfo struct {
	Rarity  int              `json:"rarity"`
	Version NamedAPIR `json:"version"`
}

type MoveEntry struct {
	Move                NamedAPIR      `json:"move"`
	VersionGroupDetails []MoveVersionDetail   `json:"version_group_details"`
}

type MoveVersionDetail struct {
	LevelLearnedAt  int              `json:"level_learned_at"`
	VersionGroup    NamedAPIR `json:"version_group"`
	MoveLearnMethod NamedAPIR `json:"move_learn_method"`
	Order           int              `json:"order"`
}

type Sprites struct {
	BackDefault       *string            `json:"back_default"`
	BackFemale        *string            `json:"back_female"`
	BackShiny         *string            `json:"back_shiny"`
	BackShinyFemale   *string            `json:"back_shiny_female"`
	FrontDefault      *string            `json:"front_default"`
	FrontFemale       *string            `json:"front_female"`
	FrontShiny        *string            `json:"front_shiny"`
	FrontShinyFemale  *string            `json:"front_shiny_female"`
	Other             SpriteOther        `json:"other"`
	Versions          SpriteVersions     `json:"versions"`
}

type SpriteOther struct {
	DreamWorld        SpriteSet `json:"dream_world"`
	Home              SpriteSet `json:"home"`
	OfficialArtwork   SpriteSet `json:"official-artwork"`
	Showdown          SpriteSet `json:"showdown"`
}

type SpriteSet struct {
	FrontDefault      *string `json:"front_default"`
	FrontFemale       *string `json:"front_female"`
	FrontShiny        *string `json:"front_shiny"`
	FrontShinyFemale  *string `json:"front_shiny_female"`
	BackDefault       *string `json:"back_default"`
	BackFemale        *string `json:"back_female"`
	BackShiny         *string `json:"back_shiny"`
	BackShinyFemale   *string `json:"back_shiny_female"`
}

type SpriteVersions map[string]map[string]any

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type StatEntry struct {
	BaseStat int              `json:"base_stat"`
	Effort   int              `json:"effort"`
	Stat     NamedAPIR `json:"stat"`
}

type TypeEntry struct {
	Slot int              `json:"slot"`
	Type NamedAPIR `json:"type"`
}

type PastType struct {
	Generation NamedAPIR `json:"generation"`
	Types      []TypeEntry      `json:"types"`
}

type PastAbility struct {
	Generation NamedAPIR        `json:"generation"`
	Abilities  []PastAbilityEntry      `json:"abilities"`
}

type PastAbilityEntry struct {
	Ability  *NamedAPIR `json:"ability"`
	IsHidden bool              `json:"is_hidden"`
	Slot     int               `json:"slot"`
}

func GetPokemonDetails(url, pokemon string) (Pokemon, error) {
	res, err := http.Get(url + pokemon + "/")
	if err != nil {
		return Pokemon{}, errors.New("Error: Calling API failed\n")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return Pokemon{}, errors.New("Error: Bad response\n")
	}
	response, err := unmarshal[Pokemon](body)
	if err != nil {
		return Pokemon{}, errors.New("Error: Unmarshaling body\n")
	}
	return response, nil
}