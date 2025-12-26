package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleanedInput := cleanInput(input)
		cmdName := cleanedInput[0]
		cmd, ok := cliCommands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}
}