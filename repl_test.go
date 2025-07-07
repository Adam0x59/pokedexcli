package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  hello      ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HelLo WorlD",
			expected: []string{"hello", "world"},
		},
	}

	for c_num, c := range cases {
		actual := cleanInput(c.input)
		// Check length of actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Case %d: Length of response does not match expected length - Fail!", c_num)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Case %d: Word: %s, dos not match expected: %s", c_num, word, expectedWord)
			}
		}
		for i := range actual {
			for x := range actual[i] {
				// Check each letter in each word in slice of strings
				// if a character is within the range of capital letters
				// return a message and fail the test
				if 'A' <= actual[i][x] && actual[i][x] <= 'Z' {
					t.Errorf("%s from word at index %d should be lowercase!", string(actual[i][x]), i)
				}
			}
		}
	}
}
