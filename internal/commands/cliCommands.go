package commands

import (
	"fmt"
	"errors"
	"os"
	"github.com/Bention99/pokedexcli/internal/api"
	"github.com/Bention99/pokedexcli/internal/pokecache"
)

type Config struct {
	PokeCache pokecache.Cache
	nextURL *string
	previousURL *string
	Arg string
}

type cliCommand struct {
	name string
	description string
	Callback func(*Config) error
}

var CliCommands = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		Callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Help for Pokedex commands",
		Callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		Callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the names of the previous 20 location areas in the Pokemon world",
		Callback: commandMapB,
	},
	"explore": {
		name: "explore",
		description: "Lists found Pokemons in Location",
		Callback: commandExplore,
	},
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("exit requested")
}

func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return errors.New("help requested")
}

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.nextURL != nil {
		url = *c.nextURL
	}
	apiResponse, err := api.GetLocationAreas(c.PokeCache, url)
	if err != nil {
		fmt.Printf("Error in API handling: %v", err)
		os.Exit(0)
		return errors.New("exit due to API problem")
	}

	for _, r := range apiResponse.Results {
		fmt.Println(r.Name)
	}
	c.nextURL = apiResponse.Next
	c.previousURL = apiResponse.Previous
	return errors.New("map requested")
}

func commandMapB(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.previousURL != nil {
		url = *c.previousURL
	}
	apiResponse, err := api.GetLocationAreas(c.PokeCache, url)
	if err != nil {
		fmt.Printf("Error in API handling: %v", err)
		os.Exit(0)
		return errors.New("exit due to API problem")
	}
	if apiResponse.Previous == nil {
		fmt.Println("you're on the first page")
		return errors.New("mapb requested - but currently on first page")
	}
	for _, r := range apiResponse.Results {
		fmt.Println(r.Name)
	}
	c.previousURL = apiResponse.Previous
	return errors.New("mapb requested")
}

func commandExplore(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	apiResponse, err := api.GetPokemonInLocation(url, c.Arg)
	if err != nil {
		fmt.Printf("Error in API handling: %v", err)
		os.Exit(0)
		return errors.New("exit due to API problem")
	}
	fmt.Printf("Exploring %s ...\n", apiResponse.Name)
	fmt.Println("Found Pokemon:")
	for _, p := range apiResponse.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}