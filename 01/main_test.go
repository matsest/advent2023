package main

import "testing"

func Test_p1(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want int
	}{
		{"default", []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := p1(tt.args); gotCount != tt.want {
				t.Errorf("p1(%v) = %v, want %v", tt.args, gotCount, tt.want)
			}
		})
	}
}

func Test_p2(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		want int
	}{
		{"default", []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}, 281},
		{"multiple", []string{
			"7tsdxhfrfdkeighteightnine3eight",
		}, 78},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := p2(tt.args); gotCount != tt.want {
				t.Errorf("p2(%v) = %v, want %v", tt.args, gotCount, tt.want)
			}
		})
	}
}