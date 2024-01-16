package cellsinrange

import (
	"reflect"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "example 1",
			s:    "K1:L2",
			want: []string{"K1", "K2", "L1", "L2"},
		},
		{
			name: "example 2",
			s:    "A1:F1",
			want: []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
