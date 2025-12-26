package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splittedText := strings.Fields(strings.ToLower(text))
	return splittedText
}