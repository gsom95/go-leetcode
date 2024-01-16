package removingstarsfromastring

import "testing"

func Test_removeStars(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "example 1",
			s:    "leet**cod*e",
			want: "lecoe",
		},
		{
			name: "example 2",
			s:    "erase*****",
			want: "",
		},
		{
			name: "my test 1",
			s:    "noerase",
			want: "noerase",
		},
		{
			name: "my test 2",
			s:    "noerase*",
			want: "noeras",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStars(tt.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
