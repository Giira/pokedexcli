package main

import "testing"

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
			input:    "normal spacing",
			expected: []string{"normal", "spacing"},
		},
		{
			input:    "           oh   ",
			expected: []string{"oh"},
		},
		{
			input:    "PLEASE Make loweR",
			expected: []string{"please", "make", "lower"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Error: Incorrect number of words: %s vs %s", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s is different from %s", word, expectedWord)
			}
		}
	}
}
