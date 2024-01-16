package minsumoffourdigit

import "testing"

func Test_minimumSum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "example 1",
			num: 2932,
			want: 52,
		},
		{
			name: "example 2",
			num: 4009,
			want: 13,
		},
		{
			name: "corner case 1",
			num: 1000,
			want: 1,
		},
		{
			name: "corner case 2",
			num: 9999,
			want: 198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.num); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
