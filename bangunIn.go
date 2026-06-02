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
	db[*jumlahData] = supplier{1, "PT_Cahaya", kontak{"021-098-010", "email@.com", "Jakarta"}, layanan{"Semen", 4.5, 16}}
	*jumlahData++

	db[*jumlahData] = supplier{2, "PT_Baja_Utama", kontak{"021-555-234", "info@bajautama.com", "Jakarta"}, layanan{"Baja", 4.7, 60}}
	*jumlahData++

	db[*jumlahData] = supplier{3, "PT_Kayu_Abadi", kontak{"022-444-123", "sales@kayuabadi.com", "Bandung"}, layanan{"Kayu", 4.5, 30}}
	*jumlahData++

	db[*jumlahData] = supplier{4, "PT_Pasir_Mas", kontak{"031-333-123", "admin@pasirmas.com", "Surabaya"}, layanan{"Pasir", 4.6, 80}}
	*jumlahData++

	db[*jumlahData] = supplier{5, "PT_Bata_Merah", kontak{"024-222-123", "halo@batamerah.com", "Semarang"}, layanan{"Bata", 4.4, 110}}
	*jumlahData++

	db[*jumlahData] = supplier{6, "PT_Cat_Cerah", kontak{"061-666-123", "order@catcerah.com", "Medan"}, layanan{"Cat", 4.9, 150}}
	*jumlahData++

	db[*jumlahData] = supplier{7, "PT_Keramik_Indah", kontak{"021-555-789", "cs@keramikindah.com", "Tangerang"}, layanan{"Keramik", 4.3, 40}}
	*jumlahData++

	db[*jumlahData] = supplier{8, "PT_Besi_Kokoh", kontak{"021-888-123", "marketing@besikokoh.com", "Bekasi"}, layanan{"Besi", 4.7, 95}}
	*jumlahData++

	db[*jumlahData] = supplier{9, "PT_Pipa_Aliran", kontak{"021-777-123", "pipa@aliran.com", "Depok"}, layanan{"Pipa", 4.5, 75}}
	*jumlahData++

	db[*jumlahData] = supplier{10, "PT_Genteng_Kuat", kontak{"0251-333-444", "support@gentengkuat.com", "Bogor"}, layanan{"Genteng", 4.6, 55}}
	*jumlahData++

	db[*jumlahData] = supplier{11, "PT_Kabel_Nusantara", kontak{"021-555-999", "kabel@nusantara.com", "Jakarta"}, layanan{"Kabel", 4.9, 200}}
	*jumlahData++

	db[*jumlahData] = supplier{12, "PT_Kaca_Bening", kontak{"031-333-777", "sales@kacabening.com", "Surabaya"}, layanan{"Kaca", 4.4, 35}}
	*jumlahData++

	db[*jumlahData] = supplier{13, "PT_Paku_Bumi", kontak{"061-666-888", "paku@bumi.com", "Medan"}, layanan{"Paku", 4.2, 120}}
	*jumlahData++

	db[*jumlahData] = supplier{14, "PT_Beton_Perkasa", kontak{"0411-444-555", "beton@perkasa.com", "Makassar"}, layanan{"Beton", 4.8, 180}}
	*jumlahData++

	db[*jumlahData] = supplier{15, "PT_Aluminium_Jaya", kontak{"0711-333-222", "alum@jaya.com", "Palembang"}, layanan{"Aluminium", 4.5, 65}}
	*jumlahData++

	db[*jumlahData] = supplier{16, "PT_Semen_Sentosa", kontak{"031-398-111", "info@semensentosa.com", "Gresik"}, layanan{"Semen", 4.7, 140}}
	*jumlahData++

	db[*jumlahData] = supplier{17, "PT_Baja_Sakti", kontak{"0254-333-222", "baja@sakti.com", "Cilegon"}, layanan{"Baja", 4.6, 85}}
	*jumlahData++

	db[*jumlahData] = supplier{18, "PT_Kayu_Lestari", kontak{"024-354-111", "kontak@kayulestari.com", "Semarang"}, layanan{"Kayu", 4.4, 50}}
	*jumlahData++

	db[*jumlahData] = supplier{19, "PT_Pasir_Putih", kontak{"0721-444-333", "pasir@putih.com", "Lampung"}, layanan{"Pasir", 4.3, 70}}
	*jumlahData++

	db[*jumlahData] = supplier{20, "PT_Bata_Ringan", kontak{"031-894-111", "bata@ringan.com", "Sidoarjo"}, layanan{"Bata", 4.7, 130}}
	*jumlahData++

	db[*jumlahData] = supplier{21, "PT_Cat_Utama", kontak{"031-532-444", "sales@catutama.com", "Surabaya"}, layanan{"Cat", 4.5, 90}}
	*jumlahData++

	db[*jumlahData] = supplier{22, "PT_Keramik_Mulia", kontak{"021-296-444", "mulia@keramik.com", "Jakarta"}, layanan{"Keramik", 4.8, 165}}
	*jumlahData++

	db[*jumlahData] = supplier{23, "PT_Besi_Baja", kontak{"022-601-222", "admin@besibaja.com", "Bandung"}, layanan{"Besi", 4.6, 105}}
	*jumlahData++

	db[*jumlahData] = supplier{24, "PT_Pipa_Jaya", kontak{"021-824-111", "info@pipajaya.com", "Bekasi"}, layanan{"Pipa", 4.4, 55}}
	*jumlahData++

	db[*jumlahData] = supplier{25, "PT_Genteng_Metal", kontak{"0264-201-222", "metal@genteng.com", "Purwakarta"}, layanan{"Genteng", 4.5, 40}}
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
	fmt.Printf("ID Supplier: %d\n", db[*jumlahData].ID)

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

func hapusDataSupplier(db *DatabaseSUpplier, jumlahData *int) {

}

func tampilkanDataSupplier(db DatabaseSUpplier, jumlahData int) {

}

func pencarianDataSupplier(db DatabaseSUpplier, jumlahData int) {

}

func pengurutanDataSupplier(db *DatabaseSUpplier, jumlahData int) {

}

func tampilkanStatistik(db DatabaseSUpplier, jumlahData int) {

}
