package main

import "fmt"

func main () {
	var Jari int
	phi :=  3.1415926535
	fmt.Print("Masukkan Jari Jari = ")
	fmt.Scan(&Jari)

	volume := (4.0 / 3.0) * phi * ( float64(Jari) * float64(Jari) * float64(Jari) )
	luas := 4 * phi * ( float64(Jari) * float64(Jari))

	fmt.Printf("Bola dengan Jari Jari %d memiliki volume %.4f dan luas kulit %.4f ", Jari, volume, luas)
}
