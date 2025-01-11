package constructkpalindromestrings

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "my test 1",
			args: args{
				s: "aabbc",
				k: 4,
			},
			want: true,
		},
		{
			name: "Example 1",
			args: args{
				s: "annabelle",
				k: 2,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				s: "leetcode",
				k: 3,
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				s: "true",
				k: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
