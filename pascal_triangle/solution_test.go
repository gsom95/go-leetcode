package pascaltriangle

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "example 1",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "example 2",
			numRows: 1,
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
