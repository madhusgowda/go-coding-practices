package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2}
	fmt.Println("Slice1 1st data:", slice1[:1])
	slice2 := []int{3, 4}
	slice3 := slice1
	copy(slice3, slice2)
	//slice1 = slice2
	fmt.Println("Slices: ", slice1, slice2, slice3)
}
