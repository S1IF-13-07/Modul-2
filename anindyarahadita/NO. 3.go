package main
import "fmt"

func main() {
	var luas, pi, r float64
	fmt.Scan(&r, &pi)
	pi = 3.14
	r = 14
	luas = pi*r*r
	fmt.Println("Luas lingkaran adalah: ", luas)
}