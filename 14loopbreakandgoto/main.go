package main

import "fmt"

func main() {
	fmt.Println("Loops, Break and Goto in GoLang")

	days := []string{"sunday", "monday", "tuesday","saturday"}

	fmt.Println(days)

	for d:=0; d< len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougueValue := 1 

	for rougueValue < 10 {

		if rougueValue==5 {
			break
		}

		if rougueValue ==2 {
			goto lco
			
		}



		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	lco:
	fmt.Println("Jumping at lco")

}
