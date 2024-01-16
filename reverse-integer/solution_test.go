package reverseinteger

import (
	"fmt"
	"math"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    math.MaxInt32,
			want: 0,
		},
		{
			x:    math.MinInt32,
			want: 0,
		},
		{
			x:    math.MaxInt32 + 1,
			want: 0,
		},
		{
			x:    math.MinInt32 - 1,
			want: 0,
		},
		{
			x:    9,
			want: 9,
		},
		{
			x:    19,
			want: 91,
		},
		{
			x:    10,
			want: 1,
		},
		{
			x:    123,
			want: 321,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    120,
			want: 21,
		},
		{
			x:    0,
			want: 0,
		},
		{
			x:    1,
			want: 1,
		},
		{
			x:    -1,
			want: -1,
		},
		{
			x:    94,
			want: 49,
		},
		{
			x:    -45,
			want: -54,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.x), func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
