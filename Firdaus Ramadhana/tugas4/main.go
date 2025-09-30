package main

import "fmt"

func main() {
	var f int
	fmt.Print("Masukkan suhu dalam Fahrenheit : ")
	fmt.Scanln(&f)
	c := (f - 32) * 5 / 9
	fmt.Printf("Suhu dalam Celcius: %d\n", c)
}
