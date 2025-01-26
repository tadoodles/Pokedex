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
			input:    "PIKACHU    CHARIZARD",
        	expected: []string{"pikachu", "charizard"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input:    "PiKaChU",
			expected: []string{"pikachu"},
		},
		{
			input:    "Charmander    Bulbasaur      Squirtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		// add more cases here
	}
	

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Got %d words, expected %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
                t.Errorf("Got word %q at position %d, expected %q", word, i, expectedWord)
            }
		}
	}
}
	