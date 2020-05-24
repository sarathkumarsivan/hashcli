package util

import "strings"

// IsEmpty returns true if the string is empty and false otherwise.
// Empty spaces and all leading and trailing white spaces are removed
// before checking the length.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty returns true if the string is NOT empty and false otherwise.
// Empty spaces and all leading and trailing white spaces are removed
// before checking the length.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
