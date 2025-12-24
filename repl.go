package main

import (
	"strings"
)

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	splittedText := strings.Fields(strings.ToLower(text))
	return splittedText
}