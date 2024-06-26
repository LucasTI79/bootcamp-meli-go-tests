package main

import "testing"

func TestFactorial(t *testing.T) {
	// table test
	tests := []struct {
		arg  int
		want int
	}{{0, 1}, {5, 120}, {4, 24}}

	for i, d := range tests {
		got := factorial(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
