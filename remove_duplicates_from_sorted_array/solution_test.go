package removeduplicatesfromsortedarray

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
		want     int
	}{
		{
			name:     "my test 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
			want:     3,
		},
		{
			name:     "example 1",
			nums:     []int{1, 1, 2},
			expected: []int{1, 2},
			want:     2,
		},
		{
			name:     "example 2",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
			want:     5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
				return
			}

			for i, v := range tt.expected {
				if v != tt.nums[i] {
					t.Errorf("removeDuplicates(): got[%d]=%d, expected[%d]=%d", i, v, i, tt.expected[i])
					return
				}
			}
		})
	}
}
