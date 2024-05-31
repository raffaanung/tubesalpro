package main

import (
	"fmt"
)

const (
	maxPetugas   = 100
	maxTransaksi = 1000
	mobilBiaya   = 5000
	motorBiaya   = 2000
)

type Petugas struct {
	id       string
	nama     string
	password string
}

type Transaksi struct {
	jenisKendaraan string
	biaya          int
}

var petugasArray [maxPetugas]Petugas
var transaksiArray [maxTransaksi]Transaksi
var totalPetugas int
var totalTransaksi int

func main() {
	var pilihan int
	for {
		fmt.Println("Aplikasi Parkir")
		fmt.Println("1. Login sebagai Admin")
		fmt.Println("2. Login sebagai Pengguna")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			if loginAdmin() {
				adminMenu()
			}
		} else if pilihan == 2 {
			if loginPetugas() {
				penggunaMenu()
			}
		} else if pilihan == 3 {
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func loginAdmin() bool {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "admin" {
		return true
	} else {
		fmt.Println("Login gagal.")
		return false
	}
}

func loginPetugas() bool {
	var id, password string
	fmt.Print("ID Petugas: ")
	fmt.Scan(&id)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	index := sequentialSearchPetugas(id)
	if index != -1 && petugasArray[index].password == password {
		return true
	} else {
		fmt.Println("Login gagal.")
		return false
	}
}

func adminMenu() {
	var pilihan int
	for {
		fmt.Println("Menu Admin")
		fmt.Println("1. Tambah Petugas")
		fmt.Println("2. Ubah Petugas")
		fmt.Println("3. Hapus Petugas")
		fmt.Println("4. Logout")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPetugas()
		case 2:
			ubahPetugas()
		case 3:
			hapusPetugas()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func penggunaMenu() {
	var pilihan int
	for {
		fmt.Println("Menu Petugas")
		fmt.Println("1. Transaksi Parkir")
		fmt.Println("2. Laporan Harian")
		fmt.Println("3. Urutkan Transaksi (Ascending)")
		fmt.Println("4. Urutkan Transaksi (Descending)")
		fmt.Println("5. Logout")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			transaksiParkir()
		case 2:
			laporanHarian()
		case 3:
			selectionSortTransaksiAsc()
		case 4:
			insertionSortTransaksiDesc()
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
func tambahPetugas() {
	if totalPetugas >= maxPetugas {
		fmt.Println("Data petugas penuh.")
		return
	}

	var id, nama, password string
	fmt.Print("ID Petugas: ")
	fmt.Scan(&id)
	fmt.Print("Nama Petugas: ")
	fmt.Scan(&nama)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	petugasArray[totalPetugas] = Petugas{id, nama, password}
	totalPetugas++
	fmt.Println("Petugas berhasil ditambahkan.")
}

func ubahPetugas() {
	var id, nama, password string
	fmt.Print("ID Petugas yang akan diubah: ")
	fmt.Scan(&id)

	index := sequentialSearchPetugas(id)
	if index == -1 {
		fmt.Println("Petugas tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru: ")
	fmt.Scan(&nama)
	fmt.Print("Password baru: ")
	fmt.Scan(&password)

	petugasArray[index].nama = nama
	petugasArray[index].password = password
	fmt.Println("Petugas berhasil diubah.")
}

func hapusPetugas() {
	var id string
	fmt.Print("ID Petugas yang akan dihapus: ")
	fmt.Scan(&id)

	index := sequentialSearchPetugas(id)
	if index == -1 {
		fmt.Println("Petugas tidak ditemukan.")
		return
	}

	for i := index; i < totalPetugas-1; i++ {
		petugasArray[i] = petugasArray[i+1]
	}
	totalPetugas--
	fmt.Println("Petugas berhasil dihapus.")
}

func sequentialSearchPetugas(id string) int {
	for i := 0; i < totalPetugas; i++ {
		if petugasArray[i].id == id {
			return i
		}
	}
	return -1
}
func transaksiParkir() {
	if totalTransaksi >= maxTransaksi {
		fmt.Println("Data transaksi penuh.")
	}

	var jenisKendaraan string
	var waktu int
	fmt.Print("Jenis Kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Waktu: ")
	fmt.Scan(&waktu)

	biaya := 0
	if jenisKendaraan == "mobil" {
		if waktu > 1 {
			biaya = 7000 + (mobilBiaya * (waktu - 1))
		} else if waktu == 1 {
			biaya = mobilBiaya
		}
	} else if jenisKendaraan == "motor" {
		if waktu > 1 {
			biaya = 4000 + (motorBiaya * (waktu - 1))
		} else if waktu == 1 {
			biaya = motorBiaya
		}
	} else {
		fmt.Println("Jenis kendaraan tidak valid.")
	}

	transaksiArray[totalTransaksi] = Transaksi{jenisKendaraan, biaya}
	totalTransaksi++
	fmt.Println("Transaksi berhasil dicatat.")
}

func laporanHarian() {
	fmt.Println("Laporan Harian")
	fmt.Println("Kendaraan\tBiaya")

	totalPendapatan := 0
	for i := 0; i < totalTransaksi; i++ {
		fmt.Printf("%s\t%d\n", transaksiArray[i].jenisKendaraan, transaksiArray[i].biaya)
		totalPendapatan += transaksiArray[i].biaya
	}

	fmt.Printf("Total Pendapatan: %d\n", totalPendapatan)
}
func selectionSortTransaksiAsc() {
	for i := 0; i < totalTransaksi-1; i++ {
		minIndex := i
		for j := i + 1; j < totalTransaksi; j++ {
			if transaksiArray[j].biaya < transaksiArray[minIndex].biaya {
				minIndex = j
			}
		}
		transaksiArray[i], transaksiArray[minIndex] = transaksiArray[minIndex], transaksiArray[i]
	}
	fmt.Println("Transaksi berhasil diurutkan")
}

func insertionSortTransaksiDesc() {
	for i := 1; i < totalTransaksi; i++ {
		key := transaksiArray[i]
		j := i - 1
		for j >= 0 && transaksiArray[j].biaya < key.biaya {
			transaksiArray[j+1] = transaksiArray[j]
			j--
		}
		transaksiArray[j+1] = key
	}
	fmt.Println("Transaksi berhasil diurutkan")
}

func binarySearchTransaksi(biaya int) int {
	low, high := 0, totalTransaksi-1
	for low <= high {
		mid := (low + high) / 2
		if transaksiArray[mid].biaya == biaya {
			return mid
		} else if transaksiArray[mid].biaya < biaya {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
