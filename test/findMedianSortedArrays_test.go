package test

import (
	"fmt"
	"testing"
)

type sFindMedianSortedArraysData struct {
	Num1   []int
	Num2   []int
	Result float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{ 3}
	num2 := []int{-2, -1}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
