package main

import "fmt"
var idx int
const NMAX int = 1000

type Tugas struct {
	namaTugas string
	prioritas int
	deadline  int
}
type tabtugas [NMAX]Tugas

func main() {
	var A tabtugas
	var tgl string
	fmt.Scan(&tgl)
	menuUtama(&A)
}
func menuUtama(A *tabtugas){
	fmt.Printf("===== ClearMind =====\n")
	fmt.Println("[1] Menu Produktivitas")
	fmt.Println("[2] Menu Kesehatan Mental")
	fmt.Println("[0] Keluar Progam")
	fmt.Println("Pilih 1 atau 2: ")
	
	var pilihan int
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
			produktifitas()
	case 2:
			MenuKesehatanMental()
	case 0:
			fmt.Print("Terima Kasih sudah menggunakan ClearMind. Semoga hari Anda menyenangkan.")
	}
}

func produktifitas(A tabtugas) {
	var pilihan int
	fmt.Println("===HALO SELAMAT DATANG DI PEMBANTU PRODUKTIFITAS===")
	fmt.Println("APA YANG KAMI BISA BANTU HARI INI?")
	fmt.Println("[1] Input Tugas")
	fmt.Println("[2] Tugas Prioritas")
	fmt.Println("[3] Cari Tugas")
	fmt.Println("[0] Kembali")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
			inputDataProduk()
	case 2:
			MenuPrioritas()
	case 3:

	case 0:
			menuUtama()
	}
}
func inputDataProduk(A *tabtugas) {
	var n int
	fmt.Println("==INPUT DATA PRODUKTIFITAS==")
	fmt.Println("BERAPA BANYAK DATA YANG INGIN ANDA INPUT")
	fmt.Scan(&n)
	fmt.Println("INPUT DATA ANDA SECARA BERURUTAN (NAMA, PRIORITAS, DEADLINE)")
	for idx < n+idx {
		fmt.Scan(&A[idx].namaTugas, &A[idx].prioritas, &A[idx].deadline)
		idx ++
	}
}
func menuPrioritas(A *tabtugas){
	var pilih int
	fmt.Println("==MENU PRIORITAS==")
	fmt.Println("1.Ascending")
	fmt.Println("2.Descending")
	fmt.Println("0.Kembali")
    switch pilih{
		case 1:
			tugasPrioritasAscending(A)
		case 2:
			tugasPrioritasDescending(A)
		case 0:
			produktivitas(A)
	}
}
func tugasPrioritasDescend(A tabtugas){
	var pass, temp, j, indeks int
	pass = 0
	for pass < idx{
		indeks = i
		j = pass +1
		for j < idx{
			if arr[j].prioritas>arr[indeks].prioritas{
				indeks = j
			}
		}
		temp = arr[pass]
		arr[pass] = arr[indeks]
		temp = arr[indeks]
	}
}
func tugasPrioritasAscend(A tabtugas){
	var pass, temp, j, indeks int
	pass = 0
	for pass < idx{
		indeks = i
		j = pass +1
		for j < idx{
			if arr[j].prioritas<arr[indeks].prioritas{
				indeks = j
			}
		}
		temp = arr[pass]
		arr[pass] = arr[indeks]
		temp = arr[indeks]
	}
}

func MenuKesehatanMental(){
	fmt.Printf("=== CEK KESEHATAN MENTAL ===\n")
	fmt.Println("APA YANG BISA KAMI BANTU HARI INI?")
	fmt.Println("[1] Koin Mental dan Stress Meter")
	fmt.Println("[2] Isi Jurnal Suasana Hati")
	fmt.Println("[3] Lihat Stastistik Mingguan")
	fmt.Println("[0] Kembali ke Menu Utama")
	
	var pilihan int
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
			fmt.Print("On Progress")
	case 2:
			fmt.Print("On Progress")
	case 3:
			fmt.Print("On Progress")
	case 0:
			menuUtama()
	}
}
