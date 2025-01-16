package finalpriceswithaspecialdiscountinashop

import (
	"reflect"
	"testing"
)

func Test_finalPrices(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "My test 1",
			args: args{
				prices: []int{2, 2},
			},
			want: []int{0, 2},
		},
		{
			name: "My test 2",
			args: args{
				prices: []int{2, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "My test 3",
			args: args{
				prices: []int{8, 7, 8},
			},
			want: []int{1, 7, 8},
		},
		{
			name: "Example 1",
			args: args{
				prices: []int{8, 4, 6, 2, 3},
			},
			want: []int{4, 2, 4, 2, 3},
		},
		{
			name: "Example 2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Example 3",
			args: args{
				prices: []int{10, 1, 1, 6},
			},
			want: []int{9, 0, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := finalPrices(tt.args.prices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("finalPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
