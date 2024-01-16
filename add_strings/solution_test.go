package add_strings

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{"11", "123"},
			want: "134",
		},
		{
			name: "example 2",
			args: args{"456", "77"},
			want: "533",
		},
		{
			name: "example 3",
			args: args{"0", "0"},
			want: "0",
		},
		{
			name: "example 4",
			args: args{"1", "9"},
			want: "10",
		},
		{
			name: "example 5",
			args: args{"9", "99"},
			want: "108",
		},
		{
			name: "example 6",
			args: args{"408", "5"},
			want: "413",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
