package equalfrequency

import "testing"

func TestEqualFrequency(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "wrong case 1",
			word: "bac",
			want: true,
		},
		{
			name: "example 1",
			word: "abcc",
			want: true,
		},
		{
			name: "example 2",
			word: "aazz",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualFrequency(tt.word); got != tt.want {
				t.Errorf("EqualFrequency('%s') = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
