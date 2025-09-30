package main

import "fmt"

func main() {
	fahrenheit := readFahrenheit()
	celsius := toCelsius(fahrenheit)
	fmt.Printf("Suhu dalam Celsius: %.0f\n", celsius)
}

func readFahrenheit() float64 {
	fmt.Print("Masukkan suhu dalam Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanln(&fahrenheit)
	return fahrenheit
}

func toCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
