package rotatearray

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		expect []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			expect: []int{3, 99, -1, -100},
		},
		{
			name: "my test 1",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    4,
			},
			expect: []int{-1, -100, 3, 99},
		},
		{
			name: "failed test 1",
			args: args{
				nums: []int{1},
				k:    1,
			},
			expect: []int{1},
		},
		{
			name: "failed test 2",
			args: args{
				nums: []int{1, 2},
				k:    1,
			},
			expect: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.expect) {
				t.Errorf("expect: %v, got: %v", tt.expect, tt.args.nums)
			}
		})
	}
}
