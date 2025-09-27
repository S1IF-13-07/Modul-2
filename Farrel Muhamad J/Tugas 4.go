package main

import (
	"fmt"
)

func main() {
	var fahrenheit int

	fmt.Print("Masukkan suhu dalam Fahrenheit (°F): ")
	fmt.Scanln(&fahrenheit)

	// Rumus konversi: C = (F - 32) * 5 / 9
	celsius := float64(fahrenheit-32) * 5 / 9

	fmt.Printf("Suhu %d°F sama dengan %.2f°C\n", fahrenheit, celsius)
}
