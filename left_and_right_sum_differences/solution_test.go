package leftandrightsumdifferences

import (
	"reflect"
	"testing"
)

func Test_leftRigthDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{10, 4, 8, 3},
			},
			want: []int{15, 1, 11, 22},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{1},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRigthDifference(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftRigthDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil slice",
			args: args{
				nums: nil,
			},
			want: 0,
		},
		{
			name: "empty slice",
			args: args{
				nums: nil,
			},
			want: 0,
		},
		{
			name: "one element",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "many elements",
			args: args{
				nums: []int{42, 54, 2},
			},
			want: 98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.nums); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
