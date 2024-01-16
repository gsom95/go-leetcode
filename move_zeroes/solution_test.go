package movezeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "my test 1",
			nums:   []int{1, 0, 3, 12},
			expect: []int{1, 3, 12, 0},
		},
		{
			name:   "my test 2",
			nums:   []int{1, 0, 0, 3, 12},
			expect: []int{1, 3, 12, 0, 0},
		},
		{
			name:   "example 1",
			nums:   []int{0, 1, 0, 3, 12},
			expect: []int{1, 3, 12, 0, 0},
		},
		{
			name:   "example 2",
			nums:   []int{0},
			expect: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expect) {
				t.Errorf("expect: %v, got: %v", tt.expect, tt.nums)
			}
		})
	}
}
