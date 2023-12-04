package main

import (
	"reflect"
	"testing"

	"github.com/matsest/advent2023/utils"
)

func TestCard_NMatchingNumbers(t *testing.T) {
	type fields struct {
		index          int
		winningNumbers []int
		numbers        []int
	}
	tests := []struct {
		name    string
		fields  fields
		wantSum int
	}{
		{
			name: "four",
			fields: fields{
				index:          1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				numbers:        []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			wantSum: 4,
		},
		{
			name: "two",
			fields: fields{
				index:          2,
				winningNumbers: []int{13, 32, 20, 16, 61},
				numbers:        []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			wantSum: 2,
		},
		{
			name: "none",
			fields: fields{
				index:          5,
				winningNumbers: []int{87, 83, 26, 28, 32},
				numbers:        []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			wantSum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Card{
				index:          tt.fields.index,
				winningNumbers: tt.fields.winningNumbers,
				numbers:        tt.fields.numbers,
			}
			if gotSum := g.NMatchingNumbers(); gotSum != tt.wantSum {
				t.Errorf("Card.NMatchingNumbers() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestCard_Points(t *testing.T) {
	type fields struct {
		index          int
		winningNumbers []int
		numbers        []int
	}
	tests := []struct {
		name    string
		fields  fields
		wantSum int
	}{
		{
			name: "four matching",
			fields: fields{
				index:          1,
				winningNumbers: []int{41, 48, 83, 86, 17},
				numbers:        []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			wantSum: 8,
		},
		{
			name: "two matching",
			fields: fields{
				index:          2,
				winningNumbers: []int{13, 32, 20, 16, 61},
				numbers:        []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			wantSum: 2,
		},
		{
			name: "none",
			fields: fields{
				index:          5,
				winningNumbers: []int{87, 83, 26, 28, 32},
				numbers:        []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			wantSum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Card{
				index:          tt.fields.index,
				winningNumbers: tt.fields.winningNumbers,
				numbers:        tt.fields.numbers,
			}
			if gotSum := g.Points(); gotSum != tt.wantSum {
				t.Errorf("Card.Points() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name      string
		args      args
		wantCards []Card
	}{
		{
			name: "simple",
			args: args{
				lines: []string{
					"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
					"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
					"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
				},
			},
			wantCards: []Card{
				{

					index:          1,
					winningNumbers: []int{41, 48, 83, 86, 17},
					numbers:        []int{83, 86, 6, 31, 17, 9, 48, 53},
				},
				{

					index:          2,
					winningNumbers: []int{13, 32, 20, 16, 61},
					numbers:        []int{61, 30, 68, 82, 17, 32, 24, 19},
				},
				{
					index:          5,
					winningNumbers: []int{87, 83, 26, 28, 32},
					numbers:        []int{88, 30, 70, 12, 93, 22, 82, 36},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCards := parseInput(tt.args.lines); !reflect.DeepEqual(gotCards, tt.wantCards) {
				t.Errorf("parseInput() = %v, want %v", gotCards, tt.wantCards)
			}
		})
	}
}

func Test_p1(t *testing.T) {
	test_input, _ := utils.ReadLines("test.txt")
	cards := parseInput(test_input)
	type args struct {
		cards []Card
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "testinput",
			args: args{
				cards: cards,
			},
			wantSum: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := p1(tt.args.cards); gotSum != tt.wantSum {
				t.Errorf("p1() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	test_input, _ := utils.ReadLines("test.txt")
	cards := parseInput(test_input)
	type args struct {
		cards []Card
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name: "testinput",
			args: args{
				cards: cards,
			},
			wantSum: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := p2(tt.args.cards); gotSum != tt.wantSum {
				t.Errorf("p2() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}