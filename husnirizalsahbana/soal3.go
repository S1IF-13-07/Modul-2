package main

import (
	"fmt"
	"math"
)

func main() {
	var r, luas float64

	// Input jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Hitung luas
	luas = math.Pi * r * r

	// Output hasil
	fmt.Printf("Luas lingkaran dengan jari-jari %.1f adalah %.1f\n", r, luas)
}
