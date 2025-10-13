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
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander BULbaSAUR PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("In TestCleanInput: length of result - %d does not match expected - %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			if word != actual[i] {
				t.Errorf("In TestCleanInput: actual word - %s does not match expected word - %s", actual[i], word)
			}
		}
	}
}
