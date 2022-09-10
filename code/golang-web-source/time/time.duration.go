package main

import (
	"fmt"
	"time"
)

func main() {
	// menghitung time duration
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds: ", duration.Seconds())
	fmt.Println("time elapsed in minutes: ", duration.Minutes())
	fmt.Println("time elapsed in hours: ", duration.Hours())

	// menghitung durasi yang dilakukan pada objek waktu
	fmt.Println("===============================================")
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration2 := t2.Sub(t1)

	fmt.Println("time elapsed in second: ", duration2.Seconds())
	fmt.Println("time elapsed in minutes: ", duration2.Minutes())
	fmt.Println("time elapsed in hours: ", duration2.Hours())

	// konversi angka ke time.Duration
	// 12 * time.Minute             // 12 menit
	// 65 * time.Hour                 // 65 jam
	// 150000 * time.Milisecond     // 150k milidetik atau 150 detik
	// 45 * time.Microsecond         // 45 microdetik
	// 233 * time.Nanosecond         // 233 nano detik

}
