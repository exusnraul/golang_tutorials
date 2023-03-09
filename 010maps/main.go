package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcoe to Maps")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["PH"] = "PHP"

	fmt.Println("List of Languages", languages)
	fmt.Println("List of Languages", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of Languages", languages)

	//Loops

	for key, val := range languages {
		fmt.Printf("Key %v Value %v \n", key, val)
	}

}
