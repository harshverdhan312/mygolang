package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Mangoes", "Banana"}
	fmt.Println("fruits :", fruitList)
	fmt.Printf("type of fruit list is %T\n", fruitList)

	fruitList = append(fruitList, "Pineapple", "Peaches")
	fmt.Println("fruits :", fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 897
	highScores[2] = 764
	highScores[3] = 837
	//highScores[4] = 777
	fmt.Println(highScores)	

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))


	// how to remove a value from slices based on index

	var course = []string{"reactJs","Javascript","Swift","Python","Ruby"}
	fmt.Println(course)

	var index int = 2 

	course = append(course[:index],course[index+1:]...)

	fmt.Println(course)

}
