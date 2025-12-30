package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
    "github.com/Bention99/pokedexcli/internal/commands"
	"github.com/Bention99/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &commands.Config{
		PokeCache: pokeCache,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleanedInput := cleanInput(input)
		cmdName := cleanedInput[0]
		if len(cleanedInput) > 1 {
			cfg.Arg = cleanedInput[1]
		}
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