// This package has custom validators that are used to validate field
package validation


import (
	"regexp"
	"unicode"
)


// Checks if `value` is present or not in `members`
func OneOf(value interface{}, members []interface{}) bool {
	for _, item := range members {
		if value == item {
			return true
		}
	}
	return false
}


// Checks if `value` is absent or not in `members`
func NoneOf(value interface{}, members []interface{}) bool {
	for _, item := range members {
		if value == item {
			return false
		}
	}
	return true
}


// Validated if a string contains at least one uppercase character or not
func HasUppercase(value string) bool {
	return regexp.MustCompile(`[A-Z]`).MatchString(value)
}


// Validated if a string contains at least one lower character or not
func HasLowercase(value string) bool {
	return regexp.MustCompile(`[a-z]`).MatchString(value)
}


// Validated if a string contains at least one digit or not
func HasDigit(value string) bool {
	return regexp.MustCompile(`[0-9]`).MatchString(value)
}


// Validated if a string contains at least one punctuation or not
func HasPunctuation(value string) bool {
	for _, item := range value {
		if unicode.IsPunct(item) {
			return true
		}
	}
	return false
}
