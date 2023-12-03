package main

import (
	"testing"
)

func Test_possibleGame(t *testing.T) {
	tests := []struct {
		name string
		args Game
		want bool
	}{
		{"possible1", Game{1, []Set{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}}, true},
		{"possible2", Game{2, []Set{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}}, true},
		{"impossible", Game{3, []Set{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}}}, false},
		{"impossible2", Game{4, []Set{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}}}, false},
		{"possible3", Game{5, []Set{{6, 3, 1}, {1, 2, 2}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := possibleGame(tt.args); gotCount != tt.want {
				t.Errorf("possibleGame(%v) = %v, want %v", tt.args, gotCount, tt.want)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"default", "args", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := p2(tt.args); gotCount != tt.want {
				t.Errorf("p2(%v) = %v, want %v", tt.args, gotCount, tt.want)
			}
		})
	}
}