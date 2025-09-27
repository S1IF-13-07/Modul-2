package main

import "fmt"

func main() {
	var nama string

	fmt.Print("Masukan Nama Anda : ")
	fmt.Scan(&nama)
	fmt.Printf("Halo, %s",nama)
}