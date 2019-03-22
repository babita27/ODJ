package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var EvenNum [5]int
	EvenNum[0] = 0
	slice1 := make([]int, 5)
	sliced := EvenNum[0:]
	numSlice := []int{2, 2, 3, 4, 5}
	numSlice := append(numSlice, 3, 4)
	//copy(slice1, numSlice)
	fmt.Println(slice1)
	fmt.Println(numSlice)
	//fmt.Println(apSlice)
	fmt.Println(sliced)
}
