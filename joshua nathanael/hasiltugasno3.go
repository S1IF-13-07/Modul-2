package main
import "fmt"

func main() { 
	var r, luas float64 
	var Pi float64 = 3.14
	fmt.Print("MASUKAN JARI JARI LINGKARAN")
    fmt.Scanln(&r)
	luas = Pi * r * r
	fmt.Printf("Luas lingkaran = %v",luas)
}