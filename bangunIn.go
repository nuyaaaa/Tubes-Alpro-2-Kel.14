package main

import "fmt"

const NMAX int = 100

type kontak struct {
	telepon, email, lokasi string
}

type layanan struct {
	jenisMaterial string
	rating        float64
	riwayatOrder  int
}

type supplier struct {
	ID            int
	namaPT        string
	detailKontak  kontak
	detailLayanan layanan
}

type databaseSupplier [NMAX]supplier

func main() {
	var db databaseSupplier
	var n int = 0
	var menu bool = true

	isiDataDummy(&db, &n)

	for menu {
		menu = menuUtama(&db, &n)
	}
}

//database yang dimiliki diawal
func isiDataDummy(db *databaseSupplier, n *int) {
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

	db[*n] = supplier{6, "PT_Cat_Cerah", kontak{"061-666-123", "order@catcerah.com", "Bekasi"}, layanan{"Cat", 4.9, 150}}
	*n++

	db[*n] = supplier{7, "PT_Keramik_Indah", kontak{"021-555-789", "cs@keramikindah.com", "Tangerang"}, layanan{"Keramik", 4.3, 40}}
	*n++

	db[*n] = supplier{8, "PT_Besi_Kokoh", kontak{"021-888-123", "marketing@besikokoh.com", "Medan"}, layanan{"Besi", 4.7, 95}}
	*n++

	db[*n] = supplier{9, "PT_Pipa_Aliran", kontak{"021-777-123", "pipa@aliran.com", "Depok"}, layanan{"Pipa", 4.5, 75}}
	*n++

	db[*n] = supplier{10, "PT_Genteng_Kuat", kontak{"0251-333-444", "support@gentengkuat.com", "Bogor"}, layanan{"Genteng", 4.6, 55}}
	*n++
}

//tampilan untuk mengoperasikan program
func menuUtama(db *databaseSupplier, n *int) bool {
	var input string

	fmt.Println("_____________________________________________________")
	fmt.Println("|                                                   |")
	fmt.Println("|                  SELAMAT DATANG                   |")
	fmt.Println("|                    DI BANGUNIN                    |")
	fmt.Println("|___________________________________________________|")
	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("        BANGUNIN - DATABASE SUPPLIER MATERIAL        ")
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
	fmt.Print("Pilih menu (1-8): ")

	fmt.Scan(&input)

	switch input {
	case "1":
		tambahDataSupplier(db, n)
	case "2":
		ubahDataSupplier(db, *n)
	case "3":
		hapusDataSupplier(db, n)
	case "4":
		tampilkanDataSupplier(*db, *n)
	case "5":
		pencarianDataSupplier(*db, *n)
	case "6":
		pengurutanDataSupplier(db, *n)
	case "7":
		tampilkanStatistik(*db, *n)
	case "8":
		fmt.Println("Terima kasih telah menggunakan BANGUNIN.")
		return false
	}
	return true
}

