package main
import "fmt"
func main() {
 var ( 
 nama, kelas string
 nim int
 )
 fmt.Print("Masukan nama: ")
 fmt.Scanln(&nama)
 fmt.Print("Masukan kelas: ")
 fmt.Scanln(&kelas)
 fmt.Print("Masukan nim: ")
 fmt.Scanln(&nim)
 fmt.Printf("Hii, Perkenalkan saya %s mahasiswa TELKOM UNIVERSITY PURWOKERTO kelas %s dengan nim %d", nama,  kelas, nim)
}