package main

import "fmt"

func main() {
	fmt.Println("Welcome to LOOps")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v \n", index, day)
	// }

	roguevalue := 1

	for roguevalue < 10 {
		if roguevalue == 9 {
			break
		} else if roguevalue == 3 {
			roguevalue++
			continue
		} else if roguevalue == 6 {
			goto lco
		}
		fmt.Println(roguevalue)
		roguevalue++
	}

lco:
	fmt.Println("goto statement excuted")

}
