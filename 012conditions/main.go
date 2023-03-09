package main

import "fmt"

func main() {
	fmt.Println("Welcome to conditional statements")

	logincount := 10
	var result string

	if logincount < 10 {
		result = "Regular User"
	} else if logincount > 10 {
		result = "Wathout"
	} else {
		result = "exactly 10 Logins"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even NNumber")
	} else {
		fmt.Println("Odd Number")
	}

	if num := 3; num < 10 {
		fmt.Println("Num Less Than 10")
	} else {
		fmt.Println("Num is Not Less Than 10")
	}
}
