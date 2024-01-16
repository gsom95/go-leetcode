package maximumsubarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "example 2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "example 3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "my test 1",
			nums: []int{-2, 3, -1, 2},
			want: 4,
		},
		{
			name: "my test 2",
			nums: []int{-1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
