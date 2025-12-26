package main

import (
	"fmt"
	"errors"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("exit requested")
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return errors.New("help requested")
}