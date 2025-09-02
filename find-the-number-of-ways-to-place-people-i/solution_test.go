package findthenumberofwaystoplacepeoplei

import (
	"encoding/json"
	"testing"
)

func Test_numberOfPairs(t *testing.T) {
	tests := []struct {
		name   string
		points []byte
		want   int
	}{
		{
			name:   "my test 1",
			points: []byte(`[[1,1], [2,1]]`),
			want:   1,
		},
		{
			name:   "my test 2",
			points: []byte(`[[1,1], [1,2]]`),
			want:   1,
		},
		{
			name:   "my test 3",
			points: []byte(`[[2,1], [1,1]]`),
			want:   1,
		},
		{
			name:   "my test 4",
			points: []byte(`[[1,2], [1,1]]`),
			want:   1,
		},
		{
			name:   "example 1",
			points: []byte(`[[1,1],[2,2],[3,3]]`),
			want:   0,
		},
		{
			name:   "example 2",
			points: []byte(`[[6,2],[4,4],[2,6]]`),
			want:   2,
		},
		{
			name:   "example 3",
			points: []byte(`[[3,1],[1,3],[1,1]]`),
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var points [][]int

			if err := json.Unmarshal(tt.points, &points); err != nil {
				t.Errorf("Unmarshal() error = %v", err)
				return
			}

			if got := numberOfPairs(points); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
