package main
import "fmt"

//Variabel global untuk diapaki semua function dan procedure
var idx, pilihan int

const NMAX int = 1000
const BatasKoinMental int = 30
const BatasStress int = 10 //Untuk sementara segini

type Mental struct{
	tanggal int
	skorEmosi int
	catatanEmosi string
	sisaKoin int //Untuk menyimpan data dari Array Tugas
	stressMeter int //Untuk menimpan data dari Array Tugas
}
type tabmental [NMAX]JurnalMental

type Tugas struct {
	tanggal int
	namaTugas string
	prioritas int
	deadline  int
}
type tabtugas [NMAX]Tugas

func main() {
	var A tabtugas
	var B tabmental
	var tglAktif int
	setTanggalSesi(&tanggalAktif)
	menuUtama(&A, &B)
}

func setTanggalSesi(tanggal *int) {//Tanggal untuk kedua menu, pada Produktivitas dan Mental
	fmt.Println("Masukkan Tanggal Sesi Ini (Format: DDMMYYYY, cth: 12062026): ")
	fmt.Scan(&tanggal)
	fmt.Printf("Program")
}

func menuUtama(A *tabtugas, B *tabJurnal, tglAktif *int){
	fmt.Printf("===== ClearMind =====\n")
	fmt.Print("[Tanggal Aktif Saat Ini: %d]\n", *tglAktif)//Untuk tanggal yang aktif pada kedua menu
	fmt.Println("[1] Menu Produktivitas")
	fmt.Println("[2] Menu Kesehatan Mental")
	fmt.Println("[0] Keluar Progam")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
			produktifitas(A)
	case 2:
			MenuKesehatanMental(A, B)
	case 0:
			fmt.Print("Terima Kasih sudah menggunakan ClearMind. Semoga hari Anda menyenangkan.")
	}
}

func produktifitas(A *tabtugas) {
	fmt.Println("===HALO SELAMAT DATANG DI PEMBANTU PRODUKTIVITAS===")
	fmt.Println("APA YANG KAMI BISA BANTU HARI INI?")
	fmt.Println("[1] Input Tugas")
	fmt.Println("[2] Tugas Prioritas")
	fmt.Println("[3] Cari Tugas")
	fmt.Println("[0] Kembali")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
			inputDataTugas(A)
	case 2:
			MenuPrioritas(A)
	case 3:
			fmt.Print("On Progress")
	case 0:
			return
	}
}
func inputDataTugas(A *tabtugas) {
	var n int
	fmt.Println("[ INPUT DATA PRODUKTIVITAS ]")
	fmt.Println("BERAPA BANYAK DATA YANG INGIN ANDA INPUT")
	fmt.Scan(&n)
	fmt.Println("INPUT DATA ANDA SECARA BERURUTAN (NAMA, PRIORITAS, DEADLINE)")
	for idx < n+idx {
		A[idx].tanggal = tanggalBatch
		
		fmt.Scan(&A[idx].namaTugas, &A[idx].prioritas, &A[idx].deadline)
		idx ++
	}
}
func menuPrioritas(A *tabtugas){
	fmt.Println("==MENU PRIORITAS==")
	fmt.Println("[1] Ascending")
	fmt.Println("[2] Descending")
	fmt.Println("[0] Kembali")
	fmt.Scan(&pilihan)
    switch pilihan{
		case 1:
			tugasPrioritasAscending(A)
		case 2:
			tugasPrioritasDescending(A)
		case 0:
			return
	}
}
func tugasPrioritasDescend(A *tabtugas){
	var pass, j, indeks int
	var temp Tugas
	pass = 0
	for pass < idx{
		indeks = pass
		j = pass +1
		for j < idx{
			if A[j].prioritas>A[indeks].prioritas{
				indeks = j
			}
		}
		temp = A[pass]
		A[pass] = A[indeks]
		A[indeks] = temp
	}
}
func tugasPrioritasAscend(A *tabtugas){
	var pass, j, indeks int
	var temp Tugas
	pass = 0
	for pass < idx{
		indeks = pass
		j = pass +1
		for j < idx{
			if A[j].prioritas<A[indeks].prioritas{
				indeks = j
			}
		}
		temp = A[pass]
		A[pass] = A[indeks]
		A[indeks] = temp
		pass++
	}
}

func MenuKesehatanMental(){
	fmt.Printf("=== CEK KESEHATAN MENTAL ===\n")
	fmt.Println("APA YANG BISA KAMI BANTU HARI INI?")
	fmt.Println("[1] Kalkulasi Koin Mental dan Stress Meter")
	fmt.Println("[2] Isi Jurnal Suasana Hati")
	fmt.Println("[3] Lihat Stastistik Mingguan")
	fmt.Println("[0] Kembali ke Menu Utama")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
			hitungLimitMental(A)
	case 2:
			fmt.Print("On Progress")
	case 3:
			fmt.Print("On Progress")
	case 0:
			return
	}

	func hitungLimitMental(A *tabtugas, B *tabmental, tglAktif int){
		fmt.Println("[ CEK KOIN MENTAL DAN STRESS METER ]")

		bebanTotal := 0
		jumlahTugasYngAda := 0

		//Pembacaan/Perhitungan tugas dan koin mental

		for i := 0; i < idxTugas; i++ {
			if A[i].tanggal == tglAktif {
				jumlahTugasYngAda++
				bebanTambahan := A[i].durasi / 10
				if A[i].prioritas >= 4 {
					bebanTambahan += 5
				}
				bebanTotal += bebanTambahan
			}
		}
		sisaKoin := BatasKoin - bebanTotal
		stressMeter := 0

		if sisaKoin < 0 {
			stressMeter = (sisaKoin * -1) / 5
			sisaKoin = 0
			if stressMeter > BatasStress {
				stressMeter = LimitStress
			}
		}
		
	}

}
