package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello you there",
			expected: []string{"hello", "you", "there"},
		},
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "HELLO world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			fmt.Println("expected", c.expected, len(c.expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("Expected %v but Got %v", expectedWord, actualWord)
			}
		}
	}
}
