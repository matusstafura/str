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