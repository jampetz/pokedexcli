package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello", "world",
			},
		},
		{
			input: "hI JaKE!",
			expected: []string{
				"hi", "jake!",
			},
		},
	}

	for _, command := range cases {
		actual := cleanInput(command.input)
		if len(actual) != len(command.expected) {
			t.Errorf("Lengths are not equal: %v vs %v", len(actual), len(command.expected))
			continue
		}
		for i := range actual {
			actualCommand := actual[i]
			expectedCommand := command.expected[i]
			if actualCommand != expectedCommand {
				t.Errorf("Commands are not equal: %v vs %v", actualCommand, expectedCommand)
				continue
			}
		}
	}

}
