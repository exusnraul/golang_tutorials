package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcoe to Study time")
	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createdate := time.Date(2020, time.November, 26, 19, 19, 0, 0, time.UTC)
	fmt.Println(createdate)
	fmt.Println(createdate.Format("01-02-2006 Monday"))

}
