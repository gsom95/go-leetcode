package reversestring

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name   string
		s      []byte
		expect []byte
	}{
		{
			name:   "example 1",
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			expect: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:   "example 2",
			s:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expect: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name:   "my test 1",
			s:      []byte{'H'},
			expect: []byte{'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.expect) {
				t.Errorf("expect: %v, got: %v", tt.expect, tt.s)
			}
		})
	}
}
