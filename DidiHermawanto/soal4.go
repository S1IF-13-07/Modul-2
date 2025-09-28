package main

import "fmt"

func main() {
	var F float64

	// Input suhu dalam Fahrenheit
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	fmt.Scanln(&F)

	// Rumus konversi ke Celcius
	C := (F - 32) * 5 / 9

	// Output hasil
	fmt.Printf("Suhu dalam Celcius adalah %.0f\n", C)
}
