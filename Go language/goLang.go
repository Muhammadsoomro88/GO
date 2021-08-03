package main

import (
	"fmt"
	"math"
	"strconv"
)

func getAdd(num1 int, num2 int) int {
	return num1 + num2
}

func greet(name string) string {
	return "Hello" + name
}

func findSqrt(num, index int) int {
	if index*index == num {
		return index
	} else {
		return findSqrt(num, index+1)
	}
}

//global scope
var company string = "Contour Software Pvt Ltd"

const (
	val1 = iota
	val2 = iota
	val3 = iota
)

//Student Struct
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

//Embedding-> Composition relation
type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoryCapacity int
	memoryType     string
}

type Computer struct { //Computer has a processor and Memory
	Processor //has a relationship
	Memory
	price int
}

func main() {
	fmt.Printf("Starting Go Language from today\n")
	var num1 int = 10
	var num2 int = 15
	if num1 < num2 {
		fmt.Print(num1, " is smaller than ", num2)
	} else {
		fmt.Print(num1, " is greater than ", num2)
	}
	fmt.Print("\n")
	fmt.Print(math.Floor(3.1), "\n")
	fmt.Print(math.Ceil(3.8), "\n")
	fmt.Print(math.Sqrt(9), "\n")

	//calling a int function
	fmt.Print(getAdd(1, 10))
	fmt.Print("\n")

	//calling a string parameterised function
	fmt.Print(greet("Muhammad"))
	fmt.Print("\n")

	//calling a Recursion function
	fmt.Print(findSqrt(16, 1))

	//%v shows the variable and %T shows the type of variabel..like int, string..
	var val string
	val = "abc"
	fmt.Printf("\n%v, %T\n", val, val)

	fmt.Print("\n", company)

	var intVal int = 5

	//type casting a variable in GO
	var foo string = strconv.Itoa(intVal)
	fmt.Printf("\n%v, %T", foo, foo)

	//float numbers
	i := 3.14
	j := 1.7e12
	k := 2.3e12
	fmt.Printf("\n%v , %T\n", i, i)
	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", k, k)

	//complex number
	var comp complex128 = 1 + 2i
	fmt.Printf("\n%v, %T\n", comp, comp)
	fmt.Printf("\n%v, %T\n", real(comp), real(comp))
	fmt.Printf("\n%v, %T\n", imag(comp), imag(comp))

	const name string = "M.Soomro"
	fmt.Printf("%v, %T\n", name, name)

	//user constant
	const (
		user   string = "Admin"
		Procut string = "Product"
	)
	fmt.Printf("\n%v, %T", user, user)
	fmt.Printf("\n%v, %T\n", Procut, Procut)

	fmt.Println(val1)
	fmt.Println(val2)
	fmt.Println(val3)

	//Array in Go
	myArr := [...]int{10, 20, 30} //if you dont want to specify your size of Array so add [...] 3 dots in the index..
	fmt.Println(myArr)
	fmt.Println(len(myArr))

	//copying an array  [:]-> complete Array, [2:]-> from 2 to complete Array, [3:]-> from 3rd element to complete Array,[:3]->up to last three
	newArr := myArr[:]
	fmt.Println("Copied Array is:")
	fmt.Print(newArr)

	fmt.Println("\nShrinked Array is:")
	shrinkArr := myArr[1:]
	fmt.Println(shrinkArr)

	//3x3 Array
	threeArr := [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}
	fmt.Println("3x3 Array is:")
	fmt.Println(threeArr[1][1]) //we can access this array by setting a dimensions like [i][j] -> ith row & jth column

	//Slices
	Arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\nSlices:-")
	fmt.Println(Arr2)

	//copying slices
	Arr3 := Arr2
	fmt.Println("Copyed Slice is:-")
	fmt.Println(Arr2)
	fmt.Println(Arr3)
	Arr3[0] = 4
	fmt.Println(Arr2)
	fmt.Println(Arr3) //single change will affect both the Arrays

	//Slice another Syntax
	var Arr4 []int = make([]int, 3, 10) //size=3, capacity=10
	fmt.Println((len(Arr4)))
	fmt.Println((cap(Arr4)))

	//append
	Arr5 := append(Arr4, 5) //copy the existing array along with the appended element 5 with it
	fmt.Println(Arr5)

	//MAPS -> key value pair, string is key & int is value
	shoppingCart := map[string]int{
		"Keyboard": 100,
		"Mouse":    50,
		"Laptop":   1000,
	}
	fmt.Println(shoppingCart)
	fmt.Println(len(shoppingCart))
	fmt.Println(shoppingCart["Keyboard"])
	shoppingCart["Monitor"] = 1200
	fmt.Println(shoppingCart)
	delete(shoppingCart, "Monitor")
	fmt.Println(shoppingCart)

	//creating an object of an struct
	std := Student{
		name:     "Muhammad",
		rollNo:   3742,
		subjects: []string{"FYP", "ICC", "ISPM", "IOT", "RM"},
	}

	fmt.Println(std.name)
	fmt.Println(std.rollNo)
	fmt.Println(std.subjects[1])

	//object of Computer struct
	computer := Computer{}
	computer.price = 50000
	computer.processorName = "Intel Core i7"
	computer.cores = 6
	computer.memoryCapacity = 8
	computer.memoryType = "DDR4"
	fmt.Println(computer)

	//for loop syntax
	fmt.Println("\nInside For Loop")
	for v := range shoppingCart {
		fmt.Println(v)
	}
}
