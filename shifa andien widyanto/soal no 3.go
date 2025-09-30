package main

import "fmt"

func main() {
	var luas, r float64
	fmt.Scan(&r)
	luas = 3.14*r*r
	fmt.Printf("%.1f", luas)
}