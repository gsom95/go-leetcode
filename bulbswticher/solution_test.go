package bulbswitcher

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 2},
	{6, 2},
	{7, 2},
	{8, 2},
	{9, 3},
	{12, 3},
	{16, 4},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d", test.input), func(t *testing.T) {
			got := bulbSwitch(test.input)
			if got != test.expected {
				t.Errorf("Got %d, expected %d\n", got, test.expected)
			}
		})
	}
}
