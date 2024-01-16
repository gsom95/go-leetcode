package pascaltriangle2

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{
			name:     "example 1",
			rowIndex: 3,
			want:     []int{1, 3, 3, 1},
		},
		{
			name:     "example 2",
			rowIndex: 0,
			want:     []int{1},
		},
		{
			name:     "example 3",
			rowIndex: 1,
			want:     []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
