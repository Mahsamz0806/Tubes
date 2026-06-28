package main
import "fmt"

//Variabel global dan konstanta
var pilihan, idxTugas, idxJurnal int
const NMAX int = 1000
const BatasKoinMental int = 30
const BatasStress int = 10

type Mental struct{
	tanggal int//Format DDMMYYYY
	skorEmosi int//Skala 1- 10
	catatanEmosi string
	sisaKoin int//Untuk menyimpan data dari Array Tugas
	stressMeter int//Untuk menimpan data dari Array Tugas
}
type tabmental [NMAX]Mental

type Tugas struct {
	tanggal int//Format DDMMYYYY
	namaTugas string
	prioritas int//Skala 1 sampai 5
	deadline int//Format 24 jam
}
type tabtugas [NMAX]Tugas

//=================================================================
func main() {
	var A tabtugas
	var B tabmental
	var tglAktif int
	tanggalAktif(&tglAktif)
	menuUtama(&A, &B, &tglAktif)
}

func tanggalAktif(tanggal *int) {//Tanggal untuk kedua menu, pada Produktivitas dan Mental
	fmt.Println("Masukkan Tanggal Sesi Ini (Format: DDMMYYYY, cth: 12062026): ")
	fmt.Scan(tanggal)
	fmt.Printf("Program beroperasi pada tanggal: %d\n", *tanggal)
}

