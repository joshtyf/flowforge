package helper

import (
	"regexp"
)

func ExtractParameterFromString(s string) []string {
	re := regexp.MustCompile(`\${(.+?)}`)
	matches := re.FindAllStringSubmatch(s, -1)

	var words []string
	for _, match := range matches {
		if len(match) == 2 {
			words = append(words, match[1])
		}
	}

	return words
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
