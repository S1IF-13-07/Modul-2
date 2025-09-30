package main

import "fmt"

func main() {
	var fahrenheit int
	var celcius float64
	fmt.Print("Masukkan Suhu Farenheit : ")
	fmt.Scanln(&fahrenheit)
	celcius = float64(fahrenheit-32) * 5 / 9
	fmt.Printf("Suhu Celcius : %0.2f°C", celcius)
}
