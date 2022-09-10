package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC) // 2022-09-10 17:48:56.1200948 +0700 WIB m=+0.000077701
	fmt.Printf("time2 %v\n", time2)                             // 2011-12-24 10:20:00 +0000 UTC

	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month(), "day:", now.Day()) // year: 2022 month: September day: 10

}