//menambahkan data baru ke database dengan jumlah maks 100
func tambahDataSupplier(db *databaseSupplier, n *int) {
	if *n >= NMAX {
		fmt.Println("\n[GAGAL] PENUH EUY! Tidak dapat menambahkan data supplier baru.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Println("                TAMBAH DATA SUPPLIER")
	fmt.Println("=====================================================")

	fmt.Print("ID Supplier: ")
	fmt.Scan(&db[*n].ID)

	//nambahin kondisi dimana klo id yang di inputkan(baru) sama/udh ada di database, maka disuruh input ulang
	if adaID(db, *n, db[*n].ID) {
		fmt.Println("\n[GAGAL] ID Supplier sudah digunakan. Silakan masukkan ID yang berbeda.")
		return
	}
	var idBaru int
	isDuplicate := true

	for isDuplicate {
		fmt.Print("ID Supplier: ")
		fmt.Scan(&idBaru)

		isDuplicate = false

		for i := 0; i < *n && !isDuplicate; i++ {
			if db[i].ID == idBaru {
				isDuplicate = true // Penanda diubah jika ditemukan ID yang sama
			}
		}

		if isDuplicate {
			fmt.Println("[ERROR] ID Supplier sudah terdaftar! Silakan masukkan ID yang berbeda.")
		} else {
			db[*n].ID = idBaru
		}
	}

	db[*n].namaPT = mintaInputHuruf("Nama Perusahaan: ")

	db[*n].detailKontak.telepon = inputValidTelepon("Nomor Telepon: ")

	fmt.Print("Email: ")
	fmt.Scan(&db[*n].detailKontak.email)

	db[*n].detailKontak.lokasi = mintaInputHuruf("Lokasi: ")
	db[*n].detailLayanan.jenisMaterial = mintaInputHuruf("Jenis Material: ")

	db[*n].detailLayanan.rating = mintaInputFloat("Rating (0.0 - 5.0): ")
	db[*n].detailLayanan.riwayatOrder = mintaInputAngka("Riwayat Order: ")

	*n++

	fmt.Println("\n[BERHASIL] Data supplier berhasil ditambahkan!")
}

func adaID(db *databaseSupplier, n int, id int) bool {
	for i := 0; i < n; i++ {
		if db[i].ID == id {
			return true
		}
	}
	return false
}

//mengubah data supplier yang sudah ada berdasarkan id
func ubahDataSupplier(db *databaseSupplier, n int) {
	var idCari, i int
	var idxFound int

	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa diubah.")
		return
	}

	fmt.Print("\nMasukkan ID Supplier yang ingin diubah: ")
	fmt.Scan(&idCari)

	//pke metode sequential search buat nyari id
	idxFound = -1
	i = 0
	for i < n && idxFound == -1 {
		if db[i].ID == idCari {
			idxFound = i
		}
		i++
	}

	if idxFound == -1 {
		fmt.Println("\n[GAGAL] Data supplier dengan ID tersebut tidak ditemukan.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Printf("           UBAH DATA SUPPLIER (ID: %d)               \n", idCari)
	fmt.Println("=====================================================")

	db[idxFound].namaPT = mintaInputHuruf("Nama Perusahaan Baru: ")

	db[idxFound].detailKontak.telepon = inputValidTelepon("Nomor Telepon Baru: ")

	fmt.Print("Email Baru: ")
	fmt.Scan(&db[idxFound].detailKontak.email)

	db[idxFound].detailKontak.lokasi = mintaInputHuruf("Lokasi Baru: ")
	db[idxFound].detailLayanan.jenisMaterial = mintaInputHuruf("Jenis Material Baru: ")

	db[idxFound].detailLayanan.rating = mintaInputFloat("Rating Baru (0.0 - 5.0): ")
	db[idxFound].detailLayanan.riwayatOrder = mintaInputAngka("Riwayat Order Baru: ")

	fmt.Println("\n[BERHASIL] Data supplier berhasil diubah!")
}

//menghapus data yang udh ada di database, nyari data yang mau di ubah dari id nya
func hapusDataSupplier(db *databaseSupplier, n *int) {
	var cariID, idx, i int

	if *n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa dihapus.")
		return
	}
	fmt.Println("\n=====================================================")
	fmt.Println(" 		          HAPUS DATA SUPPLIER                 ")
	fmt.Println("=====================================================")
	fmt.Print("Masukkan ID Supplier yang ingin dihapus: ")
	fmt.Scan(&cariID)

	cariID = mintaInputAngka("Masukkan ID Supplier yang ingin dihapus: ")

	//ini juga pke sequential search
	idx = -1
	i = 0
	for i < *n && idx == -1 {
		if db[i].ID == cariID {
			idx = i
		}
		i++
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

//nampilin semua data supplier yang udh ada di database
func tampilkanDataSupplier(db databaseSupplier, n int) {
	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong.")
		return
	}

	fmt.Println("\n=======================================================================================================================================")
	fmt.Printf("%-5s | %-20s | %-15s | %-25s | %-15s | %-15s | %-6s | %-15s\n", "ID", "Nama PT", "Telepon", "Email", "Lokasi", "Jenis Material", "Rating", "Riwayat Order")
	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%-5d | %-20s | %-15s | %-25s | %-15s | %-15s | %-6.1f | %-15d\n",
			db[i].ID, db[i].namaPT, db[i].detailKontak.telepon, db[i].detailKontak.email, db[i].detailKontak.lokasi, db[i].detailLayanan.jenisMaterial, db[i].detailLayanan.rating, db[i].detailLayanan.riwayatOrder,
		)
	}
	fmt.Println("=======================================================================================================================================")
}

func pencarianDataSupplier(db databaseSupplier, n int) {
	var pilihMetode int
	var findData string
	var found bool = false
	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa dicari.")
		return
	}

	fmt.Println("\n=====================================================")
	fmt.Println("               PENCARIAN DATA SUPPLIER ")
	fmt.Println("=====================================================")
	fmt.Println("1. Cari berdasarkan Lokasi")
	fmt.Println("2. Cari berdasarkan Jenis Material")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Println("Pilih metode pencarian (1-3): ")
	fmt.Scan(&pilihMetode)
	fmt.Println("Pilih metode pencarian (1-2): ")

	pilihMetode = mintaInputAngka("Pilih metode pencarian (1-2): ")

	if pilihMetode < 1 || pilihMetode > 3 {
		fmt.Println("\n[INFO] Pilihan metode tidak valid. Silakan pilih 1, 2, atau 3.")
		pencarianDataSupplier(db, n)
		return
	}

	switch pilihMetode {
	case 1:
		findData = mintaInputHuruf("Masukkan Lokasi yang ingin dicari: ")
		fmt.Println("\n-- HASIL PENCARIAN --")

		//sequntial search buat nyari data yang lokasi nya sesuai sama inputan user
		for i := 0; i < n; i++ {
			if db[i].detailKontak.lokasi == findData {
				fmt.Printf("ID: %d | Nama PT: %s | Telepon: %s | Email: %s | Jenis Material: %s | Rating: %.1f | Riwayat Order: %d\n",
					db[i].ID, db[i].namaPT, db[i].detailKontak.telepon, db[i].detailKontak.email, db[i].detailLayanan.jenisMaterial, db[i].detailLayanan.rating, db[i].detailLayanan.riwayatOrder)
				found = true
			}
		}

		if !found {
			fmt.Println("Data supplier dengan lokasi tersebut tidak ditemukan.")
		}

	case 2:
		findData = mintaInputHuruf("Masukkan Jenis Material yang ingin dicari: ")
		fmt.Println("\n-- HASIL PENCARIAN --")

		//binarysearch buat nyari data yang jenis material nya sesuai sama inputan user
		insertSort(&db, n)

		//masuk ke binary search
		found = false
		left := 0
		right := n - 1
		mid := (left + right) / 2
		for left <= right && !found {
			if db[mid].detailLayanan.jenisMaterial == findData {
				found = true
			} else if db[mid].detailLayanan.jenisMaterial < findData {
				left = mid + 1
			} else {
				right = mid - 1
			}
			mid = (left + right) / 2
		}
		if found {
			fmt.Printf("ID: %d | Nama PT: %s | Telepon: %s | Email: %s | Lokasi: %s | Rating: %.1f | Riwayat Order: %d\n",
				db[mid].ID, db[mid].namaPT, db[mid].detailKontak.telepon, db[mid].detailKontak.email, db[mid].detailKontak.lokasi, db[mid].detailLayanan.rating, db[mid].detailLayanan.riwayatOrder)
		} else {
			fmt.Println("Data supplier dengan jenis material tersebut tidak ditemukan.")
		}

	case 3:
		return
	}
}

