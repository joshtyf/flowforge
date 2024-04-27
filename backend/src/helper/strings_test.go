package helper

import (
	"testing"
)

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

func TestStringInSlice(t *testing.T) {
	testCases := []struct {
		s        string
		sl       []string
		expected bool
	}{
		{"a", []string{"a", "b", "c"}, true},
		{"b", []string{"a", "b", "c"}, true},
		{"c", []string{"a", "b", "c"}, true},
		{"d", []string{"a", "b", "c"}, false},
		{"", []string{"a", "b", "c"}, false},
		{"", []string{}, false},
		{"a", []string{}, false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if StringInSlice(tc.s, tc.sl) != tc.expected {
				t.Errorf("Expected: %v, Got: %v", tc.expected, !tc.expected)
			}
		})
	}
}

func TestReplacePlaceholders(t *testing.T) {
	testCases := []struct {
		input    string
		values   map[string]string
		expected string
		err      error
	}{
		{
			"Hello ${name}, you are ${age} years old",
			map[string]string{
				"name": "john",
				"age":  "50",
			},
			"Hello john, you are 50 years old",
			nil,
		},
		{
			"Hello ${{name}, you are ${age} years old",
			map[string]string{
				"name": "john",
				"age":  "50",
			},
			"",
			ErrPlaceholderNotReplaced,
		},
		{
			"Hello",
			map[string]string{},
			"Hello",
			nil,
		},
	}

	for _, tc := range testCases {
		replaced, err := ReplacePlaceholders(tc.input, tc.values)
		if tc.err != err {
			t.Errorf("Expected: %v, Got: %v", tc.err, err)
		}
		if replaced != tc.expected {
			t.Errorf("Expected: %v, Got: %v", tc.expected, replaced)
		}
	}
}
