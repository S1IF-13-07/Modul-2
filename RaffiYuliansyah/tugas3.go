package main

import "fmt"



func main(){
	const PI = 3.14159
	var jarijari float64
	fmt.Print("Input jari-jari : ")
	fmt.Scanf("%f", &jarijari)
	fmt.Printf("Luas lingkaran : %.1f", PI*jarijari*jarijari)


}