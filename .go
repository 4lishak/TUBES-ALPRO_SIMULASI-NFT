package main

import "fmt"

type NFT struct {
	namaAset         string
	kenaikanPerBulan float64
}

type jenisNFT []NFT

func main() {
	var pilihan int

	fmt.Println("SIMULASI INVESTASI NFT")
	fmt.Println("Menu:")
	fmt.Println("1. Daftar Aset NFT")
	fmt.Println("2. Tambah NFT")
	fmt.Println("3. Hapus NFT")
	fmt.Println("4. Ubah NFT")
	fmt.Println("5. Hitung Investasi NFT")
	fmt.Println("6. Selesai")
	fmt.Print("Pilih menu:")

	fmt.Scan(&pilihan)

	if pilihan == 1 {
		// panggil func pilihanNFT
		// gimana caranya
	} else if pilihan == 2 {
		// nambah nft
	} else if pilihan == 3 {
		// hapus nft
	} else if pilihan == 4 {
		// ubah nft
	} else if pilihan == 5 {
		var aset string
		var uang float64
		var bulan1, bulan3, bulan6, profit float64

		fmt.Println("Masukan Jenis Aset:")
		fmt.Scanln(&aset)

		fmt.Println("Masukkan Jumlah Investasi:")
		fmt.Scanln(&uang)

		keuntungan(aset, uang, &bulan1, &bulan3, &bulan6, &profit)

		fmt.Println("Hasil Investasi")
		fmt.Printf("Investasi setelah 1 bulan: %.2f\n", bulan1)
		fmt.Printf("Investasi setelah 3 bulan: %.2f\n", bulan3)
		fmt.Printf("Investasi setelah 6 bulan: %.2f\n", bulan6)
		fmt.Printf("Provit yang didapatkan: %.2f\n", profit)
	} else if pilihan == 6 {
		fmt.Println("Investasi selesai")
	} else {
		fmt.Println()
	}

}

func pilihanNFT(jenisNFT) { // ini buat nampilin jenis jenis nftnya, pilihan ke 1
	if jenisNFT == 0 {
		fmt.Println("Belum ada NFT")
	} else {
		fmt.Println("Daftar NFT:")
		fmt.Println("1. Bitcoin")  // kenaikannya 1.05 atau 5%
		fmt.Println("2. Ethereum") // naik 1.04 atau 4%
		fmt.Println("3. Solana")   // naik 1.03 atau 3%

	}

}

func nambahinNFT() { // buat nambahin nft, pilihan kedua

}

func hapusNFT() {

}

func mengubahNFT() {

}

func keuntungan(aset string, uang float64, bulan1, bulan3, bulan6, profit *float64) { // procedure buat itung keuntungannya
	var i int

	*bulan1 = perBulan(aset, uang)

	*bulan3 = uang
	for i = 0; i < 3; i++ {
		*bulan3 = perBulan(aset, *bulan3)
	}

	*bulan6 = uang
	for i = 0; i < 6; i++ {
		*bulan6 = perBulan(aset, *bulan6)
	}

	*profit = *bulan6 - uang
}
func perBulan(aset string, uang float64) float64 { // (function) menghitung investasi perbulan
	if aset == "Bitcoin" {
		return uang * 105 / 100
	} else if aset == "Ethereum" {
		return uang * 104 / 100
	} else if aset == "Solana" {
		return uang * 103 / 100
	} else {
		return 0
	}
}
