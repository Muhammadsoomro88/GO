// Online IDE - Code Editor, Compiler, Interpreter

package main

import "fmt"

func getSqrt(num int, index int) int {
	if index*index != num {
		return getSqrt(num, index+1)
	} else {
		return index
	}
}

var company string = "Contour Software Pvt Ltd"

func main() {
	fmt.Println("Welcome to Online IDE!! Happy Coding :)")
	fmt.Print(getSqrt(16, 1))
	fmt.Println("\n", company)
}
