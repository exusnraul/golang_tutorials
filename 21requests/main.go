package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("We will use get post")
	Performgetrequest()
	Performpostjsonrequest()
	Prformpostformrequest()
}

func Performgetrequest() {
	const myurl = "http://localhost:1111/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

	//othrway
	var responsestring strings.Builder
	bytecount, _ := responsestring.Write(content)
	fmt.Println(bytecount)
	fmt.Println(responsestring.String())
}

func Performpostjsonrequest() {
	const myurl = "http://localhost:1111/post"

	//fakejsn payload
	requestBody := strings.NewReader(`
	{
		"coursename":"Lets go with goLang",
		"platform":"Youtube"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	} else {
		var responsestring strings.Builder
		fmt.Println(responsestring.String())
	}
	defer response.Body.Close()
}

func Prformpostformrequest() {
	const myurl = "http://localhost:1111/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname", "Rahul")
	data.Add("lastname", "Saha")
	data.Add("email", "Saha@go.com")
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(content))

}
