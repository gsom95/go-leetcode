package happy_number

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    19,
			want: true,
		},
		{
			name: "Example 2",
			n:    2,
			want: false,
		},
		{
			name: "Example 3",
			n:    7,
			want: true,
		},
		{
			name: "Example 4",
			n:    1111111,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
