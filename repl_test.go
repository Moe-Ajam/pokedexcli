package main

import (
	"testing"
)

func TestCLeanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(actual),
				len(cs.expected))
			continue
		}
		for i := range actual {
			actualWorld := actual[i]
			expectedWorld := cs.expected[i]
			if actualWorld != expectedWorld {
				t.Errorf("Mismatched word: %s vs %s", actualWorld, expectedWorld)
			}
		}
	}
}
