package str

import "strings"

// Repeat returns a new string consisting of count copies of the string str
func Repeat(str string, count int) string {
	if IsEmpty(str) || count == 0 {
		return str
	}
	return strings.Repeat(str, count)
}

// IsEmpty returns a boolean of s equal "" or len(s) == 0
func IsEmpty(str string) bool {
	return "" == str || 0 == len(str)
}

// IsNotEmpty returns a boolean of s != "" and len(s) != 0
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
