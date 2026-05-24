package main

import "fmt"

const NMAX int = 100

type kontak struct {
	Telepon, Email, Lokasi string
}

type layanan struct {
	JenisMaterial string
	Rating        float64
	RiwayatOrder  int
}

type supplier struct {
	ID            int
	NamaPT        string
	DetailKontak  kontak
	DetailLayanan layanan
}

type DatabaseSUpplier [NMAX]supplier

func main() {
	var db DatabaseSUpplier
	var jumlahData int = 0

	//paggil data dummy
	isiDataDummy(&db, &jumlahData)

	for {
		menu := menuUtama(&db, &jumlahData)
		if !menu {
			break
		}
	}

}

func isiDataDummy(db *DatabaseSUpplier, jumlahData *int) {
	db[*jumlahData] = supplier{
		//cara ngisinya ikut template ini {1, "PT_Cahaya", kontak{"021-098-010", "email@.com", "Jakarta"},layanan{ "Semen", 4.5, 16}}
	}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++

	db[*jumlahData] = supplier{}
	*jumlahData++
}

func menuUtama(db *DatabaseSUpplier, jumlahData *int) bool {
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
		tambahDataSupplier(db, jumlahData)
	case 2:
		ubahDataSupplier(db, *jumlahData)
	case 3:
		hapusDataSupplier(db, jumlahData)
	case 4:
		tampilkanDataSupplier(*db, *jumlahData)
	case 5:
		pencarianDataSupplier(*db, *jumlahData)
	case 6:
		pengurutanDataSupplier(db, *jumlahData)
	case 7:
		tampilkanStatistik(*db, *jumlahData)
	case 8:
		fmt.Println("Terima kasih telah menggunakan BANGUNIN.")
		return false
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
	}
	return true
}

func tambahDataSupplier(db *DatabaseSUpplier, jumlahData *int) {
	if *jumlahData >= NMAX {
		fmt.Println("\n[GAGAL] PENUH EUY! Tidak dapat menambahkan data supplier baru.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Println(" 		        TAMBAH DATA SUPPLIER                  ")
	fmt.Println("=====================================================")

	db[*jumlahData].ID = *jumlahData + 1
	fmt.Print("ID Supplier: %d\n", db[*jumlahData].ID)

	fmt.Print("Nama Perusahaan: ")
	fmt.Scan(&db[*jumlahData].NamaPT)

	fmt.Print("Nomor Telepon: ")
	fmt.Scan(&db[*jumlahData].DetailKontak.Telepon)

	fmt.Print("Email: ")
	fmt.Scan(&db[*jumlahData].DetailKontak.Email)

	fmt.Print("Lokasi: ")
	fmt.Scan(&db[*jumlahData].DetailKontak.Lokasi)

	fmt.Print("Jenis Material: ")
	fmt.Scan(&db[*jumlahData].DetailLayanan.JenisMaterial)

	fmt.Print("Rating (0.0 - 5.0): ")
	fmt.Scan(&db[*jumlahData].DetailLayanan.Rating)

	fmt.Print("Riwayat Order: ")
	fmt.Scan(&db[*jumlahData].DetailLayanan.RiwayatOrder)

	*jumlahData++

	fmt.Println("\n[BERHASIL] Data supplier berhasil ditambahkan!")
}

func ubahDataSupplier(db *DatabaseSUpplier, jumlahData int) {

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
