package setmatrixzeroes

import (
	"slices"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name   string
		args   args
		expect [][]int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			expect: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			expect: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			for i, row := range tt.args.matrix {
				if !slices.Equal(row, tt.expect[i]) {
					t.Fatalf("setZeroes() = %v, want %v — row %d", row, tt.expect[i], i)
				}
			}
		})
	}
}
