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
	err := getLocationAreas(c)
	if err != nil {
		fmt.Printf("Error calling API: %v", err)
		os.Exit(0)
		return errors.New("exit due to API error")
	}
	return errors.New("map requested")
}