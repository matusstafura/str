package str_test

import (
	"testing"

	"github.com/matusstafura/str"
)

func TestAppend(t *testing.T) {
	got := str.Append("Matus"," Stafura")
	want := "Matus Stafura"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestLength(t *testing.T) {
	quote := str.Length("Acta non verba.")
	length := 15

	if quote != length {
		t.Fatalf("got %q, wanted %q", quote, length)
	}
}

func TestLower(t *testing.T) {
	got := str.Lower("JANUARY")
	want := "january"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
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