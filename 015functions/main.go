package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is : ", result)
	proresult, mymessage := proadder(2, 5, 5, 4, 8, 9, 9, 8, 7, 7, 4, 5)
	fmt.Println("ProResult is : ", proresult, mymessage)

}

func greeter() {
	fmt.Println("Namaste")
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proadder(value ...int) (int, string) {
	total := 0
	for _, val := range value {
		total += val
	}
	return total, "returning string"
}
