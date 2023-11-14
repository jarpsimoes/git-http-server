package utils

import (
	"regexp"
)

type PathSecurityCheck struct {
}

// IsValidPath checks if the input string adheres to the specified rules.
func (psc *PathSecurityCheck) IsValidPath(input string) bool {
	// Rule 1: Do not allow more than a single "." character.
	if countDots := psc.countOccurrences(input, '.'); countDots > 1 {
		return false
	}

	if psc.containsDirectorySeparator(input) {
		return false
	}

	allowList := []*regexp.Regexp{
		regexp.MustCompile(`^[a-zA-Z0-9_-]+$`),                        // Alphanumeric characters, underscore, and hyphen.
		regexp.MustCompile(`^([a-zA-Z0-9_-]+\.){0,2}[a-zA-Z0-9_-]+$`), // Allow up to two dots in the filename.
	}

	for _, pattern := range allowList {
		if pattern.MatchString(input) {
			return true
		}
	}

	return false
}

// countOccurrences counts the occurrences of a specific character in a string.
func (psc *PathSecurityCheck) countOccurrences(s string, c byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			count++
		}
	}
	return count
}

// containsDirectorySeparator checks if the input string contains directory separators.
func (psc *PathSecurityCheck) containsDirectorySeparator(s string) bool {
	return psc.containsAny(s, []byte{'/', '\\'})
}

// containsAny checks if the input string contains any of the specified characters.
func (psc *PathSecurityCheck) containsAny(s string, chars []byte) bool {
	for i := 0; i < len(s); i++ {
		for _, c := range chars {
			if s[i] == c {
				return true
			}
		}
	}
	return false
}
