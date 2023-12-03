package main

import (
	"testing"
)

func Test_isSpecialCharacter(t *testing.T) {
	type args struct {
		input rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Basic test", args{'*'}, true},
		{"Period character", args{'.'}, false},
		{"Digit 1", args{'1'}, false},
		{"Digit 5", args{'5'}, false},
		{"Digit 9", args{'9'}, false},
		{"Special char 1", args{'#'}, true},
		{"Special char 2", args{'@'}, true},
		{"Special char 3", args{'^'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSpecialCharacter(tt.args.input); got != tt.want {
				t.Errorf("isSpecialCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateLocalSum(t *testing.T) {
	type args struct {
		row    int
		column int
		matrix [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Basic test",
			args{
				row:    1,
				column: 1,
				matrix: [][]rune{
					{'.', '1', '.'},
					{'.', '*', '.'},
					{'1', '0', '.'},
				},
			},
			11,
		},
		{
			"Special character in the first row",
			args{
				row:    0,
				column: 0,
				matrix: [][]rune{
					{'*', '.', '.'},
					{'3', '0', '.'},
					{'1', '0', '.'},
				},
			},
			30,
		},
		{
			"Special character in the last row",
			args{
				row:    2,
				column: 2,
				matrix: [][]rune{
					{'.', '.', '.'},
					{'3', '0', '.'},
					{'1', '0', '*'},
				},
			},
			40,
		},
		{
			"Special character at (1, 0)",
			args{
				row:    1,
				column: 0,
				matrix: [][]rune{
					{'.', '.', '.'},
					{'*', '3', '5'},
					{'1', '9', '0'},
				},
			},
			225,
		},
		{
			"Special character in the last column",
			args{
				row:    1,
				column: 0,
				matrix: [][]rune{
					{'.', '.', '.'},
					{'3', '5', '*'},
					{'1', '9', '0'},
				},
			},
			225,
		},
		{
			"Special character with no number nearby",
			args{
				row:    2,
				column: 0,
				matrix: [][]rune{
					{'.', '1', '.'},
					{'.', '.', '4'},
					{'*', '.', '8'},
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLocalSum(tt.args.row, tt.args.column, tt.args.matrix); got != tt.want {
				t.Errorf("calculateLocalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
