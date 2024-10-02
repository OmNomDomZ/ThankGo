package IntMinTest

import "testing"

func TestIntMin(t *testing.T) {
	got := IntMin(2, -2)
	want := -2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestIntMin2(t *testing.T) {
	got := IntMin(-2, 2)
	want := -2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
