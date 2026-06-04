package subsets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "Example 2",
			nums: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
