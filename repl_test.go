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

func TestIdExtraction(t *testing.T) {
	cases := []struct {
		input    string
		expected struct {
			value int
			err   error
		}
	}{
		{
			input: "https://pokeapi.co/api/v2/location-area/13/",
			expected: struct {
				value int
				err   error
			}{13, nil},
		},
		{
			input: "https://pokeapi.co/api/v2/location-area/12345/",
			expected: struct {
				value int
				err   error
			}{12345, nil},
		},
		{
			input: "https://pokeapi.co/api/v2/location-area/id1234/",
			expected: struct {
				value int
				err   error
			}{0, fmt.Errorf("error")},
		},
		{
			input: "https://pokeapi.co/api/v2/location-area/abcd/",
			expected: struct {
				value int
				err   error
			}{0, fmt.Errorf("error")},
		},
	}

	for _, id := range cases {
		actual, err := getIDFromURL(id.input)
		if err != nil && id.expected.err == nil {
			t.Errorf("ID extraction error: %v", err)
			continue
		}

		if actual != id.expected.value {
			t.Errorf("ID mismatch: %v vs %v", actual, id.expected.value)
			continue
		}
	}
}
