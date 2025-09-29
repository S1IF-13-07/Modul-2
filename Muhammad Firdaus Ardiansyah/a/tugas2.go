package main
import "fmt"
func main() {
	var nama, kelas string
	var nim int 
	fmt.Print("masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("masukan kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("masukan nim: ")
	fmt.Scanln(&nim)
	fmt.Printf("Perkenalkan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %d\n", nama, kelas, nim)
}