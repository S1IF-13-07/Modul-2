package main
import "fmt"
func main(){
	var f int
	fmt.Print("masukan fahreinheit:") 
	fmt.Scan(&f)
	c := (f - 32) * 5 / 9
	fmt.Printf("celciusnya adalah: %d\n", c)
}