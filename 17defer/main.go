package main

import "fmt"

func main() {

	defer fmt.Println("Hello World 1")
	defer fmt.Println("Hello World 2")
	fmt.Println("defer")

	myDefer()
}


func myDefer()  {

	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
	
}