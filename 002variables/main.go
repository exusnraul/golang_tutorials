package main

import "fmt"

var Logintoken = "This is a Public Variable"
var logintoken1 = "This is not a Public Variable"

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("ariable is of type : %T \n", username)

	var isbool bool = true
	fmt.Println(isbool)
	fmt.Printf("ariable is of type : %T \n", isbool)

	var isvariable int = 266
	fmt.Println(isvariable)
	fmt.Printf("ariable is of type : %T \n", isvariable)

	var smallfloat float64 = 111.4345345
	fmt.Println(smallfloat)
	fmt.Printf("ariable is of type : %T \n", smallfloat)

	// defaul values and aliases

	var defaultvar int
	fmt.Println(defaultvar)
	fmt.Printf("Variable is of type : %T \n", defaultvar)

	//Implicit type
	var website = "google.com"
	fmt.Println(website)

	//Walrusoperator
	numberofuser := 332343234
	fmt.Println(numberofuser)

	fmt.Println(Logintoken)
	fmt.Printf("Variable is of type : %T \n", Logintoken)

	fmt.Println(logintoken1)
	fmt.Printf("Variable is of type : %T \n", logintoken1)

}
