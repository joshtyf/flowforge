package helper

import (
	"testing"
)

func TestExtractParameterFromString(t *testing.T) {
	testCases := []struct {
		input         string
		expectedWords []string
	}{
		{"my name is ${username}", []string{"username"}},
		{"This is a ${test} string with ${multiple} placeholders", []string{"test", "multiple"}},
		{"NoBracesHere", []string{}},
		{"${}", []string{}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			extractedWords := ExtractParameterFromString(tc.input)
			if !StringSliceEqual(extractedWords, tc.expectedWords) {
				t.Errorf("Expected: %v (len: %d), Got: %v (len: %d)", tc.expectedWords, len(tc.expectedWords),
					extractedWords, len(extractedWords))
			}
		})
	}
}

func TestStringSliceEqual(t *testing.T) {
	testCases := []struct {
		a        []string
		b        []string
		expected bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "d", "c"}, false},
		{[]string{}, []string{}, true},
		{[]string{}, []string{"a"}, false},
		{[]string{"a"}, []string{}, false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if StringSliceEqual(tc.a, tc.b) != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, !tc.expected)
			}
		})
	}
}
