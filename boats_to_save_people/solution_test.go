package boatstosavepeople

import "testing"

func Test_numRescueBoats(t *testing.T) {
	tests := []struct {
		name   string
		people []int
		limit  int
		want   int
	}{
		{
			name:   "example 1",
			people: []int{1, 2},
			limit:  3,
			want:   1,
		},
		{
			name:   "example 2",
			people: []int{3, 2, 2, 1},
			limit:  3,
			want:   3,
		},
		{
			name:   "example 3",
			people: []int{3, 5, 3, 4},
			limit:  5,
			want:   4,
		},
		{
			name:   "wrong answer 1",
			people: []int{5, 1, 4, 2},
			limit:  6,
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.people, tt.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
