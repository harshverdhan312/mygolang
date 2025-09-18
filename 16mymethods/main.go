package main

import "fmt"

func main() {
	fmt.Println("Sturcts in GO")

	// no inheritence in GO LANG; no super or parent concepts in GO LANG.

	harsh := User{"Harsh", "harsh@go.dev", true, 21 }
	fmt.Println(harsh)
	fmt.Printf("Harsh details are: %+v\n", harsh) // common syntax
	fmt.Printf("Name is: %v and email is : %v\n", harsh.Name, harsh.Email)
	harsh.GetStatus()
	harsh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("IS user active:", u.Status)
}

func (u User) NewMail()  {
	u.Email = "test@go.dev"
	fmt.Println("new email is :", u.Email)
}

