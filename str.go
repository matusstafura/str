package str

import (
	"strings"
)

func Append(s ...string) string {
	return strings.Join(s, "")
}

func Before(s, c string) string {
	ma := strings.Split(s, c)
	return string(ma[0])
}

func Length(s string) int {
	return len(s)
}

func Limit(s string, n int) string {
	return s[0:n]
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Reverse(s string) string {
	f := make([]string, len(s))

	for i, c := range s {
		f[len(s)-1-i] = string(c)
	}
	return strings.Join(f, "")
}
