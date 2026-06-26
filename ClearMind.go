package main
import "fmt"

//Variabel global untuk diapaki semua function dan procedure
var pilihan, idxTugas, idxJurnal int
const NMAX int = 1000
const BatasKoinMental int = 30
const BatasStress int = 10//Untuk sementara segini

type Mental struct{
	tanggal int//Format DDMMYYYY
	skorEmosi int
	catatanEmosi string
	sisaKoin int//Untuk menyimpan data dari Array Tugas
	stressMeter int//Untuk menimpan data dari Array Tugas
}
type tabmental [NMAX]JurnalMental

type Tugas struct {
	tanggal int//Format DDMMYYYY
	namaTugas string
	prioritas int//Skala 1 sampai 5
	deadline  int//Format 24 jam
}
type tabtugas [NMAX]Tugas
//=================================================================
func main() {
	var A tabtugas
	var B tabmental
	var tglAktif int
	setTanggalSesi(&tanggalAktif)
	menuUtama(&A, &B, &tanggalAktif)
}

func tanggalAktif(tanggal *int) {//Tanggal untuk kedua menu, pada Produktivitas dan Mental
	fmt.Println("Masukkan Tanggal Sesi Ini (Format: DDMMYYYY, cth: 12062026): ")
	fmt.Scan(&tanggal)
	fmt.Printf("Program beroperasi pada tanggal: %d\n", *tanggal)
}

func menuUtama(A *tabtugas, B *tabJurnal, tglAktif *int){
	fmt.Printf("===== ClearMind =====\n")
	fmt.Print("[Tanggal Aktif Saat Ini: %d]\n", *tglAktif)//Untuk tanggal yang aktif pada kedua menu
	fmt.Println("[1] Menu Produktivitas")//Untuk masuk ke menu Produktivitas
	fmt.Println("[2] Menu Kesehatan Mental")//Untuk masuk ke menu Kesehatan Mental
	fmt.Println("[3] Ubah Tanggal")
	fmt.Println("[0] Keluar Progam")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
			produktifitas(A, B, tglAktif)
	case 2:
			MenuKesehatanMental(A, B, tglAktif)
	case 3:
			tanggalAktif(tglAktif)
	case 0:
			fmt.Print("Terima Kasih sudah menggunakan ClearMind. Semoga hari Anda menyenangkan.")
	default:
			fmt.Println("Pilihan tidak valid")
	}
	menuUtama(A, B, tglAktif)
}
func produktifitas(A *tabtugas, tglAktif *int) {
	var pilih int
	fmt.Println("===HALO SELAMAT DATANG DI PEMBANTU PRODUKTIFITAS===")
	fmt.Println("APA YANG KAMI BISA BANTU HARI INI?")
	fmt.Println("1. Input Tugas")
	fmt.Println("2. Tampilkan Tugas")
	fmt.Println("3. Tugas Prioritas")
	fmt.Println("4. Cari Tugas")
	fmt.Println("5. Ubah Tugas")
	fmt.Println("6. Kembali")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		inputDataProduk(A, *tglAktif)
	case 2:
		tampilkanTugas(A, *tglAktif)
	case 3:
		menuPrioritas(A)
	case 4:
		MenucariTugas(A)
	case 5:
		ubahTugas(A, *tglAktif)
	case 6:
		menuUtama(A)
	}
}
func inputDataProduk(A *tabtugas, tglAktif int) {
	var n, batas int
	fmt.Println("==INPUT DATA PRODUKTIFITAS==")
	fmt.Println("BERAPA BANYAK DATA YANG INGIN ANDA INPUT")
	fmt.Scan(&n)
	batas = idxTugas + n
	fmt.Println("INPUT DATA ANDA SECARA BERURUTAN (NAMA, PRIORITAS, DEADLINE)")
	for idxTugas < batas {
		A[idxTugas].tanggal = tglAktif
		fmt.Scan(&A[idxTugas].namaTugas, &A[idxTugas].prioritas, &A[idxTugas].deadline)
		idxTugas++
	}
}
func tampilkanTugas(A *tabtugas, tglAktif int) {
	var i int
	fmt.Println("\n==DAFTAR TUGAS==")
	if idxTugas == 0 {
		fmt.Println("Belum Ada Tugas")
	} else {
		fmt.Printf("%-5s %-20s %-10s %-10s", "No", "Nama", "Prioritas", "Deadline")
		for i = 0; i < idxTugas; i++ {
			if A[i].tanggal == tglAktif {
				fmt.Printf("%-5d %-20s %-10d %-10d\n",
					i+1,
					A[i].namaTugas,
					A[i].prioritas,
					A[i].deadline,
				)
			}
		}
	}
}

