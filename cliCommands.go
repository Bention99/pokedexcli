package main

import (
	"fmt"
	"errors"
	"os"
)

type config struct {
	nextURL *string
	previousURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

var cliCommands = map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Help for Pokedex commands",
		callback: commandHelp,
	},
	"map": {
		name: "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays the names of the previous 20 location areas in the Pokemon world",
		callback: commandMapB,
	},
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("exit requested")
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return errors.New("help requested")
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" 
	if c.nextURL != nil {
		url = *c.nextURL
	}
	apiResponse, err := getLocationAreas(url)
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

func commandMapB(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/" 
	if c.previousURL != nil {
		url = *c.previousURL
	}
	apiResponse, err := getLocationAreas(url)
	if err != nil {
		fmt.Printf("Error in API handling: %v", err)
		os.Exit(0)
		return errors.New("exit due to API problem")
	}
	for _, r := range apiResponse.Results {
		fmt.Println(r.Name)
	}
	c.previousURL = apiResponse.Previous
	return errors.New("map requested")
}