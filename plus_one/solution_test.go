package plus_one

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{
			name: "example 1",
			digits: []int{
				1, 2, 3,
			},
			want: []int{
				1, 2, 4,
			},
		},
		{
			name: "example 2",
			digits: []int{
				4, 3, 2, 1,
			},
			want: []int{
				4, 3, 2, 2,
			},
		},
		{
			name: "example 3",
			digits: []int{
				9,
			},
			want: []int{
				1, 0,
			},
		},
		{
			name: "my test 1",
			digits: []int{
				9, 9, 9,
			},
			want: []int{
				1, 0, 0, 0,
			},
		},
		{
			name: "my test 2",
			digits: []int{
				1, 9, 9,
			},
			want: []int{
				2, 0, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
