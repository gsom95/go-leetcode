package fruitintobaskets

import "testing"

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{
			name:   "My Test 1",
			fruits: []int{1},
			want:   1,
		},
		{
			name:   "My Test 2",
			fruits: []int{1, 1, 1},
			want:   3,
		},
		{
			name:   "My Test 3",
			fruits: []int{1, 2},
			want:   2,
		},
		{
			name:   "Example 1",
			fruits: []int{1, 2, 1},
			want:   3,
		},
		{
			name:   "Example 2",
			fruits: []int{0, 1, 2, 2},
			want:   3,
		},
		{
			name:   "Example 3",
			fruits: []int{1, 2, 3, 2, 2},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
