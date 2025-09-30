package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	// Input jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Hitung luas (π * r^2)
	luas := math.Pi * r * r

	// Output
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
}

