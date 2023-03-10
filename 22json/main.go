package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("We are Learning about JSON Data")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Python", 299, "Freecodecamp", "abc123", []string{"web-dev", "python"}},
		{"MERN", 299, "Freecodecamp", "abc1234", []string{"web-dev", "MERN", "full-stack"}},
		{"Angular", 299, "Freecodecamp", "abc12345", []string{"web-dev", "Angular"}},
		{"React", 299, "Freecodecamp", "abc1234567", nil},
	}
	//package as jsondata
	finalJson, _ := json.Marshal(lcoCourses)
	finalJson2, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
	fmt.Println(string(finalJson2))
}

func DecodeJson() {
	jsonfromWeb := []byte(`
	{
		"coursename": "Python",
		"Price": 299,
		"website": "Freecodecamp",
		"tags": [
				"web-dev",
				"python"
		]
}
	`)

	var lcocourse course

	checkValid := json.Valid(jsonfromWeb)
	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonfromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json Not valid ")
	}

	//somecases where adding data to key value pair

	var myonlinedata map[string]interface{}
	json.Unmarshal(jsonfromWeb, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for key, val := range myonlinedata {
		fmt.Printf("key is %v and val is %v and type is %T\n", key, val, val)
	}

}
