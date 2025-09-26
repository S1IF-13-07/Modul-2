package main

import "fmt"

func main(){
	var farenheit, celcius float64
	fmt.Print("Masukan Suhu farenheit: ")
	fmt.Scan(&farenheit)
	celcius = (farenheit-32)*5/9
	fmt.Printf("Maka Hasil Suhu farenheit: %f", celcius)
}