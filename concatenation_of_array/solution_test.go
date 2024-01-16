package concatenationofarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "test 2",
			nums: []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getConcatenation(tt.nums)

			require.Equal(t, tt.want, got)
		})
	}
}
