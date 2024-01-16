package multiply_strings

import (
	"testing"
)

func Test_multiply(t *testing.T) {
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
			name: "Example 1",
			args: args{"2", "3"},
			want: "6",
		},
		{
			name: "Example 2",
			args: args{"123", "456"},
			want: "56088",
		},
		{
			name: "Example 3",
			args: args{"0", "0"},
			want: "0",
		},
		{
			name: "Example 4",
			args: args{"123", "10"},
			want: "1230",
		},
		{
			name: "Example 5",
			args: args{"10", "123"},
			want: "1230",
		},
		{
			name: "Failed test 1",
			args: args{"9133", "0"},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
