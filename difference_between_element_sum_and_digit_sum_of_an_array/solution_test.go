package differencebetweenelementsumanddigitsumofanarray

import "testing"

func Test_differenceOfSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{1, 15, 6, 3}},
			want: 9,
		},
		{
			name: "example 2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSum(tt.args.nums); got != tt.want {
				t.Errorf("differenceOfSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{1},
			want: 1,
		},
		{
			name: "negative number",
			args: args{-1},
			want: 1,
		},
		{
			name: "zero",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.a); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
