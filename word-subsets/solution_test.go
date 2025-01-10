package wordsubsets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example",
			args: args{
				words1: []string{"warrior", "world"},
				words2: []string{"wrr"},
			},
			want: []string{"warrior"},
		},
		{
			name: "Example 1",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			want: []string{"facebook", "google", "leetcode"},
		},
		{
			name: "Example 2",
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"l", "e"},
			},
			want: []string{"apple", "google", "leetcode"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordSubsets(tt.args.words1, tt.args.words2)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
