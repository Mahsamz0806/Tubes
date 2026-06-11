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
	var pilih int
	fmt.Println("====SELAMAT DATANG DI CLEARMIND CHATBOT====")
	fmt.Println("APA YANG BISA SAYA BANTU HARI INI?")
	fmt.Println("1. Produktifitas")
	fmt.Println("2. Kesehatan Mental")
	fmt.Scan(&pilih)
	if pilih == 1 {
		produktifitas(A)
	} else if pilih == 2 {
}
func produktifitas(A tabtugas) {
	var pilihan int
	fmt.Println("===HALO SELAMAT DATANG DI PEMBANTU PRODUKTIFITAS===")
	fmt.Println("APA YANG KAMI BISA BANTU HARI INI?")
	fmt.Println("1. Input Tugas")
	fmt.Println("2. Tugas Prioritas")
	fmt.Println("3. Cari Tugas")
	fmt.Println("0. Kembali")
	fmt.Scan(&pilihan)
	if pilihan == 1 {

	} else if pilihan == 2 {

	} else if pilihan == 3 {

	}else if pilihan == 4{
		func main
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
func tugasPrioritas(A tabtugas){
	var i, temp, j int
	i = 0
	for i < idx{
		j = 
	}
}

