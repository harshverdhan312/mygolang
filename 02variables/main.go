package main

import "fmt"

// number := 10 --> this is not allowed outside of any method
/*
var numberOfUsers = 1000
var numberOFUsers int = 1000

this is allowed for global declarations

there are some constants too

const LoginToken string = "login" --- LoginToken is publicily availble; if the first letter of vas is capital then public.
*/

func main() {
	var username string = "Harsh"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type: %T \n", smallVal)

	var smallFloat float64 = 255.45387742322556
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type: %T \n", smallFloat)

	// default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of Type: %T \n", anotherVariable)

	var anotherVariableString string
	fmt.Println(anotherVariableString)
	fmt.Printf("Variable is of Type: %T \n", anotherVariableString)

	//implicit type

	var website = "meetkats.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000.0  // walrus operator
	fmt.Println(numberOfUser)
}
