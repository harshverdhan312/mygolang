package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in go lang")

	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)


	// fairly simple ludo dice roller
	
	switch diceNumber {
	case 1:
		fmt.Println("Number 1 came, you can open your move")
	case 2:
		fmt.Println("Number 2 came, move 2 places")
	case 3:
		fmt.Println("Number 3 came, move 3 places")
	case 4:
		fmt.Println("Number 4 came, move 4 places")
	case 5:
		fmt.Println("Number 5 came, move 5 places")
	case 6:
		fmt.Println("Number 6 came, move 6 places and roll the dice again")
	default:
		fmt.Println("LUDO !!")
	}
}
