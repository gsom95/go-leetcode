package remove_all_occurrences_of_a_substring

import (
	"testing"
)

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s:    "daabcbaabcbc",
				part: "abc",
			},
			want: "dab",
		},
		{
			name: "Example 2",
			args: args{
				s:    "axxxxyyyyb",
				part: "xy",
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOccurrences(tt.args.s, tt.args.part); got != tt.want {
				t.Errorf("removeOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