func MenucariTugas(A *tabtugas, tglAktif int) {
	var pilih int
	fmt.Println("\n==Cari Tugas==")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Prioritas")
	fmt.Println("3. Kembali")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		cariSequential(A)
	case 2:
		cariBinary(A)
	case 3:
		menuPrioritas(A)

	}
}

func cariSequential(A *tabtugas, tglAktif int) {
	var nama string
	var found bool
	var i int
	fmt.Println("Nama Tugas: ")
	fmt.Scan(&nama)
	i = 0
	found = false
	for i < idxTugas && found == false {
		if A[i].namaTugas == nama {
			fmt.Printf("%s| %d | %d \n", A[i].namaTugas, A[i].prioritas, A[i].deadline)
			found = true
		}
		i = i + 1
	}
}

func cariBinary(A *tabtugas, tglAktif int) {
	var target, low, high, mid, hasil int
	fmt.Println("Prioritas: ")
	fmt.Scan(&target)
	low = 0
	high = idxTugas - 1
	hasil = -1
	for low <= high && hasil == -1 {
		mid = (low + high) / 2
		if A[mid].prioritas == target {
			hasil = mid
			low = high + 1
		} else if A[mid].prioritas < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if hasil == -1 {
		fmt.Println("Tidak Ditemukan")
	} else {
		fmt.Printf("%s|%d|%d\n", A[hasil].namaTugas, A[hasil].prioritas, A[hasil].deadline)
	}

}
func menuPrioritas(A *tabtugas, tglAktif int) {
	var pilih int
	fmt.Println("==MENU PRIORITAS==")
	fmt.Println("1.Ascending")
	fmt.Println("2.Descending")
	fmt.Println("3.Kembali")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		tugasPrioritasAscend(A)
	case 2:
		tugasPrioritasDescend(A)
	case 3:
		produktifitas(A)
	}
}
func tugasPrioritasDescend(A *tabtugas) {
	var pass, j, indeks int
	var temp Tugas
	pass = 0
	for pass < idxTugas {
		indeks = pass
		j = pass + 1
		for j < idxTugas {
			if A[j].prioritas > A[indeks].prioritas {
				indeks = j
			}
		}
		temp = A[pass]
		A[pass] = A[indeks]
		A[indeks] = temp
	}
}
func tugasPrioritasAscend(A *tabtugas, tglAktif int) {
	var pass, i int
	var temp Tugas
	pass = 1
	for pass < idxTugas {
		temp = A[pass]
		i = pass - 1
		for i >= 0 && A[i].prioritas > temp.prioritas {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = temp
		pass = pass + 1
	}

}
func ubahTugas(A *tabtugas, tglAktif int) {
	var nomor, i int
	fmt.Println("\n==UBAH TUGAS==")
	tampilkanTugas(A)
	if idxTugas == 0 {
		fmt.Println("Tidak ada data")
	} else {
		fmt.Print("Nomor Tugas: ")
		fmt.Scan(&nomor)
		if nomor < 1 || nomor > idxTugas {
			fmt.Println("Nomor tidak valid")
		} else {
			i = nomor - 1
			fmt.Print("Masukan data baru (nama, prioritas, deadline): ")
			fmt.Scan(&A[i].namaTugas, &A[i].prioritas, &A[i].deadline)
			fmt.Println("Tugas diubah")
		}
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
		jurnalDitemukan := false
		for i := 0; i < idxJurnal; i++ {
			if B[i].tanggal == tglAktif {
				B[i].sisaKoin = sisaKoin
				B[i].stressMeter = stressMeter
				jurnalDitemukan = true
			}
		}
		if !jurnalDitemukan && idxJurnal < NMAX {
			B[idxJurnal].tanggal = tglAktif
			B[idxJurnal].sisaKoin = sisaKoin
			B[idxJurnal].stressMeter = stressMeter
			idxJurnal++
		}
		fmt.Printf("Total Tugas Hari Ini      : %d Tugas\n", JumlahTugasYngAda)
		fmt.Printf("Koin Mental yang Terpakai : %d\n", bebanTotal)
		fmt.Printf("Sisa Koin Mental          : %d / %d\n", sisaKoin, BatasKoinHarian)
		fmt.Printf("Stress Meter              : %d / %d\n", stressMeter, BatasStress)
	}

}
