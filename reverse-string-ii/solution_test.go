package reversestringii

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "example 2",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
		{
			name: "example 3",
			args: args{
				s: "abcde",
				k: 2,
			},
			want: "bacde",
		},
		{
			name: "exapmle 4",
			args: args{
				s: "abcd",
				k: 4,
			},
			want: "dcba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
