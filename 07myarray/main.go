package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("Fruit List is:", fruitList)
	fmt.Println("Fruit List is:", len(fruitList))

	var vegList = [3]string{"potato","tomato","mushroom"}
	fmt.Println("Vegetables :", vegList)
	fmt.Println(len(vegList))

}
