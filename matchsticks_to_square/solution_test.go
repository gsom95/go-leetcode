package matchstickstosquare

import "testing"

func Test_makesquare(t *testing.T) {
	tests := []struct {
		name        string
		matchsticks []int
		want        bool
	}{
		// {
		// 	name:        "example 1",
		// 	matchsticks: []int{1, 1, 2, 2, 2},
		// 	want:        true,
		// },
		// {
		// 	name:        "example 2",
		// 	matchsticks: []int{3, 3, 3, 3, 4},
		// 	want:        false,
		// },
		// {
		// 	name:        "len < 4",
		// 	matchsticks: nil,
		// 	want:        false,
		// },
		// {
		// 	name:        "false 1",
		// 	matchsticks: []int{8, 4, 4, 4},
		// 	want:        false,
		// },
		{
			name:        "fail 1",
			matchsticks: []int{12, 18, 2, 2, 16, 8, 7, 3, 10, 12, 3, 20, 2, 10, 19},
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
