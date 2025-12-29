package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
    "github.com/Bention99/pokedexcli/internal/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &commands.Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleanedInput := cleanInput(input)
		cmdName := cleanedInput[0]
		cmd, ok := commands.CliCommands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.Callback(cfg)
	}
}

func cleanInput(text string) []string {
	splittedText := strings.Fields(strings.ToLower(text))
	return splittedText
}