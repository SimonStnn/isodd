package isodd_test

import (
	. "github.com/SimonStnn/isodd"
	"testing"
)

func TestIsOdd(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{-5, true},
		{-4, false},
		{-3, true},
		{-2, false},
		{-1, true},
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}

	for _, tt := range tests {
		if got := IsOdd(tt.input); got != tt.want {
			t.Errorf("IsOdd(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
