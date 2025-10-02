package removelettertoequalizefrequency

import "testing"

func Test_equalFrequency(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{word: "abcc"},
			want: true,
		},
		{
			name: "example 2",
			args: args{word: "aazz"},
			want: false,
		},
		{
			name: "failed test 1",
			args: args{word: "zz"},
			want: true,
		},
		{
			name: "failed test 2",
			args: args{word: "bac"},
			want: true,
		},
		{
			name: "failed test 3",
			args: args{word: "ddaccb"},
			want: false,
		},
		{
			name: "failed test 4",
			args: args{word: "abbcc"},
			want: true,
		},
		{
			name: "failed test 5",
			args: args{word: "cccd"},
			want: true,
		},
		{
			name: "failed test 6",
			args: args{word: "cbccca"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalFrequency(tt.args.word); got != tt.want {
				t.Errorf("equalFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
