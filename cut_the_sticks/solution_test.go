package cutthesticks

import (
	"fmt"
	"testing"

	"github.com/gsom95/go-leetcode/utils"
)

type testpair struct {
	inputArray     []int
	expectedResult []int
}

var tests = []testpair{
	{inputArray: []int{1, 2, 2}, expectedResult: []int{3, 2}},
	{inputArray: []int{1, 2, 3}, expectedResult: []int{3, 2, 1}},
	{inputArray: []int{5, 4, 4, 2, 2, 8}, expectedResult: []int{6, 4, 2, 1}},
	{inputArray: []int{1, 2, 3, 4, 3, 3, 2, 1}, expectedResult: []int{8, 6, 4, 1}},
}

func TestCutTheSticks(t *testing.T) {
	for i, testCase := range tests {
		testInput := testCase.inputArray
		testName := fmt.Sprintf("Test #%d: %v", i, testInput)
		t.Run(testName, func(t *testing.T) {
			result := cutTheSticks(testInput)
			if !utils.EqualIntSlices(result, testCase.expectedResult) {
				t.Error(
					"\nexpected:", testCase.expectedResult, "\n",
					"got:", result,
				)
			}
		})

	}
}
