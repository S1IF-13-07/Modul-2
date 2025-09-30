package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Printf("Nama anda %s", nama)
}