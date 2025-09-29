package main

import "fmt"

func main () {
    var (
        nama, nim, kelas string
    )
        
    fmt.Print("Masukan Nama : ")
    fmt.Scanln(&nama)

    fmt.Print("Masukan NIM : ")
    fmt.Scanln(&nim)
    
    fmt.Print("Masukan Kelas : ")
    fmt.Scanln(&kelas)

    fmt.Println("Perkenalkan Nama saya "+nama+"", "Saya Mahasiswa Prodi S1 Dari "+kelas+"", "Dengan NIM "+""+nim)
}
