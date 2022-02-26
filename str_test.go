package str

import "testing"

func TestAfter(t *testing.T) {
	got := After("Do not let the behavior of others destroy your inner peace. - Dalai Lama", "- ")
	want := "Dalai Lama"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestAppend(t *testing.T) {
	got := Append("Matus", " Stafura")
	want := "Matus Stafura"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestBefore(t *testing.T) {
	got := Before("Do not let the behavior of others destroy your inner peace. - Dalai Lama", " -")
	want := "Do not let the behavior of others destroy your inner peace."

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestEndsWith(t *testing.T) {
	got := EndsWith("abcdef", "ef")
	want := true

	if got != want {
		t.Fatalf("got %t, wanted %t", got, want)
	}
}

func TestLength(t *testing.T) {
	quote := Length("Acta non verba.")
	length := 15

	if quote != length {
		t.Fatalf("got %q, wanted %q", quote, length)
	}
}

func TestLimit(t *testing.T) {
	got := Limit("Pack my box with five dozen liquor jugs.", 10)
	want := "Pack my bo"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestLower(t *testing.T) {
	got := Lower("JANUARY")
	want := "january"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestMask(t *testing.T) {
	got := Mask("4242 4242 4242 4242 4242", 4, "#")
	want := "4242####################"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestReverse(t *testing.T) {
	got := Reverse("Never give up")
	want := "pu evig reveN"

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}

func TestStartsWith(t *testing.T) {
	got := StartsWith("abcdef", "ab")
	want := true

	if got != want {
		t.Fatalf("got %t, wanted %t", got, want)
	}
}

func TestUpper(t *testing.T) {
	got := Upper("Be the friend you wish you had.")
	want := "BE THE FRIEND YOU WISH YOU HAD."

	if got != want {
		t.Fatalf("got %q, wanted %q", got, want)
	}
}
