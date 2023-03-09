package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to working with Files")
	content := "this will be going in a file"
	file, err := os.Create("../test.txt")
	checkNillErr(err)
	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println(length)
	file.Close()
	readFile("../test.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNillErr(err)
	fmt.Println("Text Data inside the file is \n", string(databyte))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
