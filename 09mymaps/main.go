package main

import "fmt"

func main() {
	fmt.Println("Maps in go")


	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["GO"] = "GoLang"


	fmt.Println("List of all languages:", languages)

	fmt.Println("JS shorts for:",languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages:", languages)

	// loops are intresting in GOLANG

	for key, value := range languages {
		fmt.Printf("for Key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("for Key v, value is %v\n", value)
	}
}