func menuUtama(A *tabtugas, B *tabmental, tglAktif *int){//Persimpangan antara dua menu
	var keluar int
	keluar = 0
	for keluar == 0{
		fmt.Printf("===== ClearMind =====\n")
		fmt.Printf("[Tanggal Aktif Saat Ini: %d]\n", *tglAktif)//Untuk tanggal yang aktif pada kedua menu
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
				keluar = 1
		default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}

func StressMaksimal(A *tabtugas, tglAktif int)bool {
	var bebanTotal, i, sisaKoin, stressMeter int
	bebanTotal = 0
	for i = 0; i < idxTugas; i++ {
		if A[i].tanggal == tglAktif {
			bebanTotal += (A[i].prioritas * 3)
		}
	}
	sisaKoin = BatasKoinMental- bebanTotal
	if sisaKoin < 0 {
		stressMeter = (sisaKoin * -1) / 5
		if stressMeter >= BatasStress {
			return true
		}
	}
	return false
}

func produktifitas(A *tabtugas, B *tabmental, tglAktif *int) {
	var pilih, kembali int
	kembali = 0
	for kembali == 0{
		fmt.Println("===HALO SELAMAT DATANG DI PEMBANTU PRODUKTIFITAS===")
		fmt.Println("APA YANG KAMI BISA BANTU HARI INI?")
		fmt.Println("1. Input Tugas")
		fmt.Println("2. Tampilkan Tugas")
		fmt.Println("3. Tugas Prioritas")
		fmt.Println("4. Cari Tugas")
		fmt.Println("5. Ubah Tugas")
		fmt.Println("6. Hapus Tugas")
		fmt.Println("7. Kembali")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			if StressMaksimal(A, *tglAktif) {
				fmt.Println("\n===========================================")
				fmt.Println("!!! PERINGATAN: STRESS METER MAKSIMAL !!!")
				fmt.Println("BEBAN MENTAL ANDA TELAH PENUH")
				fmt.Println("Akses menambah tugas DITUTUP.")
				fmt.Println("=============================================")
				tambahJurnalOtomatis(B, *tglAktif)
			}else{
			inputDataProduk(A, *tglAktif)
			}
		case 2:
			tampilkanTugas(A, *tglAktif)
		case 3:
			menuPrioritas(A, B, tglAktif)
		case 4:
			MenucariTugas(A, B, tglAktif)
		case 5:
			ubahTugas(A, *tglAktif)
		case 6:
			hapusTugas (A, *tglAktif)
		case 7 :
			kembali = 1
		default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}

func inputDataProduk(A *tabtugas, tglAktif int) {//Untuk menangani data tugas baru ke array
	var n, batas int
	fmt.Println("==INPUT DATA PRODUKTIFITAS==")
	fmt.Println("BERAPA BANYAK DATA YANG INGIN ANDA INPUT")
	fmt.Scan(&n)
	batas = idxTugas + n
	if batas > NMAX{
		batas = NMAX
		fmt.Println("Data melebihi kapasitas, hanya sebagian yang disimpan")
	}
	fmt.Println("INPUT DATA ANDA SECARA BERURUTAN (NAMA, PRIORITAS, DEADLINE)")
	for idxTugas < batas {
		A[idxTugas].tanggal = tglAktif
		fmt.Scan(&A[idxTugas].namaTugas, &A[idxTugas].prioritas, &A[idxTugas].deadline)
		idxTugas++
	}
}

func tampilkanTugas(A *tabtugas, tglAktif int) {//Menampilkan daftar tugas yang ada pada tanggal aktif
	var i, no int
	no = 1
	fmt.Println("\n==DAFTAR TUGAS==")
	if idxTugas == 0 {
		fmt.Println("Belum Ada Tugas")
	} else {
		fmt.Printf("\n%-5s %-20s %-10s %-10s\n", "No", "Nama", "Prioritas", "Deadline")
		for i = 0; i < idxTugas; i++ {
			if A[i].tanggal == tglAktif {
				fmt.Printf("%-5d %-20s %-10d %-10d\n",
					no,
					A[i].namaTugas,
					A[i].prioritas,
					A[i].deadline,
				)
				no++
			}
		}
	}
}

func MenucariTugas(A *tabtugas, B *tabmental, tglAktif *int) {//Untuk mencari tugas-tugas yang sudah diinput
	var pilih, kembali int
	for kembali == 0{
		fmt.Println("\n==Cari Tugas==")
		fmt.Println("1. Berdasarkan Nama")
		fmt.Println("2. Berdasarkan Prioritas")
		fmt.Println("3. Kembali")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilih)
		switch pilih {
			case 1:
				cariSequential(A, *tglAktif)
			case 2:
				cariBinary(A, *tglAktif)
			case 3:
				kembali =1
			default:
				fmt.Println("Pilihan tidak valid")
		}
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
		if A[i].namaTugas == nama && A[i].tanggal == tglAktif {
			fmt.Printf("%s| %d | %d \n", A[i].namaTugas, A[i].prioritas, A[i].deadline)
			found = true
		}
		i = i + 1
	}
	if found == false{
		fmt.Println("Tugas tidak ditemukan")
	}
}

func cariBinary(A *tabtugas, tglAktif int) {
	var target, low, high, mid, hasil int
	tugasPrioritasAscend(A, tglAktif)
	fmt.Println("Prioritas: ")
	fmt.Scan(&target)
	low = 0
	high = idxTugas - 1
	hasil = -1
	for low <= high && hasil == -1 {
		mid = (low + high) / 2
		if A[mid].prioritas == target && A[mid].tanggal == tgl aktif {
			hasil = mid
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

func menuPrioritas(A *tabtugas, B *tabmental, tglAktif *int) {
	var pilih, kembali int
	kembali = 0
	for kembali == 0{
		fmt.Println("==MENU PRIORITAS==")
		fmt.Println("1.Ascending")
		fmt.Println("2.Descending")
		fmt.Println("3.Kembali")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilih)
		switch pilih {
			case 1:
				tugasPrioritasAscend(A, *tglAktif)
			case 2:
				tugasPrioritasDescend(A, *tglAktif)
			case 3:
				kembali = 1
			default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}

func tugasPrioritasDescend(A *tabtugas, tglAktif int) {
	var pass, j, indeks int
	var temp Tugas
	pass = 0
	for pass < idxTugas-1 {
		indeks = pass
		j = pass + 1
		for j < idxTugas {
			if A[j].prioritas > A[indeks].prioritas {
				indeks = j
			}
			j++
		}
		temp = A[pass]
		A[pass] = A[indeks]
		A[indeks] = temp
		pass++
	}
	tampilkanTugas(A, tglAktif)
}

func tugasPrioritasAscend(A *tabtugas, tglAktif int) {
	var pass, i int
	var temp Tugas
	pass = 1
	for pass < idxTugas {
			if A[pass].tanggal == tglAktif{
				temp = A[pass]
			i = pass - 1
			for i >= 0 && A[i].prioritas > temp.prioritas {
				A[i+1] = A[i]
				i = i - 1
			}
			A[i+1] = temp
		}
		pass = pass + 1
	}
	tampilkanTugas(A, tglAktif)
}
func ubahTugas(A *tabtugas, tglAktif int) {//untuk mengubah tugas
	var nomor, i, hitung, indeksAsli int
	fmt.Println("\n==UBAH TUGAS==")
	tampilkanTugas(A, tglAktif)
	if idxTugas == 0{
		fmt.Println("Tidak ada data tugas")
	}else{
		fmt.Println("Nomor tugas: ")
		fmt.Scan(&nomor)
		hitung = 0
		indeksAsli = -1
		for i = 0; i< idxTugas;i++{
			if A[i].tanggal == tglAktif{
				hitung++
				if hitung == nomor{
					indeksAsli = i
				}
			}
		}
		if indeksAsli == -1{
			fmt.Println("Nomor tidak valid")
		}else{
			fmt.Print("Masukan data baru(nama, prioritas, deadline): ")
			fmt.Scan(&A[indeksAsli].namaTugas,&A[indeksAsli].prioritas,&A[indeksAsli].deadline)
			fmt.Println("Tugas berhasil diubah")
		}
	}
}

func hapusTugas(A *tabtugas, tglAktif  int){// untuk menghapus tugas
	var nomor, i, hitung, indeksTarget int
	var namaTerhapus string
	fmt.Println("\n==HAPUS TUGAS==")
	tampilkanTugas(A, tglAktif)
	if idxTugas==0{
		fmt.Println("Tidak ada tugas yang bisa dihapus")
	}else{
		fmt.Print("Masukan nomor tugas yang ingin dihapus")
		fmt.Scan(&nomor)
		hitung = 0
		indeksTarget = -1
		for i = 0; i<idxTugas;i++{
			if A[i].tanggal == tglAktif{
				hitung++
				if hitung == nomor{
					indeksTarget = i
				}
			}
		}
		if indeksTarget == -1{
			fmt.Println("Nomor tidak valid")
		}else{
			namaTerhapus = A[indeksTarget].namaTugas
			for i = indeksTarget; i< idxTugas-1;i++{
				A[i] = A[i+1]
			}
			idxTugas--
			fmt.Printf("Tugas\"%s\" berhasil dihapus.\n", namaTerhapus)
		}
	}
}

func MenuKesehatanMental(A *tabtugas, B *tabmental, tglAktif *int){
	var kembali int
	kembali = 0
	for kembali == 0{
		fmt.Println("=== CEK KESEHATAN MENTAL ===\n")
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
				kembali = 1
		default: 
				fmt.Println("Pilihan tidak valid.")
		}
	}
}

	func hitungLimitMental(A *tabtugas, B *tabmental, tglAktif int){
		var bebanTotal, jumlahTugasYngAda, i, stressMeter, sisaKoin int
		var jurnalDitemukan bool
		fmt.Println("[ CEK KOIN MENTAL DAN STRESS METER ]")
		bebanTotal = 0
		jumlahTugasYngAda = 0

		//Pembacaan/Perhitungan tugas dan koin mental
		for i = 0; i < idxTugas; i++ {
			if A[i].tanggal == tglAktif {
				jumlahTugasYngAda++
				bebanTotal += (A[i].prioritas * 3)
			}
		}

		sisaKoin = BatasKoinMental - bebanTotal
		stressMeter = 0
		if sisaKoin < 0 {
			stressMeter = (sisaKoin * -1) / 5
			sisaKoin = 0
			if stressMeter > BatasStress {
				stressMeter = BatasStress
			}
		}

		//Menyambungkan data ke Array Mental
		jurnalDitemukan = false
		i = 0
		for i < idxJurnal && !jurnalDitemukan {
			if B[i].tanggal == tglAktif{
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
			B[idxJurnal].skorEmosi = 0
			B[idxJurnal].catatanEmosi = ""
			idxJurnal++
		}

		fmt.Printf("Total Tugas Hari Ini : %d Tugas\n", jumlahTugasYngAda)
		fmt.Printf("Beban Koin Mental    : %d\n", bebanTotal)
		fmt.Printf("Sisa Koin Anda       : %d / %d\n", sisaKoin, BatasKoinMental)
		fmt.Printf("Stress Meter         : %d / %d\n", stressMeter, BatasStress)
}

//Satu Tanggal hanya boleh satu jurnal
func tambahJurnalManual(B *tabmental, tglAktif int){
	var ada bool
	var i int
	fmt.Println("\n== TAMBAH JURNAL MENTAL ==")

	//Cek apakah tanggal sudah berjurnal
	ada = false
	for i = 0; i < idxJurnal; i++{
		if B[i].tanggal == tglAktif {
			ada = true
		}
	}
	
	if ada{
		fmt.Println("[!] Jurnal untuk tanggal ini sudah ada. Gunakan menu kalkulasi untuk memperbarui data.")
	}else if idxJurnal>= NMAX{
		fmt.Println("[!] Memori jurnal penuh")
	}else{	B[idxJurnal].tanggal = tglAktif
	fmt.Print("Masukkan Skor Emosi Anda Hari ini (1-10): ")
	fmt.Scan(&B[idxJurnal].skorEmosi)
	fmt.Println("Tuliskan deskripsi singkat perasaan Anda (Gunakan _ sebagai spasi): ")
	fmt.Scan(&B[idxJurnal].catatanEmosi)
	idxJurnal++
	fmt.Println("Jurnal baru telah disimpan")}
}

func tambahJurnalOtomatis(B *tabmental, tglAktif int){//Digunakan jika stress meter sudah melebihi maksimal
	var indeksTarget, j int
	var jurnalDitemukan bool
	indeksTarget = idxJurnal
	jurnalDitemukan = false
	j=0
	for j < idxJurnal && !jurnalDitemukan{
		if B[j].tanggal == tglAktif {
			indeksTarget = j
			jurnalDitemukan = true
		}
		j++
	}
	if jurnalDitemukan{
		B[indeksTarget].tanggal = tglAktif
		fmt.Print("Skor Emosi Akibat Burnout Hari Ini (1-10): ")
		fmt.Scan(&B[indeksTarget].skorEmosi)
		fmt.Println("Tuliskan apa yang Anda rasakan saat ini (Gunakan _ sebagai spasi): ")
		fmt.Scan(&B[indeksTarget].catatanEmosi)
		fmt.Println("Data kondisi darurat berhasil direkam. Beristirahatlah.")
	}else if !jurnalDitemukan && idxJurnal < NMAX{
		B[idxJurnal].tanggal = tglAktif
		fmt.Print("Skor Emosi Akibat Burn Out Hari Ini (1-10):")
		fmt.Scan(&B[idxJurnal].skorEmosi)
		fmt.Println("Tuliskan apa yang anda rasakan saat ini (gunakan '_' sebagai spasi):")
		fmt.Scan(&B[idxJurnal].catatanEmosi)
		idxJurnal++
		fmt.Println("Data kondisi darurat berhasil direkam. Beristirahatlah")
	}else{
		fmt.Println("Memori jurnal penuh")
	}

}

func tampilkanSemuaJurnal(B *tabmental) {//procedure untuk menampilkan semua jurnal
	var i int
	fmt.Println("\n== RIWAYAT JURNAL MENTAL ==")
	if idxJurnal > 0 {
		fmt.Printf("%-5s %-12s %-12s %-12s %-20s\n", "No", "Tanggal", "Skor Emosi", "Stress Meter", "Catatan")
		i = 0
		for i < idxJurnal {
			fmt.Printf("%-5d %-12d %-12d %-12d %-20s\n",
				i+1, B[i].tanggal, B[i].skorEmosi, B[i].stressMeter, B[i].catatanEmosi)
			i = i + 1
		}
	} else {
		fmt.Println("Belum ada riwayat jurnal.")
	}
}

func menuCariJurnal(B *tabmental) {//procedure untuk mencari jurnal berdasarkan berbagai kategori
	var target, kembali, i, hasil int
	var found bool
	kembali = 0
	for kembali == 0{
		fmt.Println("\n== CARI JURNAL MENTAL ==")
		fmt.Println("[1] Cari Berdasar Skor Emosi")
		fmt.Println("[2] Cari Berdasar Tanggal")
		fmt.Println("[0] Kembali ke Menu Kesehatan Mental")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
	
		switch pilihan {
		case 1:
			fmt.Print("Masukkan Skor Emosi yang dicari (0-10): ")
			fmt.Scan(&target)
			found = false
			i = 0
			for i < idxJurnal {
				if B[i].skorEmosi == target {
					fmt.Printf("Ditemukan -> Tgl: %d | Stres: %d | Catatan: %s\n", B[i].tanggal, B[i].stressMeter, B[i].catatanEmosi)
					found = true
				}
				i = i + 1
			}
			if found == false {
				fmt.Println("Tidak ada jurnal dengan skor emosi tersebut.")
			}
		case 2:
			fmt.Print("Masukkan Tanggal yang dicari (Pastikan tanggal sudah diurutkan): ")
			fmt.Scan(&target)
			hasil = cariJurnalTanggalRekursif(B, target, 0)
			if hasil == -1{
				fmt.Println("Jurnal tidak ditemukan")
			}else{
				fmt.Printf("Ditemukan -> Tgl: %d | Emosi: %d | Stress: %d | Catatan: %s\n",
				B[hasil].tanggal, B[hasil].skorEmosi, B[hasil].stressMeter, B[hasil].catatanEmosi)
			}
		case 0:
			kembali = 1
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

//Sorting Bagian Mental secara Selection dan Insertion tergantung pilihan
func menuSortJurnal(B *tabmental){
	var kembali, pass, j, i, indeks int
	var temp Mental
	kembali = 0
	for kembali == 0{
		fmt.Println("\n== URUTKAN JURNAL MENTAL ==")
		fmt.Println("[1] Urutkan Berdasar Skor Emosi Tertinggi")
		fmt.Println("[2] Urutkan Berdasar Tanggal Terlama")
		fmt.Println("[0] Kembali ke Menu Kesehatan Mental")
		fmt.Println("Pilih: ")
		fmt.Scan(&pilihan)
	
		switch pilihan {
			case 1: //Pengurutan Selection secara Descending
					for pass = 0; pass < idxJurnal-1; pass++ {
						indeks = pass
						for j = pass + 1; j < idxJurnal; j++ {
							if B[j].skorEmosi > B[indeks].skorEmosi {
								indeks = j
							}
						}
						temp = B[pass]
						B[pass] = B[indeks]
						B[indeks] = temp
					}
			fmt.Println("Riwayat jurnal telah diurutkan berdasar skor emosi tertinggi.")
			tampilkanSemuaJurnal(B)
			case 2:
					for pass = 1; pass < idxJurnal; pass++ {
						temp = B[pass]
						i = pass - 1
						for i >= 0 && B[i].tanggal > temp.tanggal {
							B[i+1] = B[i]
							i = i -1
						}
						B[i+1] = temp
					}
					fmt.Println("Riwayat jurnal diurutkan berdasar tanggal terlama.")
					tampilkanSemuaJurnal(B)
			case 0:
				kembali = 1
			default:
				fmt.Println("Pilihan tidak valid")
		}
	}
}

func hapusJurnal(B *tabmental, tglAktif int) {//procedure untuk menghapus jurnal
	var nomor, i, indeksTarget int
	var tglTerhapus int
	fmt.Println("\n== HAPUS JURNAL ==")
	tampilkanSemuaJurnal(B)
	if idxJurnal > 0 {
		fmt.Print("\nMasukkan Nomor Jurnal yang ingin dihapus: ")
		fmt.Scan(&nomor)
		if nomor >= 1 && nomor <= idxJurnal {
			indeksTarget = nomor - 1
			tglTerhapus = B[indeksTarget].tanggal
			i = indeksTarget
			for i < idxJurnal-1 {
				B[i] = B[i+1]
				i = i + 1
			}
			idxJurnal--
			fmt.Printf("Jurnal untuk tanggal %d berhasil dihapus.\n", tglTerhapus)
		} else {
			fmt.Println("Nomor tidak valid. Penghapusan dibatalkan.")
		}
	} else {
		fmt.Println("Tidak ada jurnal yang bisa dihapus.")
	}
}

func cariJurnalTanggalRekursif(B *tabmental, tglTarget int, idx int)int{
	if idx >= idxJurnal{
		return -1
	}
	if B[idx].tanggal == tglTarget{
		return idx
	}
	return cariJurnalTanggalRekursif(B, tglTarget, idx+1)
}
