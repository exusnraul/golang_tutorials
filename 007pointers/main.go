package main

import "fmt"

// pointer is direct reference to memory address
func main() {
	fmt.Println("Welcome to Pointers")
	var ptr *int
	fmt.Println("The Value of Pointer is", ptr)
	mynumber := 23
	var ptr1 = &mynumber
	fmt.Println("The Value of Pointer is", ptr1)
	fmt.Println("The Value of Pointer is", *ptr1)

}
