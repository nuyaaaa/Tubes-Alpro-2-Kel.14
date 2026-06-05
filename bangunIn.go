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
	var n int = 0
	var menu bool = true

	isiDataDummy(&db, &n)

	for menu {
		menu = menuUtama(&db, &n)
	}
}

func isiDataDummy(db *DatabaseSUpplier, n *int) {
	db[*n] = supplier{1, "PT_Cahaya", kontak{"021-098-010", "email@.com", "Jakarta"}, layanan{"Semen", 4.5, 16}}
	*n++

	db[*n] = supplier{2, "PT_Baja_Utama", kontak{"021-555-234", "info@bajautama.com", "Jakarta"}, layanan{"Baja", 4.7, 60}}
	*n++

	db[*n] = supplier{3, "PT_Kayu_Abadi", kontak{"022-444-123", "sales@kayuabadi.com", "Bandung"}, layanan{"Kayu", 4.5, 30}}
	*n++

	db[*n] = supplier{4, "PT_Pasir_Mas", kontak{"031-333-123", "admin@pasirmas.com", "Surabaya"}, layanan{"Pasir", 4.6, 80}}
	*n++

	db[*n] = supplier{5, "PT_Bata_Merah", kontak{"024-222-123", "halo@batamerah.com", "Semarang"}, layanan{"Bata", 4.4, 110}}
	*n++

	db[*n] = supplier{6, "PT_Cat_Cerah", kontak{"061-666-123", "order@catcerah.com", "Medan"}, layanan{"Cat", 4.9, 150}}
	*n++

	db[*n] = supplier{7, "PT_Keramik_Indah", kontak{"021-555-789", "cs@keramikindah.com", "Tangerang"}, layanan{"Keramik", 4.3, 40}}
	*n++

	db[*n] = supplier{8, "PT_Besi_Kokoh", kontak{"021-888-123", "marketing@besikokoh.com", "Bekasi"}, layanan{"Besi", 4.7, 95}}
	*n++

	db[*n] = supplier{9, "PT_Pipa_Aliran", kontak{"021-777-123", "pipa@aliran.com", "Depok"}, layanan{"Pipa", 4.5, 75}}
	*n++

	db[*n] = supplier{10, "PT_Genteng_Kuat", kontak{"0251-333-444", "support@gentengkuat.com", "Bogor"}, layanan{"Genteng", 4.6, 55}}
	*n++
}

func menuUtama(db *DatabaseSUpplier, n *int) bool {
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
		tambahDataSupplier(db, n)
	case 2:
		ubahDataSupplier(db, *n)
	case 3:
		hapusDataSupplier(db, n)
	case 4:
		tampilkanDataSupplier(*db, *n)
	case 5:
		pencarianDataSupplier(*db, *n)
	case 6:
		pengurutanDataSupplier(db, *n)
	case 7:
		tampilkanStatistik(*db, *n)
	case 8:
		fmt.Println("Terima kasih telah menggunakan BANGUNIN.")
		return false
	default:
		fmt.Println("Pilihan tidak valid. Silakan pilih menu yang tersedia.")
	}
	return true
}

