package main

import "fmt"

func main(){
	var r, phi float64
	phi = 3.14
	fmt.Print("Masukan Jari-Jari Lingkaran: ")
	fmt.Scanln(&r)
	luas := phi * r * r
	fmt.Printf("Maka Luas Lingkaran Adalah %.2f", luas)
}