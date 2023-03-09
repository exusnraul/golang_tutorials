package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Structs")
	//no super parent child no inheritance
	Rahul := User{"Rahul", "rahul@g.com", true, 30}

	fmt.Println(Rahul)
	fmt.Printf("%+v\n", Rahul)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