func tambahDataSupplier(db *DatabaseSUpplier, n *int) {
	if *n >= NMAX {
		fmt.Println("\n[GAGAL] PENUH EUY! Tidak dapat menambahkan data supplier baru.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Println(" 		        TAMBAH DATA SUPPLIER                  ")
	fmt.Println("=====================================================")

	db[*n].ID = *n + 1
	fmt.Printf("ID Supplier: %d\n", db[*n].ID)

	fmt.Print("Nama Perusahaan: ")
	fmt.Scan(&db[*n].NamaPT)

	fmt.Print("Nomor Telepon: ")
	fmt.Scan(&db[*n].DetailKontak.Telepon)

	fmt.Print("Email: ")
	fmt.Scan(&db[*n].DetailKontak.Email)

	fmt.Print("Lokasi: ")
	fmt.Scan(&db[*n].DetailKontak.Lokasi)

	fmt.Print("Jenis Material: ")
	fmt.Scan(&db[*n].DetailLayanan.JenisMaterial)

	fmt.Print("Rating (0.0 - 5.0): ")
	fmt.Scan(&db[*n].DetailLayanan.Rating)

	fmt.Print("Riwayat Order: ")
	fmt.Scan(&db[*n].DetailLayanan.RiwayatOrder)

	*n++

	fmt.Println("\n[BERHASIL] Data supplier berhasil ditambahkan!")
}

func ubahDataSupplier(db *DatabaseSUpplier, n int) {
	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa diubah.")
		return
	}

	var idCari int
	fmt.Print("\nMasukkan ID Supplier yang ingin diubah: ")
	fmt.Scan(&idCari)

	var idxFound int = -1
	for i := 0; i < n; i++ {
		if db[i].ID == idCari {
			idxFound = i
			break
		}
	}

	if idxFound == -1 {
		fmt.Println("\n[GAGAL] Data supplier dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Printf("           UBAH DATA SUPPLIER (ID: %d)               \n", idCari)
	fmt.Println("=====================================================")

	fmt.Print("Nama Perusahaan Baru: ")
	fmt.Scan(&db[idxFound].NamaPT)

	fmt.Print("Nomor Telepon Baru: ")
	fmt.Scan(&db[idxFound].DetailKontak.Telepon)

	fmt.Print("Email Baru: ")
	fmt.Scan(&db[idxFound].DetailKontak.Email)

	fmt.Print("Lokasi Baru: ")
	fmt.Scan(&db[idxFound].DetailKontak.Lokasi)

	fmt.Print("Jenis Material Baru: ")
	fmt.Scan(&db[idxFound].DetailLayanan.JenisMaterial)

	fmt.Print("Rating Baru (0.0 - 5.0): ")
	fmt.Scan(&db[idxFound].DetailLayanan.Rating)

	fmt.Print("Riwayat Order Baru: ")
	fmt.Scan(&db[idxFound].DetailLayanan.RiwayatOrder)

	fmt.Println("\n[BERHASIL] Data supplier berhasil diubah!")
}

func hapusDataSupplier(db *DatabaseSUpplier, n *int) {
	var cariID, idx int

	if *n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa dihapus.")
		return
	}
	fmt.Println("\n=====================================================")
	fmt.Println(" 		          HAPUS DATA SUPPLIER                 ")
	fmt.Println("=====================================================")
	fmt.Print("Masukkan ID Supplier yang ingin dihapus: ")
	fmt.Scan(&cariID)

	idx = -1
	for i := 0; i < *n; i++ {
		if db[i].ID == cariID {
			idx = i
			return
		}
	}

	if idx == -1 {
		fmt.Println("\n[GAGAL] Data supplier dengan ID tersebut tidak ditemukan.")
		return
	}

	for i := idx; i < *n-1; i++ {
		db[i] = db[i+1]
	}
	*n--

	fmt.Println("\n[BERHASIL] Data supplier berhasil dihapus!")
}

func tampilkanDataSupplier(db DatabaseSUpplier, n int) {
	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong.")
		return
	}

	fmt.Println("\n=======================================================================================================================================")
	fmt.Printf("%-5s | %-20s | %-15s | %-25s | %-15s | %-15s | %-6s | %-15s\n", "ID", "Nama PT", "Telepon", "Email", "Lokasi", "Jenis Material", "Rating", "Riwayat Order")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		s := db[i]
		fmt.Printf("%-5d | %-20s | %-15s | %-25s | %-15s | %-15s | %-6.1f | %-15d\n",
			s.ID,
			s.NamaPT,
			s.DetailKontak.Telepon,
			s.DetailKontak.Email,
			s.DetailKontak.Lokasi,
			s.DetailLayanan.JenisMaterial,
			s.DetailLayanan.Rating,
			s.DetailLayanan.RiwayatOrder,
		)
	}
	fmt.Println("=======================================================================================================================================")
}

func pencarianDataSupplier(db DatabaseSUpplier, n int) {

}

func pengurutanDataSupplier(db *DatabaseSUpplier, n int) {

}

func tampilkanStatistik(db DatabaseSUpplier, n int) {

}
