package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Apple1"
	fruitList[3] = "Apple2"

	fmt.Println("Fruit List is : ", fruitList)
	fmt.Println("Fruit List Length is : ", fruitList)

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println(len(vegList), vegList)
}
