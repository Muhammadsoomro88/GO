package main

import "fmt"

func findMax(myArr [5]int) int {
	max := 0
	for i := 0; i < len(myArr); i++ {
		if myArr[i] > max {
			max = myArr[i]
		}
	}
	return max
}

func findMin(myArr [5]int) int {
	min := 0
	for i := 0; i < len(myArr); i++ {
		if myArr[i] < min {
			min = myArr[i]
		}
	}
	return min
}

func main() {
	myArr := [5]int{0, 5, 2, 4, 3}
	fmt.Println("Finding Maximum in Array:")
	fmt.Println("Maximum Element in Array: ", findMax(myArr))
	fmt.Println("Minimum Element in Array: ", findMin(myArr))
}
