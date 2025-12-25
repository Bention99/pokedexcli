package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		splittedInput := strings.Fields(input)
		fmt.Printf("Your command was: %s\n", splittedInput[0])
	}
}