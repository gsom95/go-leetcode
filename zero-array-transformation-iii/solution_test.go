package zeroarraytransformationiii

import "testing"

func Test_maxRemoval(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{2, 0, 2},
				queries: [][]int{
					{0, 2},
					{0, 2},
					{1, 1},
				},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 1, 1, 1},
				queries: [][]int{
					{1, 3},
					{0, 2},
					{1, 3},
					{1, 2},
				},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2, 3, 4},
				queries: [][]int{
					{0, 3},
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRemoval(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("maxRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
