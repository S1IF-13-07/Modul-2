package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14 // pakai konstanta pi sederhana

	// Input jari-jari
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&r)

	// Hitung luas lingkaran
	luas := pi * r * r

	// Output hasil
	fmt.Println("Luas lingkaran =", luas)
}