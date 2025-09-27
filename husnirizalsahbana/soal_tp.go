package main

import "fmt"

func main() {
	nama := ""
	fmt.Print("Masukan nama kalian : ")
	fmt.Scan(&nama)
	fmt.Printf("Hallo %s", nama)
}
