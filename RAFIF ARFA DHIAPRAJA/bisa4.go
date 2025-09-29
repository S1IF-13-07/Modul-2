package main

import "fmt"

func main() {
	var F int
	var C float64

	fmt.Print("Masukkan suhu dalam fahrenheit: ")
	fmt.Scan(&F)

	C = float64(F-32) * 5 / 9
	
	fmt.Printf("Suhu dalam celcius: %.2f", C)
}