package main

import (
	"testing"

	"github.com/matsest/advent2023/utils"
)

func Test_hash(t *testing.T) {
	type args struct {
		step string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "HASH",
			args: args{"HASH"},
			want: 52,
		},
		{
			name: "rn=1",
			args: args{"rn=1"},
			want: 30,
		},
		{
			name: "cm-",
			args: args{"cm-"},
			want: 253,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.step); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "test",
			args:    args{"test.txt"},
			wantSum: 1320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle_input, _ := utils.ReadLines("test.txt")
			steps := parseInput(puzzle_input)
			if gotSum := p1(steps); gotSum != tt.wantSum {
				t.Errorf("p1() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "test",
			args:    args{"test.txt"},
			wantSum: 145,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle_input, _ := utils.ReadLines("test.txt")
			steps := parseInput(puzzle_input)
			if gotSum := p2(steps); gotSum != tt.wantSum {
				t.Errorf("p2() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