//insertion sort desc buat ngurutin data biar kepake di binary
func insertSort(db *databaseSupplier, n int) {
	var pass, i int
	var temp supplier
	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = db[pass]
		for i > 0 && temp.detailLayanan.rating > db[i-1].detailLayanan.rating {
			db[i] = db[i-1]
			i--
		}
		db[i] = temp
		pass++
	}
}

func pengurutanDataSupplier(db *databaseSupplier, n int) {
	var passI, i, j, passS, idx int
	var tempI, tempS supplier

	fmt.Println("\n=====================================================")
	fmt.Println("           PENGURUTAN DATA SUPPLIER (RATING)         ")
	fmt.Println("=====================================================")
	fmt.Println("1. Mengurutkan data supplier berdasarkan rating tertinggi")
	fmt.Println("2. Mengurutkan data supplier berdasarkan rating terendah")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih metode pengurutan (1-3): ")
	var pilihMetode int
	fmt.Scan(&pilihMetode)

	if n == 0 {
		fmt.Println("\n[INFO] Data supplier masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	if pilihMetode < 1 || pilihMetode > 3 {
		fmt.Println("\n[INFO] Pilihan metode tidak valid. Silakan pilih 1, 2, atau 3.")
		pengurutanDataSupplier(db, n)
		return
	}

	switch pilihMetode {

	case 1:
		//insertion sort desc berdasarkan rating, klo rating sama, di urutin berdasarkan riwayat order terbanyak
		passI = 1
		for passI <= n-1 {
			i = passI
			tempI = db[passI]

			for i > 0 && prioritas(tempI, db[i-1]) {
				db[i] = db[i-1]
				i--
			}
			db[i] = tempI
			passI++
		}
		fmt.Println("\n[BERHASIL] Data supplier berhasil diurutkan berdasarkan rating tertinggi")
		tampilkanDataSupplier(*db, n)

	case 2:
		//selection sort asc berdasarkan rating, klo rating sama, di urutin berdasarkan riwayat order paling sedikit
		passS = 1
		for passS <= n-1 {
			idx = passS - 1
			j = passS
			for j < n {
				if prioritas(db[idx], db[j]) { // ini kondisi dimana klo ratingnya sama bakal masuk ke fungsi prioritas buat ngedahuluin yang riwayat order lebih sedikit
					idx = j
				}
				j++
			}
			tempS = db[passS-1]
			db[passS-1] = db[idx]
			db[idx] = tempS
			passS++
		}
		fmt.Println("\n[BERHASIL] Data supplier berhasil diurutkan berdasarkan rating terendah")
		tampilkanDataSupplier(*db, n)

	case 3:
		return
	}
}

func prioritas(kandidat, pembanding supplier) bool {
	if kandidat.detailLayanan.rating != pembanding.detailLayanan.rating {
		return kandidat.detailLayanan.rating > pembanding.detailLayanan.rating
	}
	return kandidat.detailLayanan.riwayatOrder > pembanding.detailLayanan.riwayatOrder
}

func tampilkanStatistik(db databaseSupplier, n int) {
	var totalRating float64 = 0
	var avgRating float64
	var idxWilayah int
	var wilayahUnik [NMAX]string
	var jumlahPerWilayah [NMAX]int
	var jumlahWilayahUnik int = 0
	var lokasiSupplier string
	var ada bool

	fmt.Println("\n=====================================================")
	fmt.Println("   	  STATISTIK WILAYAH & KEPUASAN MITRA          ")
	fmt.Println("=====================================================")

	for i := 0; i < n; i++ {
		totalRating = totalRating + db[i].detailLayanan.rating
	}
	avgRating = totalRating / float64(n)
	fmt.Printf("Rata-rata Skor Rating Mitra: %.2f\n", avgRating)

	for i := 0; i < n; i++ {
		lokasiSupplier = db[i].detailKontak.lokasi
		ada = false
		idxWilayah = -1

		j := 0
		for j < jumlahWilayahUnik && !ada {
			if wilayahUnik[j] == lokasiSupplier {
				ada = true
				idxWilayah = j
			}
			j++
		}

		if ada {
			jumlahPerWilayah[idxWilayah]++
		} else {
			wilayahUnik[jumlahWilayahUnik] = lokasiSupplier
			jumlahPerWilayah[jumlahWilayahUnik] = 1
			jumlahWilayahUnik++
		}
	}

	fmt.Println("\nJumlah Supplier per Wilayah:")
	for i := 0; i < jumlahWilayahUnik; i++ {
		fmt.Printf("- %s : %d supplier\n", wilayahUnik[i], jumlahPerWilayah[i])
	}
}

//klo masukin sesuatu ga sesuai sama tipe datanya, nanti bikin error handling
// Validasi Huruf: Memastikan input tidak mengandung angka
func mintaInputHuruf(pesan string) string {
	var input string
	var valid bool = false

	for !valid {
		fmt.Print(pesan)
		fmt.Scan(&input)
		valid = true

		for i := 0; i < len(input); i++ {
			if input[i] >= '0' && input[i] <= '9' {
				valid = false
			}
		}

		if !valid {
			fmt.Println("   [ERROR] Input tidak boleh mengandung angka! Coba lagi.")
		}
	}
	return input
}

// Validasi Angka: input tidak boleh huruf
func mintaInputAngka(pesan string) int {
	var input string
	var valid bool = false
	var hasil int

	for !valid {
		fmt.Print(pesan)
		fmt.Scan(&input)
		valid = true
		hasil = 0

		for i := 0; i < len(input); i++ {
			if input[i] < '0' || input[i] > '9' {
				valid = false
			} else {
				hasil = hasil*10 + int(input[i]-'0')
			}
		}

		if !valid {
			fmt.Println("   [ERROR] Input harus berupa angka penuh! Coba lagi.")
		}
	}
	return hasil
}

// Validasi Telepon: input harus angka
func inputValidTelepon(pesan string) string {
	var input string
	var valid bool = false

	for !valid {
		fmt.Print(pesan)
		fmt.Scan(&input)
		valid = true

		for i := 0; i < len(input); i++ {
			if (input[i] >= 'a' && input[i] <= 'z') || (input[i] >= 'A' && input[i] <= 'Z') {
				valid = false
			}
		}

		if !valid {
			fmt.Println("   [ERROR] Nomor telepon tidak boleh mengandung huruf! Coba lagi.")
		}
	}
	return input
}

// Validasi Float / koma
func mintaInputFloat(pesan string) float64 {
	var input string
	var hasil float64
	var valid bool = false

	for !valid {
		fmt.Print(pesan)
		fmt.Scan(&input)

		valid = true
		hasil = 0.0
		var pembagi float64 = 10.0
		var adaTitik bool = false

		for i := 0; i < len(input); i++ {
			if input[i] == '.' {
				if adaTitik {
					valid = false
				}
				adaTitik = true
			} else if input[i] >= '0' && input[i] <= '9' {
				digit := float64(input[i] - '0')

				if !adaTitik {
					hasil = (hasil * 10.0) + digit
				} else {
					hasil = hasil + (digit / pembagi)
					pembagi = pembagi * 10.0
				}
			} else {
				valid = false
			}
		}
		if valid && adaTitik && hasil >= 0.0 && hasil <= 5.0 {
			valid = true
		} else {
			valid = false
			fmt.Println("   [ERROR] Input wajib menggunakan format desimal berpola titik (contoh: 4.0)! Coba lagi.")
		}
	}
	return hasil
}
