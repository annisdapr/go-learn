package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	umur := flag.Int("umur", 20, "Masukkan umur Anda")
	kota := flag.String("kota", "Jakarta", "Masukkan kota asal Anda")

	// Parsing flags
	flag.Parse()

	// Ambil argument pertama sebagai nama (jika ada)
	var nama string
	if len(flag.Args()) > 3 {
		nama = flag.Arg(0)
	} else {
		fmt.Println("Usage: go run main.go <nama> -umur=<angka> -kota=<nama_kota>")
		return
	}

	// Output
	fmt.Printf("Halo, %s! Umur kamu %d tahun dan berasal dari %s.\n", nama, *umur, *kota)
}
