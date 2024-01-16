package subrectanglequeries

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want SubrectangleQueries
	}{
		{
			"test constructor",
			[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}},
			SubrectangleQueries{[][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
