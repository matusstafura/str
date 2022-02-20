package str_test

import (
	"testing"

	"github.com/matusstafura/str"
)

var tests = []struct {
	given string
	wanted string
} {
	{"UPPER", "upper"},
	{"TEST", "test"},
}

func TestLength(t *testing.T) {
	quote := str.Length("Acta non verba.")
	length := 15

	if quote != length {
		t.Fatalf("got %q, wanted %q", quote, length)
	}
}

func TestLower(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.given, func(t *testing.T) {
			give := str.Lower(tt.given)
			if give != tt.wanted {
				t.Fatalf("got %q, wanted %q", tt.given, tt.wanted)
			}
		})
	}
}

func TestLimit(t *testing.T) {
	got := str.Limit("Pack my box with five dozen liquor jugs.", 10)
	want := "Pack my bo"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestReverse(t *testing.T) {
	got := str.Reverse("Never give up")
	want := "pu evig reveN"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}