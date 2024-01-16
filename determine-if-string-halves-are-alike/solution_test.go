package determine_if_string_halves_are_alike

import (
	"testing"
)

func Test_halvesAreAlike(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "book",
			want: true,
		},
		{
			name: "example 2",
			s:    "textbook",
			want: false,
		},
		{
			name: "example 3",
			s:    "MerryChristmas",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlike(tt.s); got != tt.want {
				t.Errorf("halvesAreAlike() = %v, want %v", got, tt.want)
			}
		})
	}
}
