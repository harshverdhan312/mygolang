package main

import "fmt"

func main() {
	fmt.Println("If-Else in GoLang")

	loginCount := 13

	var results string

	if loginCount < 10 {
		results = "regular user"
	} else if loginCount > 10 {
		results = "Watchout"
	} else {
		results = "exactly 10 login count"
	}

	fmt.Println(results)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
