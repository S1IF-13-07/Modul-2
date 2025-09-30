package main

import "fmt"

func main () {
	var F, C float64

	fmt.Print("Masukan suhu dalam Fahrenheit: ")
	fmt.Scanln(&F)
	C = (F - 32) * 5 / 9
	fmt.Println("Suhu dalam Celcius adalah:", C)
}