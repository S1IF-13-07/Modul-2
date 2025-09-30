package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var jumlahMahasiswa int

	fmt.Println("Ceritanya Program Biodata Mahasiswa Coyyy")
	fmt.Println("==========================================")

	fmt.Print("Masukkan jumlah mahasiswa yang ingin didata: ")
	fmt.Scanln(&jumlahMahasiswa)

	if jumlahMahasiswa <= 0 {
		fmt.Println("Jumlah mahasiswa harus lebih dari 0!")
		return
	}

	nama := make([]string, jumlahMahasiswa)
	nim := make([]string, jumlahMahasiswa)
	kelas := make([]string, jumlahMahasiswa)

	fmt.Printf("\nMasukkan %d data mahasiswa (format: Nama NIM Kelas)\n", jumlahMahasiswa)
	fmt.Println("Contoh: Budi 123456789 IF-A")
	fmt.Println()

	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("Mahasiswa %d: ", i+1)
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) >= 3 {
			nama[i] = parts[0]
			nim[i] = parts[1]
			kelas[i] = parts[2]
		} else {
			fmt.Printf("Format input salah! Silakan masukkan ulang data mahasiswa %d\n", i+1)
			i--
		}
	}

	fmt.Println("\n==================================")
	fmt.Println("         HASIL KELUARAN")
	fmt.Println("==================================")
	fmt.Printf("Total mahasiswa yang didata: %d\n", jumlahMahasiswa)
	fmt.Println()

	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("Mahasiswa %d:\n", i+1)
		fmt.Printf("Masukan: %s %s %s\n", nama[i], nim[i], kelas[i])
		fmt.Printf("Keluaran: Perkenalan saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s hehe.\n",
			nama[i], kelas[i], nim[i])
		fmt.Println()
	}

	var pilihan string
	fmt.Print("\nApakah ingin melihat ringkasan data? (y/n): ")
	fmt.Scanln(&pilihan)

	if strings.ToLower(pilihan) == "y" {
		fmt.Println("\n==================================")
		fmt.Println("        RINGKASAN DATA")
		fmt.Println("==================================")
		fmt.Printf("Total mahasiswa: %d\n", jumlahMahasiswa)

		kelasCount := make(map[string]int)
		for i := 0; i < jumlahMahasiswa; i++ {
			kelasCount[kelas[i]]++
		}

		fmt.Println("\nDistribusi per kelas:")
		for k, v := range kelasCount {
			fmt.Printf("- Kelas %s: %d mahasiswa\n", k, v)
		}
	}

	fmt.Println("\nJangan Lupa Sholat")
}
