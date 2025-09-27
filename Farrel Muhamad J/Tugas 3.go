package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	fmt.Print("Masukkan jari-jari lingkaran (dalam cm): ")
	fmt.Scanln(&jariJari)

	// Rumus luas lingkaran = π * r^2
	luas := math.Pi * math.Pow(jariJari, 2)

	fmt.Printf("Luas lingkaran dengan jari-jari %.2f cm adalah %.2f cm²\n", jariJari, luas)
}
