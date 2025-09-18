package main

import (
	"fmt"
	"io"
	"io/ioutil" // oldschool way
	"os"
)

func main() {
	fmt.Println("File handling")

	content := "This needs to go in a file - lets go, in GO!"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }

	checkNilErr(err)

	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./mygofile.txt")
	readFileInModernWay("./mygofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) // this is old school way

	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

func readFileInModernWay(filename string) {
	dataByte, err := os.ReadFile(filename) // modern replacement for ioutil.ReadFile
	checkNilErr(err)

	fmt.Println("Text data inside the file is:\n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
