package main 
import "fmt"
 
func main() { 
    var ( 
        satu, dua, tiga string 
        temp string 
    ) 

		// memasukkan inputan pada var satu, dua, dan tiga
    fmt.Print("Masukan input string: ") 
    fmt.Scanln(&satu) 
    fmt.Print("Masukan input string: ") 
    fmt.Scanln(&dua) 
    fmt.Print("Masukan input string: ") 
    fmt.Scanln(&tiga) 

		// print var satu, dua, dan tiga
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga) 

		// menggunakan variabel swapping untuk membalik inputan
    temp = satu 
    satu = dua 
    dua = tiga 
    tiga = temp 

		// print hasil variabel swapping
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga) 
}
