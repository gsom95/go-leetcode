package main

import "fmt"

type OrderedStream struct {
	values []string
	curPtr int
	n      int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{values: make([]string, n), n: n}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.values[idKey-1] = value
	if this.curPtr == this.n || this.values[this.curPtr] == "" {
		return nil
	}
	res := []string{}
	for ; this.curPtr < this.n && this.values[this.curPtr] != ""; this.curPtr++ {
		res = append(res, this.values[this.curPtr])
	}
	return res
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */

func main() {
	testCase := []struct {
		id    int
		value string
	}{
		{3, "ccccc"},
		{1, "aaaaa"},
		{2, "bbbbb"},
		{5, "eeeee"},
		{4, "ddddd"},
	}
	obj := Constructor(len(testCase))
	for _, tc := range testCase {
		res := obj.Insert(tc.id, tc.value)
		fmt.Println(res)
	}
}
