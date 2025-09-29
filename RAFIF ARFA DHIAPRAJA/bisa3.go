package main

import "fmt"

func main() {
	var luas float64
	var r int
	
	fmt.Print("Masukkan jari jari lingkaran: ")
	fmt.Scan(&r)

	luas = 3.14 * float64(r) * float64(r)

	fmt.Printf("jari jari %.d luasnya %.2f", r, luas )
}