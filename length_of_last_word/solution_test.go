package lengthoflastword

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "Hello World",
			want: 5,
		},
		{
			name: "example 2",
			s:    "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "example 3",
			s:    "luffy is still joyboy",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
