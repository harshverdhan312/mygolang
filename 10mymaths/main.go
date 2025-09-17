package main

import (
	"fmt"
	"math/big"
	//"math/rand/v2"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths in Go ")

	//var mynumberOne int = 2
	//var mynumberTwo float64 = 4.5

	//fmt.Println("The sum is:", mynumberOne+int(mynumberTwo))

	// random numbers in GO
	
	// fmt.Println(rand.IntN(5))

	//random from Crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
