package make_two_arrays_equal_by_reversing_subarrays

import (
	"testing"
)

type args struct {
	target []int
	arr    []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "Example 1",
		args: args{
			target: []int{1, 2, 3, 4},
			arr:    []int{2, 4, 1, 3},
		},
		want: true,
	},
	{
		name: "Example 2",
		args: args{
			target: []int{7},
			arr:    []int{7},
		},
		want: true,
	},
	{
		name: "Example 3",
		args: args{
			target: []int{3, 7, 9},
			arr:    []int{3, 7, 11},
		},
		want: false,
	},
	{
		name: "Example 4",
		args: args{
			target: []int{1, 12},
			arr:    []int{12, 1},
		},
		want: true,
	},
	{
		name: "Example 5",
		args: args{
			target: []int{1, 1, 1, 1, 1},
			arr:    []int{1, 1, 1, 1, 1},
		},
		want: true,
	},
	{
		name: "Example 6",
		args: args{
			target: []int{1, 1, 2, 2, 3, 3},
			arr:    []int{2, 2, 1, 5, 3, 3},
		},
		want: false,
	},
}

func Test_canBeEqual(t *testing.T) {
	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canBeEqualSorted(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqualSorted(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqualSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
