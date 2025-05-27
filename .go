package main

import "fmt"

type NFT struct {
	namaAset string
	invest   float64
}

const NMAX int = 100

type tabNFT [NMAX]NFT

func main() {
	var pilihan int

	tampilMenu(pilihan)

}

func pilihanNFT() { // ini buat nampilin jenis jenis nftnya, pilihan ke 1
	fmt.Println("Daftar NFT:")
	fmt.Println("1. Bitcoin")
	fmt.Println("2. Ethereum")
	fmt.Println("3. Solana")

}

func nambahinNFT(A *tabNFT, N *int) { // buat nambahin nft, pilihan kedua
	fmt.Print("Masukan nama NFT : ")
	fmt.Scan(&(*A)[*N].namaAset)
	fmt.Print("Masukan jumlah investasi yang Anda inginkan : ")
	fmt.Scan(&(*A)[*N].invest)
	*N++
}

func TampilAset(A tabNFT, N int) {
	var i int
	for i = 0; i < N; i++ {
		fmt.Println("Aset Yang Anda miliki : ", A[i].namaAset, A[i].invest)
	}
}

func perBulan(A tabNFT, N int) float64 { // (function) menghitung investasi perbulan
	var i int

	for i = 0; i < N; i++ {
		if A[i].namaAset == "Bitcoin" {
			return A[i].invest * 105 / 100
		} else if A[i].namaAset == "Ethereum" {
			return A[i].invest * 104 / 100
		} else if A[i].namaAset == "Solana" {
			return A[i].invest * 103 / 100
		} else {
			return 0
		}
	}
	return 0
}

func tampilMenu(pilihan int) {
	var data tabNFT
	var nData int

	for pilihan != 7 {
		fmt.Println("SIMULASI INVESTASI NFT")
		fmt.Println("Menu:")
		fmt.Println("1. Daftar NFT")
		fmt.Println("2. Tambah NFT")
		fmt.Println("3. Hapus NFT")
		fmt.Println("4. Ubah NFT")
		fmt.Println("5. Hitung Investasi NFT")
		fmt.Println("6. Tampilkan Aset")
		fmt.Println("7. Selesai")
		fmt.Print("Pilih menu:")

		fmt.Scan(&pilihan)

		if pilihan == 1 {
			pilihanNFT()
		} else if pilihan == 2 {
			nambahinNFT(&data, &nData)
		} else if pilihan == 3 {
			// hapus nft
		} else if pilihan == 4 {
			// ubah nft
		} else if pilihan == 5 {

		} else if pilihan == 6 {
			TampilAset(data, nData)
		}
	}
	fmt.Println("Program Selesai!")
}

// mencari nft sequential dan binary search
func seqSearch(A *tabNFT, total int, x string) int {
	var i int
	i = 0
	for i < total && A[i].namaAset != x {
		i++
	}
	if i < total {
		return i // ditemukan
	} else {
		return -1 // tidak ditemukan
	}
}

func binSearch(A *tabNFT, total int, x float64) int { // binary search brdsaekan invest
	var left, right, mid int
	var ketemu int
	left = 0
	right = total
	ketemu = -1
	for left <= right && ketemu == -1 {
		mid = (left + right) / 2
		if x < (*A[mid]).invest {
			right = mid - 1
		} else if x > (*A[mid]).invest {
			left = mid + 1
		} else {
			ketemu = mid
		}
	}
	return ketemu
}

// mengurutkan nft selection dan insertion sort
func selectionSort(A *tabNFT, total int) {
	var pass, i, idx, temp int
	pass = 1
	for pass <= total-1 {
		idx = pass - 1
		i = pass
		for i < total {
			if (*A[i].invest) < (*A[idx].invest) {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		(*A[pass-1]) = (*A[idx])
		(*A[idx]) = temp
		pass++
	}
}

func insertionSort(A *tabNFT, total int) {
	var pass, i, temp int
	pass = 1
	for pass <= total-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.invest < (*A[i-1].invest) {
			(*A[i]) = (*A[i-1])
			i--
		}
		(*A[i]) = temp
		pass++
	}
}
