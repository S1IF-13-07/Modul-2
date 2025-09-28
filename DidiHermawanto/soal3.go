package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14

	// input jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// hitung luas
	luas := pi * r * r

	// output hasil
	fmt.Println("Luas lingkaran adalah", luas)
}
