package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Webrequest")
	response, err := http.Get(url)
	checkNillErr(err)
	fmt.Printf("Type of response %T \n", response)
	defer response.Body.Close() //callers responsibility to close every request

	databytes, err := ioutil.ReadAll(response.Body)
	checkNillErr(err)
	fmt.Print(string(databytes))
	writetofile(string(databytes))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}

func writetofile(content string) {
	file, err := os.Create("../response.html")
	checkNillErr(err)
	length, err := io.WriteString(file, content)
	checkNillErr(err)
	fmt.Println(length)
	file.Close()
}
