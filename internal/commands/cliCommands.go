package commands

import (
	"fmt"
	"errors"
	"os"
	"math/rand"
	"github.com/Bention99/pokedexcli/internal/api"
	"github.com/Bention99/pokedexcli/internal/pokecache"
	"github.com/Bention99/pokedexcli/internal/pokedexCatches"
)

type Config struct {
	PokeCache pokecache.Cache
	Rand *rand.Rand
	nextURL *string
	previousURL *string
	Arg string
	Caught map[string]api.Pokemon
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
	"catch": {
		name: "catch",
		description: "Try to catch a Pokemon",
		Callback: commandCatch,
	},
	"free": {
		name: "free",
		description: "Frees all caught Pokemon",
		Callback: commandFree,
	},
	"list": {
		name: "list",
		description: "Lists all caught Pokemon",
		Callback: commandList,
	},
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("exit requested")
}

func commandHelp(c *Config) error {
	fmt.Println(`Welcome to the Pokedex!
		Usage:

		help: Displays a help message
		exit: Exit the Pokedex
		map: Displays Locations where Pokemon could hide
		mapb: Goes back to the previous Locations
		explore: Lists found Pokemons in specified Location - Argument (Location Name) necessary
		catch: Try your luck catching a specified Pokemon - Argument (Pokemon Name) necessary
		`)
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
	if c.Arg == "" {
		return errors.New("Please provide Location Name")
	}
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

func commandCatch(c *Config) error {
	if c.Arg == "" {
		return errors.New("Please provide Pokemon Name")
	}
	_, ok := c.Caught[c.Arg]
	if ok {
		fmt.Printf("Pokemon already captured: %s\n", c.Arg)
		return errors.New("Pokemon already captured.")
	}
	url := "https://pokeapi.co/api/v2/pokemon/"
	fmt.Printf("Throwing a Pokeball at %s...", c.Arg)
	apiResponse, err := api.GetPokemonDetails(url, c.Arg)
	if err != nil {
		fmt.Printf("There is no Pokemon called %s\n", c.Arg)
		return errors.New("API problem")
	}
	if c.Rand.Intn(200) <= apiResponse.BaseExperience {
		fmt.Printf("%s escaped!\n", c.Arg)
		return nil
	}
	fmt.Printf("%s was caught!\n", c.Arg)
	c.Caught[apiResponse.Name] = apiResponse
	_ = pokedexCatches.SaveCaughtJSON("data/caught.json", c.Caught)
	return nil
}

func commandFree(c *Config) error {
	_ = pokedexCatches.DeleteCaughtFile("data/caught.json")
	fmt.Println("Pokemon released in the wild.")
	return nil
}

func commandList(c *Config) error {
	fmt.Println("You caught:")
	for _, p := range c.Caught {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}