package main

import "fmt"

func main() {
	var r float64
	phi := 3.14 
	
	fmt.Print("masukkan jari-jari = ")
	fmt.Scan(&r)
	
	luas := phi*r*r

	fmt.Printf("luas nya adalah = %.1f\n", luas)
}