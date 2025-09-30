package main

import "fmt"

func main() {
	const phi = 3.14
	var jarijari float64
	fmt.Print("Masukkan jarijari lingkaran: ")
	fmt.Scan(&jarijari)

	hasil := phi * jarijari * jarijari
	fmt.Printf("luas lingkaran: ,%.1f", hasil)
}
