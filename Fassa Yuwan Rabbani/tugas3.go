package main

import (
	"fmt"
	"math"
)

func main() {
	var jarijari float64
	fmt.Print("Masukkan jari - jari Lingkaran : ")
	fmt.Scanln(&jarijari)
	var luas = math.Pi * jarijari * jarijari
	fmt.Printf("Luas Lingkaran %0.2f adalah %0.2f", jarijari, luas)
}
