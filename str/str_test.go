package str_test

import (
	"strhelpers/str"
	"testing"
)

func TestLower(t *testing.T) {
	got := str.Lower("UPPER")
	want := "upper"

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