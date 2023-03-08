package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Rating:")

	//comma ok // err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("ariable is of type : %T \n", input)

}
