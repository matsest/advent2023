package main

import (
	"reflect"
	"testing"

	"github.com/matsest/advent2023/utils"
)

func Test_allZeros(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all zeros",
			args: args{[]int{0, 0, 0}},
			want: true,
		},
		{
			name: "one non-zero",
			args: args{[]int{0, 1, 0}},
			want: false,
		},
		{
			name: "all non-zero",
			args: args{[]int{1, 1, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allZeros(tt.args.slice); got != tt.want {
				t.Errorf("allZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extraPolate(t *testing.T) {
	type args struct {
		history []int
	}
	tests := []struct {
		name             string
		args             args
		wantExtrapolated []int
	}{
		{
			name:             "test1",
			args:             args{[]int{0, 3, 6, 9, 12, 15}},
			wantExtrapolated: []int{0, 3, 6, 9, 12, 15, 18},
		},
		{
			name:             "test2",
			args:             args{[]int{1, 3, 6, 10, 15, 21}},
			wantExtrapolated: []int{1, 3, 6, 10, 15, 21, 28},
		},
		{
			name:             "test3",
			args:             args{[]int{10, 13, 16, 21, 30, 45}},
			wantExtrapolated: []int{10, 13, 16, 21, 30, 45, 68},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExtrapolated := extraPolate(tt.args.history); !reflect.DeepEqual(gotExtrapolated, tt.wantExtrapolated) {
				t.Errorf("extraPolate() = %v, want %v", gotExtrapolated, tt.wantExtrapolated)
			}
		})
	}
}

func Test_p1(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "test",
			args: args{"test.txt"},
			wantSum: 114,
		},
	}
	for _, tt := range tests {
		puzzle_input, _ := utils.ReadLines(tt.args.fileName)
		dataset := parseInput(puzzle_input)
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := p1(dataset); gotSum != tt.wantSum {
				t.Errorf("p1() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
