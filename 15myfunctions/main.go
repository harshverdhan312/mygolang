package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in GO Lang")
	greeter()
	
	greeterTwo()

	result := adder(3,5)

	fmt.Println("Result is :", result)


	proResult, myMesssage := proAdder(2,3,4,56,7,8,4,3,35)

	fmt.Println("Pro Result is :", proResult)

	fmt.Println("Pro Message is :", myMesssage)
}


func greeter() {
	fmt.Println("Namaste for Go-Lang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

func adder(a int, b int) int {
	return a+b
}


func proAdder(values ...int) (int, string) {
	total := 0
	for _ ,val := range values{
		total += val
	}

	return total, "Hi, Pro Result Functions"
}


