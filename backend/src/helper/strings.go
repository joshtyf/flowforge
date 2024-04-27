package helper

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

var (
	ErrPlaceholderNotReplaced = errors.New("some placeholders were not replaced")
)

func ReplacePlaceholders(input string, values map[string]any) (string, error) {
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

		if valueType := reflect.TypeOf(value); valueType.Kind() == reflect.String {
			// If the value is a string, replace placeholder with it
			return value.(string)
		} else if valueType.Kind() == reflect.Int {
			// If the value is an integer, convert it to string and replace placeholder with it
			return strconv.Itoa(value.(int))
		} else {
			return match
		}
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
