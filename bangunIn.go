package main

import "fmt"

const NMAX int = 100

type x struct {
}
type y [NMAX]x

func main() {
	menuUtama()
}

func menuUtama() {
	var input int

	fmt.Println("_____________________________________________________")
	fmt.Println("|                                                   |")
	fmt.Println("|                  SELAMAT DATANG                   |")
	fmt.Println("|                   DI BANGUNIN                     |")
	fmt.Println("|___________________________________________________|")
	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("        BANGUNIN - DATABASE SUPPLIER MATERIAL       ")
	fmt.Println("=====================================================")
	fmt.Println("1. Tambah Data Supplier")
	fmt.Println("2. Ubah Data Supplier")
	fmt.Println("3. Hapus Data Supplier")
	fmt.Println("4. Tampilkan Seluruh Data Supplier")
	fmt.Println("5. Pencarian Data Supplier")
	fmt.Println("6. Pengurutan Data Supplier")
	fmt.Println("7. Tampilkan Statistik Wilayah & Kepuasan Mitra")
	fmt.Println("8. Keluar")
	fmt.Println("=====================================================")
	fmt.Println("Pilih menu (1-8): ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		tambahDataSupplier()
	case 2:
		ubahDataSupplier()
	case 3:
		hapusDataSupplier()
	case 4:
		tampilkanDataSupplier()
	case 5:
		pencarianDataSupplier()
	case 6:
		pengurutanDataSupplier()
	case 7:
		tampilkanStatistik()
	case 8:
		fmt.Println("Terima kasih telah menggunakan BANGUNIN.")
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
		menuUtama()
	}
}

func tambahDataSupplier() {

}

func ubahDataSupplier() {

}

func hapusDataSupplier() {

}

func tampilkanDataSupplier() {

}

func pencarianDataSupplier() {

}

func pengurutanDataSupplier() {

}

func tampilkanStatistik() {

}
