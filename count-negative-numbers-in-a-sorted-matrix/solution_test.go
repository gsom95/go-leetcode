package count_negative_numbers_in_a_sorted_matrix

import (
	"testing"
)

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{4, 3, 2, -1},
					{3, 2, 1, -1},
					{1, 1, -1, -2},
					{-1, -1, -2, -3},
				},
			},
			want: 8,
		},
		{
			name: "Example 2",
			args: args{
				grid: [][]int{
					{3, 2},
					{1, 0},
				},
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				grid: [][]int{
					{1, -1},
					{-1, -1},
				},
			},
			want: 3,
		},
		{
			name: "Example 4",
			args: args{
				grid: [][]int{
					{5, 1, 0},
					{-5, -5, -5},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}
