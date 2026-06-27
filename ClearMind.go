package main
import "fmt"

//Variabel global untuk diapaki semua function dan procedure
var pilihan, idxTugas, idxJurnal int
const NMAX int = 1000
const BatasKoinMental int = 30
const BatasStress int = 10//Untuk sementara segini

type Mental struct{
	tanggal int//Format DDMMYYYY
	skorEmosi int//Skala 1- 10
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
func MenuKesehatanMental(A *tabtugas, B *tabmental, tglAktif *int){
	fmt.Printf("=== CEK KESEHATAN MENTAL ===\n", *tglAktif)
	fmt.Println("APA YANG BISA KAMI BANTU?")
	fmt.Println("[1] Kalkulasi Koin Mental Harian")
	fmt.Println("[2] Tambah Jurnal Baru")
	fmt.Println("[3] Tampilkan Semua Riwayat Jurnal")
	fmt.Println("[4] Cari Riwayat Jurnal")
	fmt.Println("[5] Urutkan Riwayat Jurnal")
	fmt.Println("[6] Hapus Jurnal Berdasar Tanggal")
	fmt.Println("[0] Kembali ke Menu Utama")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
			hitungLimitMental(A, B, *tglAktif)
	case 2:
			tambahJurnalManual(B, *tglAktif)
	case 3:
			tampilkanSemuaJurnal(B)
	case 4:
			menuCariJurnal(B)
	case 5:
			menuSortJurnal(B)
	case 6:
			hapusJurnal(B, *tglAktif)
	case 0:
			menuUtama(A)
	default: 
			fmt.Println("Pilihan tidak valid.")
	}
}

	func hitungLimitMental(A *tabtugas, B *tabmental, tglAktif int){
		fmt.Println("[ CEK KOIN MENTAL DAN STRESS METER ]")
		bebanTotal := 0
		jumlahTugasYngAda := 0

		//Pembacaan/Perhitungan tugas dan koin mental
		for i := 0; i < idxTugas; i++ {
			if A[i].tanggal == tglAktif {
				jumlahTugasYngAda++
				bebanTotal += (A[i].prioritas * 3)
			}
		}

		sisaKoin := BatasKoinHarian - bebanTotal
		stressMeter := 0
		if sisaKoin < 0 {
			stressMeter = (sisaKoin * -1) / 5
			sisaKoin = 0
			if stressMeter > BatasStress {
				stressMeter = BatasStress
			}
		}

		//Menyambungkan data ke Array Mental
		jurnalDitemukan : false
		i := 0
		for i < idxJurnal && !jurnalDitemukan {
			if B[i].tanggal == tglAktif
				B[i].sisaKoin = sisaKoin
				B[i].stressMeter = stressMeter
				jurnalDitemukan = true
			}
			i++
		}

		if !jurnalDitemukan && idxJurnal < NMAX {
			B[idxJurnal].tanggal = tglAktif
			B[idxJurnal].sisaKoin = sisaKoin
			B[idxJurnal].stressMeter = stressMeter
			idxJurnal++
		}

		fmt.Printf("Total Tugas Hari Ini : %d Tugas\n", jumlahTugasDitemukan)
		fmt.Printf("Beban Koin Mental    : %d\n", bebanTotal)
		fmt.Printf("Sisa Koin Anda       : %d / %d\n", sisaKoin, BatasKoinHarian)
		fmt.Printf("Stress Meter         : %d / %d\n", stressMeter, BatasStress)
}

//Satu Tanggal hanya boleh satu jurnal
func tambahJurnalManual(B *tabJurnal, tglAktif int){
	fmt.Println("\n== TAMBAH JURNAL MENTAL ==")

	//Cek apakah tanggal sudah berjurnal
	ada := false
	for I := 0; i < idxJurnal; i++{
		if B[i].tanggal == tglAktif {
			ada = true
		}
	}
	
	if ada{
		fmt.Println("[!] Jurnal untuk tanggal ini sudah ada. Gunakan menu kalkulasi untuk memperbarui data.")
		return
	}
	if idxJurnal >= NMAX {
		fmt.Println("[!] Memori jurnal penuh.")	
		return
	}

	B[idxJurnal].tanggal = tglAktif
	fmt.Print("Masukkan Skor Emosi Anda Hari ini (1-10): ")
	fmt.Scan(&B[idxJurnal].skorEmosi)
	fmt.Println("Tuliskan deskripsi singkat perasaan Anda (Gunakan _ sebagai spasi): ")
	fmt.Scan(&B[idxJurnal].catatanEmosi)
	idxJurnal++
	fmt.Println("Jurnal baru telah disimpan")
}

func tampilkanSemuaJurnal(B *tabmental){
	fmt.Println("\n== RIWAYAT JURNAL MENTAL ==")
	if idxJurnal == 0 {
		fmt.Println("Belum ada riwayat jurnal.")
		return
	}
	fmt.Printf("%-5s %-12s %-12s %-12s %-20s\n", "No", "Tanggal", "Skor Emosi", "Stress Meter", "Catatan")
	for i := 0; i < idxJurnal; i++ {
		fmt.Printf("%-5d %-12d %-12d %-12d %-20s\n", 
			i+1, B[i].tanggal, B[i].skorEmosi, B[i].stressMeter, B[i].catatanEmosi)
	}
}

