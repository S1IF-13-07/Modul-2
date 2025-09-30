package main

import (
	"fmt"
)

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari lingkaran :")
	fmt.Scanln(&r)
	Pi := 3.14 
	luas:=Pi*r*r
	fmt.Printf("Luas Lingkaran adalah %.1f\n", luas)
}