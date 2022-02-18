package str

import (
	"strings"
)

func Limit(s string, n int) string {
	var f string
	for i:=0; i<n; i++ {
		f += string(s[i])
	}
	return f
}

func Lower(s string) string {
	return strings.ToLower(s)
}