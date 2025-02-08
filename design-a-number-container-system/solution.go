package main

import (
	"fmt"
	"slices"
)

type NumberContainers struct {
	values  map[int][]int // value -> indices
	indices map[int]int   // index -> value
}

func Constructor() NumberContainers {
	return NumberContainers{
		values:  make(map[int][]int),
		indices: make(map[int]int),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	oldNumber, ok := nc.indices[index]

	// insert the number to the index
	nc.indices[index] = number
	if ok { // remove the index from the old number
		for i, v := range nc.values[oldNumber] {
			if v == index {
				nc.values[oldNumber] = append(nc.values[oldNumber][:i], nc.values[oldNumber][i+1:]...)
				break
			}
		}
	}

	if len(nc.values[number]) == 0 {
		nc.values[number] = []int{index}
		return
	}

	insertTo, _ := slices.BinarySearch(nc.values[number], index)
	nc.values[number] = slices.Insert(nc.values[number], insertTo, index)
}

func (nc *NumberContainers) Find(number int) int {
	indices, ok := nc.values[number]
	if ok && len(indices) > 0 {
		return indices[0]
	}

	return -1
}

func main() {
	example1()
	failedTest1()
}

func failedTest1() {
	nc := Constructor()

	nc.Change(1, 10)
	if nc.values[10][0] != 1 {
		fmt.Println(nc.values[10])
		panic("Test failed")
	}

	if nc.Find(10) != 1 {
		fmt.Println(nc.Find(10))
		panic("Test failed")
	}

	nc.Change(1, 20)
	if nc.values[20][0] != 1 {
		fmt.Println(nc.values[20])
		panic("Test failed")
	}

	if nc.Find(10) != -1 {
		fmt.Println(nc.Find(10))
		panic("Test failed")
	}

	if nc.Find(20) != 1 {
		fmt.Println(nc.Find(20))
		panic("Test failed")
	}

	if nc.Find(30) != -1 {
		fmt.Println(nc.Find(30))
		panic("Test failed")
	}

	fmt.Println("All tests passed")
}

func example1() {
	nc := Constructor()
	if nc.Find(10) != -1 {
		panic("Test failed")
	}
	nc.Change(2, 10)
	if nc.values[10][0] != 2 {
		fmt.Println(nc.values[10])
		panic("Test failed")
	}

	nc.Change(1, 10)
	if nc.values[10][0] != 1 {
		fmt.Println(nc.values[10])
		panic("Test failed")
	}
	nc.Change(3, 10)
	if nc.values[10][0] != 1 {
		fmt.Println(nc.values[10])
		panic("Test failed")
	}
	nc.Change(5, 10)
	if nc.values[10][0] != 1 {
		fmt.Println(nc.values[10])
		panic("Test failed")
	}
	if nc.Find(10) != 1 {
		panic("Test failed")
	}
	nc.Change(1, 20)
	if nc.values[20][0] != 1 {
		fmt.Println(nc.values[20])
		panic("Test failed")
	}
	if nc.Find(10) != 2 {
		fmt.Println(nc.Find(10))
		panic("Test failed")
	}

	fmt.Println("All tests passed")
}
