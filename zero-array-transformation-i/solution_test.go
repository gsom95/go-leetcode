package zeroarraytransformationi

import "testing"

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 0, 1},
				queries: [][]int{
					{0, 2},
				},
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{4, 3, 2, 1},
				queries: [][]int{
					{1, 3},
					{0, 2},
				},
			},
			want: false,
		},
		{
			name: "failed test 1",
			args: args{
				nums: []int{3},
				queries: [][]int{
					{0, 0},
					{0, 0},
				},
			},
			want: false,
		},
		{
			name: "failed test 2",
			args: args{
				nums: []int{1, 2, 1, 0, 0, 0},
				queries: [][]int{
					{0, 3},
					{2, 4},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isZeroArray(tt.args.nums, tt.args.queries); got != tt.want {
				t.Errorf("isZeroArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
