package rectangle_overlap

import (
	"testing"
)

func Test_isRectangleOverlap(t *testing.T) {
	type args struct {
		rec1 []int
		rec2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				rec1: []int{0, 0, 2, 2},
				rec2: []int{1, 1, 3, 3},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				rec1: []int{0, 0, 1, 1},
				rec2: []int{1, 0, 2, 1},
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				rec1: []int{0, 0, 1, 1},
				rec2: []int{2, 2, 3, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRectangleOverlap(tt.args.rec1, tt.args.rec2); got != tt.want {
				t.Errorf("isRectangleOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
