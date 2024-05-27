package helper

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

var (
	ErrPlaceholderNotReplaced               = errors.New("some placeholders were not replaced")
	ErrInvalidTypeForPlaceholderReplacement = errors.New("placeholder replacement is only supported for strings, slices, and maps")
)

func ReplacePlaceholdersInString(input string, values map[string]any) (string, error) {
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
		} else if valueType.Kind() == reflect.Float64 {
			// If the value is a float, convert it to string and replace placeholder with it
			return strconv.FormatFloat(value.(float64), 'f', -1, 64)
		} else if valueType.Kind() == reflect.Bool {
			// If the value is a boolean, convert it to string and replace placeholder with it
			return strconv.FormatBool(value.(bool))
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

func ReplacePlaceholders(input any, values map[string]any) (any, error) {
	switch reflect.TypeOf(input).Kind() {
	case reflect.String:
		return ReplacePlaceholdersInString(input.(string), values)
	case reflect.Slice:
		// If the input is a slice, iterate over each element and replace placeholders
		output := make([]any, 0)
		for _, elem := range input.([]any) {
			replaced, err := ReplacePlaceholders(elem, values)
			if err != nil {
				return nil, err
			}
			output = append(output, replaced)
		}
		return output, nil
	case reflect.Map:
		// If the input is a map, iterate over each key and value and replace placeholders
		output := make(map[string]any)
		for key, value := range input.(map[string]any) {
			replacedKey, err := ReplacePlaceholdersInString(key, values)
			if err != nil {
				return nil, err
			}
			replacedValue, err := ReplacePlaceholders(value, values)
			if err != nil {
				return nil, err
			}
			output[replacedKey] = replacedValue
		}
		return output, nil
	default:
		return nil, ErrInvalidTypeForPlaceholderReplacement
	}
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
