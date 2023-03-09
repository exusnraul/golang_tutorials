package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=python&paymentid=asdahsd"

func main() {
	fmt.Println("Url Handling")
	fmt.Println((myurl))

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("query type %T\n", qparams)
	fmt.Println(qparams["coursename"])

	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=Rahul",
	}
	anotherurl := partsofUrl.String()
	fmt.Println(anotherurl)
}