func menuCariJurnal(B *tabmental){
	var target int
	fmt.Println("\n== CARI JURNAL MENTAL ==")
	fmt.Println("[1] Cari Berdasar Skor Emosi")
	fmt.Println("[2] Cari Berdasar Tanggal")
	fmt.Println("[0] Kembali ke Menu Kesehatan Mental")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1: //Pencarian sequential berdasar skor emosi
			fmt.Print("Masukkan Skor Emosi yang dicari (0-10): ")
			fmt.Scan(&target)
			found := false
			for i := 0; i < idxJurnal; i++ {
				if B[i].skorEmosi == target {
					fmt.Printf("Ditemukan -> Tgl: %d | Stres: %d | Catatan: %s\n", B[i].tanggal, B[i].stressMeter, B[i].catatanEmosi)
				found = true
				}
			}
			if !found {
				fmt.Println("Tidak ada jurnal dengan skor emosi tersebut.")
			}
	case 2: //Pencarian binary berdasar tanggal
			fmt.Print("Masukkan Tanggal yang dicari (Pastikan tanggal sudah diurutkan): ")
			fmt.Scan(&target)
			low := 0
			high := idxJurnal - 1
			hasil := -1
			for low <= high && hasil == -1 {
				mid := (low + high) / 2
				if B[mid].tanggal == target {
					hasil = mid
				}else if B[mid].tanggal < target {
					low = mid + 1
				}else{
					high = mid - 1
				}
			}
			if hasil == -1 {
				fmt.Println("Jurnal tidak ditemukan pada tanggal tersebut.")
			}else{
				fmt.Printf("Ditemukan -> Tgl: %d | Emosi: %d | Stres: %d | Catatan: %s\n", 
				B[hasil].tanggal, B[hasil].skorEmosi, B[hasil].stressMeter, B[hasil].catatanEmosi)
			}
	case 0:
			MenuKesehatanMental(B)
}

//Sorting Bagian Mental secara Selection dan Insertion tergantung pilihan
func menuSortJurnal(B *tabmental){
	var target int
	fmt.Println("\n== URUTKAN JURNAL MENTAL ==")
	fmt.Println("[1] Urutkan Berdasar Skor Emosi Tertinggi")
	fmt.Println("[2] Urutkan Berdasar Tanggal Terlama")
	fmt.Println("[0] Kembali ke Menu Kesehatan Mental")
	fmt.Println("Pilih: ")
	fmt.Scan(&pilihan)

	switch pilihan {
		case 1: //Pengurutan Selection secara Descending
				for pass := 0; pass < idxJurnal-1; pass++ {
					indeks := pass
					for j := pass + 1; j < idxJurnal; j++ {
						if B[j].skorEmosi > B[indeks].skorEmosi {
							indeks = j
						}
					}
					temp := B[pass]
					B[pass] = B[indeks]
					B[indeks] = temp
				}
		fmt.Println("Riwayat jurnal telah diurutkan berdasar skor emosi tertinggi.")
		tampilkanSemuaJurnal(B)
		case 2:
				for pass := 1; pass < idxJurnal; pass++ {
					temp := B[pass]
					i := pass - 1
					for I >= 0 && B[i].tanggal > temp.tanggal {
						B[i+1] = B[i]
						i = i -1
					}
					B[i+1] = temp
				}
				fmt.Println("Riwayat jurnal diurutkan berdasar tanggal terlama.")
				tampilkanSemuaJurnal(B)
	}
}

func hapusJurnal(B *tabmental, tglAktif int){
	fmt.Println("\n== HAPUS JURNAL")
	if idxJurnal == 0 {
		fmt.Printf("Tidak ada data yang jurnal yang bisa dihapus.")
		return
	}
	fmt.Printf("%-5s %-12s %-12s %-20s\n", "No", "Tanggal", "Skor Emosi", "Catatan")
	for i := 0; i < idxJurnal; i++ {
		fmt.Printf("%-5d %-12d %-12d %-20s\n", i+1, B[i].tanggal, B[i].skorEmosi, B[i].catatanEmosi)
	}
	var nomor int
	fmt.Print("\nMasukkan Nomor Jurnal yang ingin dihapus: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > idxJurnal {
		fmt.Println("Nomor tidak valid. Penghapusan dibatalkan.")
		return
	}

	indeksTarget := nomor - 1
	tglTerhapus := B[indeksTarget].tanggal

	for i := indeksTarget; i < idxJurnal - 1; i++ {
		B[i] = B[i+1]
	}
	idxJurnal--
	fmt.Printf("Jurnal untuk tanggal %d berhasil dihapus dari sistem>\n", tglTerhapus)
}
