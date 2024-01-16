package kidswithcandies

import (
	"fmt"
	"testing"

	"github.com/gsom95/go-leetcode/utils"
)

type input struct {
	candies      []int
	extraCandies int
}

type testPair struct {
	i   input
	exp []bool
}

var tests = []testPair{
	{
		i:   input{[]int{2, 3, 5, 1, 3}, 3},
		exp: []bool{true, true, true, false, true},
	},
	{
		i:   input{[]int{4, 2, 1, 1, 2}, 1},
		exp: []bool{true, false, false, false, false},
	},
	{
		i:   input{[]int{12, 1, 12}, 10},
		exp: []bool{true, false, true},
	},
}

func TestKidsWithCandies(t *testing.T) {
	for i, testCase := range tests {
		testInput := testCase.i
		testName := fmt.Sprintf("Test #%d: %v\n", i, testCase.i)
		t.Run(testName, func(t *testing.T) {
			result := kidsWithCandies(testInput.candies, testInput.extraCandies)
			if !utils.EqualBoolSlices(result, testCase.exp) {
				t.Error(
					"\nExpected:", testCase.exp, "\n",
					"got:", result,
				)
			}
		})
	}
}
