package eliminate_maximum_number_of_monsters

import (
	"testing"
)

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-case-1",
			args: args{
				dist:  []int{1, 3, 4},
				speed: []int{1, 1, 1},
			},
			want: 3,
		},
		{
			name: "test-case-2",
			args: args{
				dist:  []int{1, 1, 2, 3},
				speed: []int{1, 1, 1, 1},
			},
			want: 1,
		},
		{
			name: "test-case-3",
			args: args{
				dist:  []int{3, 2, 4},
				speed: []int{5, 3, 2},
			},
			want: 1,
		},
		{
			name: "test-case-4",
			args: args{
				dist:  []int{4, 2, 8},
				speed: []int{2, 1, 4},
			},
			want: 2,
		},
		{
			name: "failed-test-1",
			args: args{
				dist:  []int{4, 2},
				speed: []int{5, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eliminateMaximum(tt.args.dist, tt.args.speed); got != tt.want {
				t.Errorf("eliminateMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
