package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
    
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  test  3132  ",
			expected: []string{"test", "3132"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length doesn't match. ACTUAL: %d EXPECTED: %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match. ACTUAL: %v EXPECTED: %v", word, expectedWord)
			}
		}
	}
}