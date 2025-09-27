package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nama, nim, kelas string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukin nama lo: ")
	nama, _ = reader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("Masukin NIM lo: ")
	nim, _ = reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	fmt.Print("Masukin kelas lo: ")
	kelas, _ = reader.ReadString('\n')
	kelas = strings.TrimSpace(kelas)

	fmt.Println("\n====================")
	fmt.Println(" Resume Anak Gaul ")
	fmt.Println("====================")
	fmt.Printf("Hai, gue %s nih!\n", nama)
	fmt.Printf("NIM gue: %s, jangan salah yaa~\n", nim)
	fmt.Printf("Sekarang gue nongkrong di kelas %s \n", kelas)
	fmt.Println("Jangan lupa kenalan ya, biar nggak cuma tau nama doang hehe ")
}
