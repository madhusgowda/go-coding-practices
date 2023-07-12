package main

import (
	"fmt"
)

func main() {
	example1()
	example2()
}

func example1() {
	println("Example 1")
	slice1 := []int{1, 2}
	fmt.Println("Slice1 1st data:", slice1[:1])
	slice2 := []int{3, 4}
	slice3 := slice1
	copy(slice3, slice2)
	//slice1 = slice2
	fmt.Println("Slices: ", slice1, slice2, slice3)
}

func example2() {
	println("\nExample 2")
	primeNums := [6]int{2, 3, 5, 7, 11, 13}
	slice := primeNums[1:4]
	fmt.Println("Prime Num slices", slice)
}
