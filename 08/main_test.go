package main

import (
	"testing"

	"github.com/matsest/advent2023/utils"
)

func Test_p1(t *testing.T) {
	tests := []struct {
		name      string
		fileName string
		want int
	}{
		{"test1", "test1.txt", 2},
		{"test2", "test2.txt", 6},
		{"input", "input.txt", 16409},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle_input, _ := utils.ReadLines(tt.fileName)
			instructions, nodes := parseInput(puzzle_input)
			if gotCount := p1(instructions, nodes); gotCount != tt.want {
				t.Errorf("p1(%v) = %v, want %v", tt.fileName, gotCount, tt.want)
			}
		})
	}
}