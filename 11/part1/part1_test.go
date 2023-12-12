package main

import "testing"

func newCoord(row int, col int) *coordinate {
	return &coordinate{row: row, col: col}
}

func Test_findDistance(t *testing.T) {
	type args struct {
		first  *coordinate
		second *coordinate
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{newCoord(6, 1), newCoord(11, 5)}, 9},
		{"Example 2", args{newCoord(0, 4), newCoord(10, 9)}, 15},
		{"Example 3", args{newCoord(2, 0), newCoord(7, 12)}, 17},
		{"Example 4", args{newCoord(11, 0), newCoord(11, 5)}, 5},
		{"Example 5", args{newCoord(0, 4), newCoord(1, 9)}, 6},
		{"Example 6", args{newCoord(0, 4), newCoord(2, 0)}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDistance(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("findDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
