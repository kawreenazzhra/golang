package main

import (
	"fmt"
)

type Admin struct {
	Username string
	Password string
}

type Petugas struct {
	ID   int
	Nama string
}

type Transaksi struct {
	ID             int
	JenisKendaraan string
	PlatKendaraan  string
	JamMasuk       int
	JamKeluar      int
	Harga          int
}

var admins = [2]Admin{
	{"k4rin", "r4wrr"},
	{"cy4can", "rr4wrr"},
}

const maxPetugas = 100
const maxTransaksi = 100

var petugasList [maxPetugas]Petugas
var transaksiList [maxTransaksi]Transaksi

var petugasCount int
var transaksiCount int

var nextPetugasID = 1
var nextTransaksiID = 1

func main() {

	var pilihan int

	for {
		fmt.Println("======== PROGRAM PARKIR =======")
		fmt.Println("1. Login as Admin")
		fmt.Println("2. Login as User")
		fmt.Println("3. Exit")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var username, password string
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)

			if loginAdmin(username, password) {
				menuAdmin()
			} else {
				fmt.Println("Login gagal!")
			}
		} else if pilihan == 2 {
			menuUser()
		} else if pilihan == 3 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func loginAdmin(username, password string) bool {
	for _, admin := range admins {
		if admin.Username == username && admin.Password == password {
			return true
		}
	}
	return false
}

func tambahPetugas() {
	if petugasCount >= maxPetugas {
		fmt.Println("Tidak bisa menambah petugas lagi, kapasitas penuh.")
		return
	}

	var id int
	fmt.Print("ID Petugas: ")
	fmt.Scan(&id)

	var Nama string
	fmt.Print("Nama Petugas: ")
	fmt.Scan(&Nama)

	petugas := Petugas{id, Nama}
	petugasList[petugasCount] = petugas
	petugasCount++
	fmt.Println("Petugas berhasil ditambahkan.")
}

func cariPetugas(id int) int {
	for i := 0; i < petugasCount; i++ {
		if petugasList[i].ID == id {
			return i
		}
	}
	return -1
}

func menghapus(id int) int {
	low := 0
	high := petugasCount-1
	for low <= high {
		mid := low + (high-low)/2
		if petugasList[mid].ID == id {
			return mid
		} else if petugasList[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func hapusPetugas(id int) bool {
	index := menghapus(id)
	if index != -1 {
		for i := index; i < petugasCount-1; i++ {
			petugasList[i] = petugasList[i+1]
		}
		petugasCount--
		return true
	}
	return false
}

func ubahPetugas(id int, nama string) bool {
	index := cariPetugas(id)
	if index != -1 {
		petugasList[index].Nama = nama
		return true
	}
	return false
}

func hitungHargaParkir(jenisKendaraan string, durasiJam int) int {
	switch jenisKendaraan {
	case "motor":
		return durasiJam * 2000
	case "mobil":
		return durasiJam * 3000
	default:
		return 0
	}
}

func tambahTransaksi(jenisKendaraan, PlatKendaraan string, jamMasuk, jamKeluar int) {
	if transaksiCount >= maxTransaksi {
		fmt.Println("Tidak bisa menambah transaksi lagi, kapasitas penuh.")
		return
	}

	var idPetugas int
	fmt.Print("ID Petugas: ")
	fmt.Scan(&idPetugas)

	indexPetugas := cariPetugas(idPetugas)
	if indexPetugas == -1 {
		fmt.Println("Petugas dengan ID tersebut tidak ditemukan.")
		return
	}

	durasi := jamKeluar - jamMasuk
	harga := hitungHargaParkir(jenisKendaraan, durasi)
	transaksi := Transaksi{nextTransaksiID, jenisKendaraan, PlatKendaraan, jamMasuk, jamKeluar, harga}
	transaksiList[transaksiCount] = transaksi
	transaksiCount++
	nextTransaksiID++
	fmt.Println("Transaksi berhasil ditambahkan.")
}

func cetakTransaksi() {
	for i := 1; i < transaksiCount; i++ {
		key := transaksiList[i]
		j := i - 1

		for j >= 0 && transaksiList[j].ID > key.ID {
			transaksiList[j+1] = transaksiList[j]
			j--
		}
		transaksiList[j+1] = key
	}

	for i := 0; i < transaksiCount; i++ {
		transaksi := transaksiList[i]
		fmt.Printf("ID: %d, Jenis : %s, Plat Nomor: %s, Jam Masuk: %d, Jam Keluar: %d, Harga: %d\n", transaksi.ID, transaksi.JenisKendaraan, transaksi.PlatKendaraan, transaksi.JamMasuk, transaksi.JamKeluar, transaksi.Harga)
	}
}

func totalHargaParkir() int {
	total := 0
	for i := 0; i < transaksiCount; i++ {
		total += transaksiList[i].Harga
	}
	return total
}

func menuAdmin() {
	var pilihan int
	for {
		fmt.Println("\nMenu Admin:")
		fmt.Println("1. Tambah Petugas Parkir")
		fmt.Println("2. Hapus Petugas Parkir")
		fmt.Println("3. Ubah Petugas Parkir")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Scanln()

		if pilihan == 1 {
			tambahPetugas()
		} else if pilihan == 2 {
			var id int
			fmt.Print("ID Petugas yang akan dihapus: ")
			fmt.Scan(&id)
			if hapusPetugas(id) {
				fmt.Println("Petugas berhasil dihapus.")
			} else {
				fmt.Println("Petugas tidak ditemukan.")
			}
		} else if pilihan == 3 {
			var id int
			var nama string
			fmt.Print("ID Petugas yang akan diubah: ")
			fmt.Scan(&id)
			fmt.Print("Nama Baru: ")
			fmt.Scan(&nama)
			if ubahPetugas(id, nama) {
				fmt.Println("Petugas berhasil diubah.")
			} else {
				fmt.Println("Petugas tidak ditemukan.")
			}
		} else if pilihan == 4 {
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuUser() {
	var pilihan int
	for {
		fmt.Println("\nMenu User:")
		fmt.Println("1. Input Transaksi Parkir")
		fmt.Println("2. Cetak Transaksi Parkir")
		fmt.Println("3. Kembali ke Menu Utama")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var jenisKendaraan, platKendaraan string
			var jamMasuk, jamKeluar int
			fmt.Print("Jenis Kendaraan (motor/mobil): ")
			fmt.Scan(&jenisKendaraan)
			fmt.Print("Plat Nomor Kendaraan: ")
			fmt.Scan(&platKendaraan)
			fmt.Print("Jam Masuk: ")
			fmt.Scan(&jamMasuk)
			fmt.Print("Jam Keluar: ")
			fmt.Scan(&jamKeluar)
			tambahTransaksi(jenisKendaraan, platKendaraan, jamMasuk, jamKeluar)
		} else if pilihan == 2 {
			cetakTransaksi()
		} else if pilihan == 3 {
			return
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
