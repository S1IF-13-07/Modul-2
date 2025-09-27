package main

import "fmt"

func main() {
	var F, C float64

	// Input dari user (suhu dalam Fahrenheit)
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&F)

	// Rumus konversi ke Celcius: C = (F - 32) * 5/9
	C = (F - 32) * 5 / 9

	// Output hasil
	fmt.Println("Suhu dalam Celcius adalah:", C)
}
