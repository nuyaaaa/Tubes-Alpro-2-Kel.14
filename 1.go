package main

import "fmt"

const NMAX = 1001
type identitas struct {
	nama string
	tahun int
	asal string
}
type tabPenduduk [NMAX]identitas

func main() {
	var n, targetTahun int
	var arr tabPenduduk
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].nama, &arr[i].tahun, &arr[i].asal)
	}
	fmt.Scan(&targetTahun)
	cariTarget(arr, n, targetTahun)
}

func cariTarget(arr tabPenduduk, n, X int) {
	var awal, akhir, tengah int
	var found bool
	awal = 0
	akhir = n - 1
	tengah = (awal + akhir) / 2
	for awal <= akhir && !found {
		if X == arr[tengah].tahun {
			found = true
		} else if X < arr[tengah].tahun {
			akhir = tengah - 1
		} else {
			awal = tengah + 1
		}
		tengah = (awal + akhir) / 2
	}
	if found {
		fmt.Print(arr[tengah].nama)
		fmt.Printf("\n%d", X)
		fmt.Printf("\n%s", arr[tengah].asal)
		fmt.Printf("\ndetemukan di index ke-%d", tengah)
	} else {
		fmt.Print("Data tidak ditemukan")
	}
}
