package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	radius, err := readRadius()
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca jari-jari:", err)
		return
	}

	if err := validateRadius(radius); err != nil {
		fmt.Println(err)
		return
	}

	area := circleArea(radius)
	fmt.Printf("Luas lingkaran: %.2f\n", area)
}

func readRadius() (float64, error) {
	fmt.Print("Masukkan jari-jari: ")
	var radius float64
	if _, err := fmt.Scanln(&radius); err != nil {
		return 0, err
	}
	return radius, nil
}

func validateRadius(radius float64) error {
	if radius <= 0 {
		return errors.New("jari-jari harus lebih besar dari nol")
	}
	return nil
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
