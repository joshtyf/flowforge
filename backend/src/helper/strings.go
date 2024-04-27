package helper

import (
	"errors"
	"regexp"
)

var (
	ErrPlaceholderNotReplaced = errors.New("some placeholders were not replaced")
)

func ReplacePlaceholders(input string, values map[string]string) (string, error) {
	// Regular expression to find placeholders
	re := regexp.MustCompile(`\$\{(.*?)\}`)

	// Replace each placeholder found in the string
	replaced := re.ReplaceAllStringFunc(input, func(match string) string {
		// Strip '${' prefix and '}' suffix
		key := match[2 : len(match)-1]

		// Retrieve value from the map
		value, exists := values[key]
		if !exists {
			// If the key doesn't exist, return an error
			return match // return the original placeholder
		}

		// Replace placeholder with value from the map
		return value
	})

	// Check if there are any leftover placeholders
	leftoverPlaceholders := re.FindString(replaced)
	if leftoverPlaceholders != "" {
		// If there are leftover placeholders, return an error
		return "", ErrPlaceholderNotReplaced
	}

	return replaced, nil
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func StringInSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
