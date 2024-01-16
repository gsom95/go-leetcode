package isomorphic_strings

import (
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "egg",
			t:    "add",
			want: true,
		},
		{
			name: "Example 2",
			s:    "foo",
			t:    "bar",
			want: false,
		},
		{
			name: "Example 3",
			s:    "paper",
			t:    "title",
			want: true,
		},
		{
			name: "Example 4",
			s:    "badc",
			t:    "baba",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
