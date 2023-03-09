package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitlist = []string{"Apple", "Peach", "Banana"}
	fmt.Printf("the type is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Tomato")
	fmt.Println(fruitlist[1:3]) //slicing
	fmt.Println(fruitlist[:3])  //slicing

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 945
	highscore[2] = 465
	highscore[3] = 867

	highscore = append(highscore, 213, 213, 5443, 6546)

	fmt.Println(highscore)

	sort.Ints(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))

	//removing a value from slice

	var courses = []string{"react", "Python", "swift", "Ruby", "JS"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
