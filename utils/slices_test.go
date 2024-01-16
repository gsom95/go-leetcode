package utils

import (
	"log"
	"testing"
)

func TestEqualInt32Slices(t *testing.T) {
	log.Println("TestEqualInt32Slices is running")
	t.Run("Empty slices", emptySlices)
	t.Run("Left slice is empty, right slice is not", emptyLeftSlice)
	t.Run("nil slice", nilSlice)
}

func emptySlices(t *testing.T) {
	left := []int32{}
	right := []int32{}

	if !EqualInt32Slices(left, right) {
		t.Error("Empty slices should be equal")
	}
}

func emptyLeftSlice(t *testing.T) {
	left := []int32{}
	right := []int32{1}

	if EqualInt32Slices(left, right) {
		t.Error("Slices shouldn't be equal")
	}
}

func nilSlice(t *testing.T) {
	left := []int32{}
	var right []int32 = nil

	if !EqualInt32Slices(left, right) {
		t.Error("Problem")
	}
}
