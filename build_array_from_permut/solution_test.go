package buildarrayfrompermut

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{0, 2, 1, 5, 3, 4},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "test 2",
			nums: []int{5, 0, 1, 2, 3, 4},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildArray(tt.nums)
			require.Equal(t, tt.want, got)
		})
	}
}
