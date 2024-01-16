package addbinary

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "failed test 1",
			args: args{
				a: "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
				b: "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			},
			want: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
		{
			name: "example 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "example 2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "my test 1",
			args: args{
				a: "0",
				b: "1",
			},
			want: "1",
		},
		{
			name: "my test 2",
			args: args{
				a: "10",
				b: "1",
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
			if got := addBinaryPlain(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinaryPlain() = %v, want %v", got, tt.want)
			}
		})
	}
}
