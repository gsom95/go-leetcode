package sqrtx

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "Example 1",
			x:    4,
			want: 2,
		},
		{
			name: "Example 2",
			x:    8,
			want: 2,
		},
		{
			name: "Example 3",
			x:    0,
			want: 0,
		},
		{
			name: "Example 4",
			x:    1,
			want: 1,
		},
		{
			name: "Example 5",
			x:    2,
			want: 1,
		},
		{
			name: "Example 6",
			x:    3,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
