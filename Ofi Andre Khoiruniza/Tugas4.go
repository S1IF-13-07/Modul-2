package main

import "fmt"

func main() {
	var f int

	fmt.Print("masukkan nilai Fahrenheit = ")
	fmt.Scan(&f)

	c := (f-32)*9/5
	fmt.Println("dalam celcius adalah = ", c)
}