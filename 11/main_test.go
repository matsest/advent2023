package main

import (
	"testing"

	"github.com/matsest/advent2023/utils"
)

func Test_countBelow(t *testing.T) {
	type args struct {
		indexes []int
		index   int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name:      "orderedEnd",
			args:      args{[]int{1, 2, 3, 4, 5}, 5},
			wantCount: 4,
		},
		{
			name:      "unordered",
			args:      args{[]int{5, 9, 1, 4, 5}, 4},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countBelow(tt.args.indexes, tt.args.index); gotCount != tt.wantCount {
				t.Errorf("countBelow() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive",
			args: args{5},
			want: 5,
		},
		{
			name: "negative",
			args: args{-5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manhattanDistance(t *testing.T) {
	type args struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no distance",
			args: args{0, 0, 0, 0},
			want: 0,
		},
		{
			name: "one",
			args: args{0, 0, 1, 0},
			want: 1,
		},
		{
			name: "diag",
			args: args{0, 0, 2, 2},
			want: 4,
		},
		{
			name: "three",
			args: args{1, 2, 3, 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattanDistance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p1(t *testing.T) {
	type args struct {
		fileName string
		expansionCoeff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "double",
			args: args{"test.txt",2},
			want: 374,
		},
	}
	for _, tt := range tests {
		puzzle_input, _ := utils.ReadLines(tt.args.fileName)
		_, galaxies := parseInput(puzzle_input, tt.args.expansionCoeff)
		t.Run(tt.name, func(t *testing.T) {
			if got := p1(galaxies); got != tt.want {
				t.Errorf("p1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	type args struct {
		fileName string
		expansionCoeff int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ten",
			args: args{"test.txt",10},
			want: 1030,
		},
		{
			name: "hundred",
			args: args{"test.txt",100},
			want: 8410,
		},
	}
	for _, tt := range tests {
		puzzle_input, _ := utils.ReadLines(tt.args.fileName)
		_, galaxies := parseInput(puzzle_input, tt.args.expansionCoeff)
		t.Run(tt.name, func(t *testing.T) {
			if got := p2(galaxies); got != tt.want {
				t.Errorf("p2() = %v, want %v", got, tt.want)
			}
		})
	}
}