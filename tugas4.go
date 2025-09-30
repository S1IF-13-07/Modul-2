package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Print("Masukkan suhu dalam satuan Fahrenheit: ")
	fmt.Scanln(&fahrenheit)
	c := (fahrenheit - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celsius adalah: %.2f\n", c)
}
