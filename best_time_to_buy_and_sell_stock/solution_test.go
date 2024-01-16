package besttimetobuyandsellstock_test

import (
	"math/rand"
	"testing"
	"time"

	. "github.com/gsom95/go-leetcode/best_time_to_buy_and_sell_stock"
)

func TestMaxProfitI(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "example 2",
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "failed test 1",
			input:    []int{1, 2},
			expected: 1,
		},
		{
			name:     "failed test 2",
			input:    []int{2, 4, 1},
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := MaxProfitI(tC.input)
			if got != tC.expected {
				t.Errorf("got %d, expect %d", got, tC.expected)
			}
		})
	}
}

func TestMaxProfitII(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "example 2",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "example 3",
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "my test 1",
			input:    []int{7, 1, 1, 3, 1},
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := MaxProfitII(tC.input)
			if got != tC.expected {
				t.Errorf("got %d, expect %d", got, tC.expected)
			}
		})
	}
}

func TestMaxProfitIII(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{3, 3, 5, 0, 0, 3, 1, 4},
			expected: 6,
		},
		{
			name:     "example 2",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "example 3",
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "failed test 1",
			input:    []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			expected: 13,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := MaxProfitIII(tC.input)
			if got != tC.expected {
				t.Errorf("got %d, expect %d", got, tC.expected)
			}
		})
	}
}

func TestMaxProfitIV(t *testing.T) {
	type input struct {
		k      int
		prices []int
	}

	var testCases = []struct {
		name string
		input
		expected int
	}{
		{
			name: "example 1",
			input: input{
				k:      2,
				prices: []int{2, 4, 1},
			},
			expected: 2,
		},
		{
			name: "example 2",
			input: input{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			expected: 7,
		},
		{
			name: "failed test 1",
			input: input{
				k:      2,
				prices: []int{},
			},
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			got := MaxProfitIV(tC.k, tC.prices)
			if got != tC.expected {
				t.Errorf("got %d, expect %d", got, tC.expected)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	rand.Seed(time.Now().UnixMilli())
	size := rand.Intn(1024)
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = rand.Intn(10000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := MaxProfitI(input)
		_ = a
	}
}
