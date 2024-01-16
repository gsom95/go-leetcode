package reshapethematrix

import (
	"reflect"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	type args struct {
		mat [][]int
		r   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "failed test 1",
			args: args{
				mat: [][]int{
					{1, 2},
				},
				r: 1,
				c: 1,
			},
			want: [][]int{
				{1, 2},
			},
		},
		{
			name: "my test 1",
			args: args{
				mat: [][]int{
					{1},
				},
				r: 1,
				c: 1,
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "my test 2",
			args: args{
				mat: [][]int{
					{1},
					{2},
				},
				r: 1,
				c: 2,
			},
			want: [][]int{
				{1, 2},
			},
		},
		{
			name: "example 1",
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
				r: 1,
				c: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name: "example 2",
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
				r: 2,
				c: 4,
			},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.mat, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
