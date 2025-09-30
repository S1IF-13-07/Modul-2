package main

import "fmt"

func main() {
	var farenheit float64
	fmt.Print("Input suhu Farenheit : ")
	fmt.Scan(&farenheit)
	fmt.Printf("Hasil suhu ke celcius = %f", (farenheit-32)*5/9)
}
