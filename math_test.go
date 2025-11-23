package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

func TestPow(t *testing.T) {
	got := Pow(2, 3)
	want := 8
	if got != want {
		t.Errorf("Pow(2, 3) = %d; want %d", got, want)
	}

	got = Pow(5, 0)
	want = 1
	if got != want {
		t.Errorf("Pow(5, 0) = %d; want %d", got, want)
	}
}
