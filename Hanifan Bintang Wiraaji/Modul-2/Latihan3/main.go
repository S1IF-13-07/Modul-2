package main

import "fmt"

func main() {
	var (
		r, v float64
		p float64 = 3.14
	)

	fmt.Println("Hitung luas lingkaran")
	fmt.Print("Masukan jari-jari: ")
	fmt.Scanln(&r)
	v = p * r * r
	fmt.Printf("Luas lingkaran adalah %v", v)
}