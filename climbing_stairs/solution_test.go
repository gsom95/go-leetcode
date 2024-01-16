package climbingstairs

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "my test 1",
			n:    1,
			want: 1,
		},
		{
			name: "my test 2",
			n:    4,
			want: 5,
		},
		{
			name: "example 1",
			n:    2,
			want: 2,
		},
		{
			name: "example 2",
			n:    3,
			want: 3,
		},
		{
			name: "failed test 1",
			n:    44,
			want: 1134903170,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
