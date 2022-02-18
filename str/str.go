package str

import (
	"strings"
)

func Limit(s string, n int) string {
	return s[0:n]
}

func Lower(s string) string {
	return strings.ToLower(s)
}