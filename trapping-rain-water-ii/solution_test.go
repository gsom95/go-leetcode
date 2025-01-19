package trappingrainwaterii

import (
	"testing"
)

func Test_trapRainWater(t *testing.T) {
	type args struct {
		heightMap [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				heightMap: [][]int{
					{1, 4, 3, 1, 3, 2},
					{3, 2, 1, 3, 2, 4},
					{2, 3, 3, 2, 3, 1},
				},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				heightMap: [][]int{
					{3, 3, 3, 3, 3},
					{3, 2, 2, 2, 3},
					{3, 2, 1, 2, 3},
					{3, 2, 2, 2, 3},
					{3, 3, 3, 3, 3},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapRainWater(tt.args.heightMap); got != tt.want {
				t.Errorf("trapRainWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
