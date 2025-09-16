package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GOLANG")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.December, 03, 23, 23, 0, 0, time.UTC)
	fmt.Println("created date is :", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
