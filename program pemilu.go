/// Tubes Alpro 2 S1 IF 11 05
/// Program Pemilu

// Anggota Kelompok :
/// Caroline Carren Aureliya R. - 2311102174
/// Moch. Aditya Sulistiawan - 2311102193
/// Yesika Widiyani - 2311102195
/// Andreas Besar Wibowo - 2311102198

package main

import (
	"fmt"
)

const NMAX = 1000

// Struct date digunakan untuk menyimpan tanggal mulai dan selesai pemilu
type date struct {
	dd1, mm1, yy1, dd2, mm2, yy2 int
}

// Struct tabDate adalah array dari struct date dengan ukuran NMAX
type tabDate [NMAX]date

// Struct provinsi digunakan untuk menyimpan data calon, pemilih, dan total suara
type provinsi struct {
	calon            calon
	pemilih          pemilih
	pilih            int
	totalSuaraCalon  int
	totalSuaraPartai int
}

// Struct calon digunakan untuk menyimpan nama dan partai calon
type calon struct {
	nama, partai string
}

// Struct pemilih digunakan untuk menyimpan nama pemilih dan pilihan mereka
type pemilih struct {
	nama    string
	pilihan int
}

// Array untuk menyimpan data provinsi dengan ukuran NMAX
type arrProvA [NMAX]provinsi
type arrProvB [NMAX]provinsi
type arrProvC [NMAX]provinsi

var (
	mainArrayA arrProvA
	mainArrayB arrProvB
	mainArrayC arrProvC
	threshold  int // Dibutuhkan untuk validasi di berbagai fungsi
)

func main() {
	threshold = 100 // Set nilai default ambang batas suara
	var A date
	var TA arrProvA
	var TB arrProvB
	var TC arrProvC
	var nA, nB, nC int
	var nPA, nPB, nPC int

	menuUtama(&A, &TA, &TB, &TC, &nA, &nB, &nC, &nPA, &nPB, &nPC)
}

// Menampilkan menu utama aplikasi
func menuUtama(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int
	fmt.Println()
	fmt.Println("=======================================")
	fmt.Println("    PEMILU RAYA INDONESIA TAHUN 2025   ")
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Println("Masuk sebagai : ")
	fmt.Println("1. KPU")
	fmt.Println("2. Pemilih")
	fmt.Println("3. Selesai")

	fmt.Println()
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 2 {
		menuPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else if pilihan == 3 {
		hasilPEMILU(*TA, *TB, *TC, *nA, *nB, *nC, *nPA, *nPB, *nPC)
		fmt.Println()
		fmt.Println("----------------------------------")
		fmt.Println("Terima Kasih")
	} else {
		fmt.Println("Pilihan salah")
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

// Menampilkan menu untuk pemilih
func menuPemilih(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var tanggal, bulan, tahun int

	fmt.Println()
	fmt.Println("==========================")
	fmt.Println("Anda Masuk Sebagai Pemilih")
	fmt.Println("==========================")
	fmt.Println()

	fmt.Println("Tanggal saat ini : (dd/m/yyyy)")

	fmt.Scan(&tanggal, &bulan, &tahun)

	if cekTanggal(*A, tanggal, bulan) == false {
		fmt.Println("----------------------------------------------")
		fmt.Println("                  * Maaf *                   ")
		fmt.Println("    tidak ada PEMILU pada tanggal tersebut   ")
		fmt.Println("----------------------------------------------")
		fmt.Println()
		cetakCalonA(*TA, *nA)
		cetakCalonB(*TB, *nB)
		cetakCalonC(*TC, *nC)
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	} else {
		menuMasukPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

// Mengisi data pemilih untuk Provinsi A
func isiPemilihA(TA *arrProvA, nPA *int) {
	var nama string
	var pilihan int

	fmt.Println("\nMasukkan data pemilih Provinsi A:")
	fmt.Print("Nama pemilih (atau STOP untuk selesai): ")
	fmt.Scan(&nama)

	for nama != "STOP" {
		// Validasi nama
		if len(nama) < 3 || len(nama) > 50 {
			fmt.Println("Nama tidak valid! (3-50 karakter)")
		} else {
			// Cek duplikasi
			isDuplikat := false
			for i := 0; i < *nPA; i++ {
				if TA[i].pemilih.nama == nama {
					isDuplikat = true
					break
				}
			}

			if isDuplikat {
				fmt.Println("Pemilih sudah terdaftar!")
			} else {
				fmt.Print("Pilihan (nomor urut calon): ")
				fmt.Scan(&pilihan)

				// Simpan data pemilih
				TA[*nPA].pemilih.nama = nama
				TA[*nPA].pemilih.pilihan = pilihan
				// Tambahkan suara ke calon yang dipilih
				TA[pilihan-1].totalSuaraCalon++ // Tambahkan ini
				*nPA++
				fmt.Println("Data pemilih berhasil ditambahkan!")
			}
		}

		fmt.Print("\nNama pemilih berikutnya (atau STOP untuk selesai): ")
		fmt.Scan(&nama)
	}
}

// Mengisi data pemilih untuk Provinsi B
func isiPemilihB(TB *arrProvB, nPB *int) {
	var nama string
	var pilihan int

	fmt.Println("\nMasukkan data pemilih Provinsi B:")
	fmt.Print("Nama pemilih (atau STOP untuk selesai): ")
	fmt.Scan(&nama)

	for nama != "STOP" {
		// Validasi nama
		if len(nama) < 3 || len(nama) > 50 {
			fmt.Println("Nama tidak valid! (3-50 karakter)")
		} else {
			// Cek duplikasi
			isDuplikat := false
			for i := 0; i < *nPB; i++ {
				if TB[i].pemilih.nama == nama {
					isDuplikat = true
					break
				}
			}

			if isDuplikat {
				fmt.Println("Pemilih sudah terdaftar!")
			} else {
				fmt.Print("Pilihan (nomor urut calon): ")
				fmt.Scan(&pilihan)

				// Simpan data pemilih
				TB[*nPB].pemilih.nama = nama
				TB[*nPB].pemilih.pilihan = pilihan
				// Tambahkan suara ke calon yang dipilih
				TB[pilihan-1].totalSuaraCalon++ // Tambahkan ini
				*nPB++
				fmt.Println("Data pemilih berhasil ditambahkan!")
			}
		}

		fmt.Print("\nNama pemilih berikutnya (atau STOP untuk selesai): ")
		fmt.Scan(&nama)
	}
}

// Mengisi data pemilih untuk Provinsi C
func isiPemilihC(TC *arrProvC, nPC *int) {
	var nama string
	var pilihan int

	fmt.Println("\nMasukkan data pemilih Provinsi C:")
	fmt.Print("Nama pemilih (atau STOP untuk selesai): ")
	fmt.Scan(&nama)

	for nama != "STOP" {
		// Validasi nama
		if len(nama) < 3 || len(nama) > 50 {
			fmt.Println("Nama tidak valid! (3-50 karakter)")
		} else {
			// Cek duplikasi
			isDuplikat := false
			for i := 0; i < *nPC; i++ {
				if TC[i].pemilih.nama == nama {
					isDuplikat = true
					break
				}
			}

			if isDuplikat {
				fmt.Println("Pemilih sudah terdaftar!")
			} else {
				fmt.Print("Pilihan (nomor urut calon): ")
				fmt.Scan(&pilihan)

				// Simpan data pemilih
				TC[*nPC].pemilih.nama = nama
				TC[*nPC].pemilih.pilihan = pilihan
				// Tambahkan suara ke calon yang dipilih
				TC[pilihan-1].totalSuaraCalon++ // Tambahkan ini
				*nPC++
				fmt.Println("Data pemilih berhasil ditambahkan!")
			}
		}

		fmt.Print("\nNama pemilih berikutnya (atau STOP untuk selesai): ")
		fmt.Scan(&nama)
	}
}

// Menampilkan menu untuk pemilih yang sudah masuk
func menuMasukPemilih(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int
	var selesai bool = false

	for !selesai {
		fmt.Println("\nPilih Provinsi Anda:")
		fmt.Println("1. Provinsi A")
		fmt.Println("2. Provinsi B")
		fmt.Println("3. Provinsi C")
		fmt.Println("4. Kembali")
		fmt.Println()

		fmt.Print("Pilihan anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 4 {
			menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			selesai = true
		} else {
			handlePilihanProvinsi(pilihan, A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}
	}
}

// Fungsi helper untuk menangani pilihan provinsi
func handlePilihanProvinsi(pilihan int, A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	if pilihan >= 1 && pilihan <= 3 {
		var pilihanMenu int
		fmt.Println("1. Cari data calon")
		fmt.Println("2. Mulai memilih!")
		fmt.Println()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihanMenu)

		if pilihanMenu == 1 {
			handlePencarianCalon(pilihan, A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		} else if pilihanMenu == 2 {
			handlePemilihan(pilihan, A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}
	}
}

// Fungsi untuk validasi pemilih
func validasiPemilih(nama string, pilihan int, maxPilihan int) bool {
	if len(nama) < 3 || len(nama) > 50 {
		fmt.Println("Nama pemilih harus antara 3-50 karakter!")
		return false
	}

	if pilihan < 1 || pilihan > maxPilihan {
		fmt.Printf("Pilihan harus antara 1-%d!\n", maxPilihan)
		return false
	}

	return true
}

// Fungsi untuk mengecek duplikasi pemilih
func cekDuplikasiPemilih(nama string, TA arrProvA, TB arrProvB, TC arrProvC, nPA, nPB, nPC int) bool {
	isDuplikat := false

	// Cek di Provinsi A
	i := 0
	for i < nPA && !isDuplikat {
		if TA[i].pemilih.nama == nama {
			isDuplikat = true
		}
		i++
	}

	// Cek di Provinsi B jika belum ditemukan
	i = 0
	for i < nPB && !isDuplikat {
		if TB[i].pemilih.nama == nama {
			isDuplikat = true
		}
		i++
	}

	// Cek di Provinsi C jika belum ditemukan
	i = 0
	for i < nPC && !isDuplikat {
		if TC[i].pemilih.nama == nama {
			isDuplikat = true
		}
		i++
	}

	return isDuplikat
}

// Menampilkan menu data calon
func menuDataCalon(TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC *int) {
	fmt.Println("==================")
	fmt.Println("  * DATA CALON *  ")
	fmt.Println("==================")

	//tampilkan data calon
	cetakCalonA(*TA, *nA)
	cetakCalonB(*TB, *nB)
	cetakCalonC(*TC, *nC)

	fmt.Println()
	fmt.Println("Daftar Menu : ")
	fmt.Println("1. Tambah Data Calon")
	fmt.Println("2. Cari Data Calon")
	fmt.Println("3. Kembali")

	fmt.Println()
}

// Menampilkan menu untuk mencari pemilih
func menuCariPemilih() {
	fmt.Println("-----------------")
	fmt.Println("Pilih Provinsi : ")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Println("4. Kembali")

	fmt.Println()
}

// Menampilkan menu untuk mencari calon
func menuCariCalon() {
	fmt.Println("-----------------")
	fmt.Println("Pilih Provinsi : ")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Println("4. Kembali")

	fmt.Println()
}

// Menampilkan dan mengelola data calon
func dataCALON(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int
	var nama string

	menuDataCalon(TA, TB, TC, nA, nB, nC)

	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1: // tambah data calon
		isiCalon(TA, TB, TC, nA, nB, nC)
		dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

	case 2: // cari data calon
		menuCariCalon()
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1: // Provinsi A
			cetakCalonA(*TA, *nA)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)

			if pilihan == 1 { // edit data calon
				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di edit : ")
				fmt.Scan(&nama)
				fmt.Println("---------------------------------------")

				found := false
				for i := 0; i < *nA; i++ {
					if TA[i].calon.nama == nama {
						editCalonA(TA, i+1, *nA)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 { // hapus data calon
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di hapus : ")
				fmt.Scan(&nama)
				fmt.Println("----------------------------------------")

				found := false
				for i := 0; i < *nA; i++ {
					if TA[i].calon.nama == nama {
						hapusCalonA(TA, i+1, nA)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}

		case 2: // Provinsi B
			cetakCalonB(*TB, *nB)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)

			if pilihan == 1 { // edit data calon
				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di edit : ")
				fmt.Scan(&nama)
				fmt.Println("---------------------------------------")

				found := false
				for i := 0; i < *nB; i++ {
					if TB[i].calon.nama == nama {
						editCalonB(TB, i+1, *nB)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 { // hapus data calon
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di hapus : ")
				fmt.Scan(&nama)
				fmt.Println("----------------------------------------")

				found := false
				for i := 0; i < *nB; i++ {
					if TB[i].calon.nama == nama {
						hapusCalonB(TB, i+1, nB)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}

		case 3: // Provinsi C
			cetakCalonC(*TC, *nC)
			fmt.Println("--------------")
			fmt.Println("Daftar Menu : ")
			fmt.Println("1. Edit data calon")
			fmt.Println("2. Hapus data calon")
			fmt.Println("3. Kembali")
			fmt.Println()
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihan)

			if pilihan == 1 { // edit data calon
				fmt.Println("---------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di edit : ")
				fmt.Scan(&nama)
				fmt.Println("---------------------------------------")

				found := false
				for i := 0; i < *nC; i++ {
					if TC[i].calon.nama == nama {
						editCalonC(TC, i+1, *nC)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)

			} else if pilihan == 2 { // hapus data calon
				fmt.Println("----------------------------------------")
				fmt.Print("Masukkan nama calon yang ingin di hapus : ")
				fmt.Scan(&nama)
				fmt.Println("----------------------------------------")

				found := false
				for i := 0; i < *nC; i++ {
					if TC[i].calon.nama == nama {
						hapusCalonC(TC, i+1, nC)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Nama calon tidak ditemukan")
				}
				dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			}
		}

	case 3: // kembali
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

// Fungsi helper untuk mencari calon berdasarkan nama
func findCalonByName[T arrProvA | arrProvB | arrProvC](arr T, n int, nama string) int {
	for i := 0; i < n; i++ {
		if arr[i].calon.nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengedit calon berdasarkan nama
func editCalonByName(arr interface{}, nama string, n int) bool {
	var idx int = -1

	switch a := arr.(type) {
	case *arrProvA:
		idx = findCalonByName(*a, n, nama)
		if idx != -1 {
			var newNama, newPartai string
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&newNama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&newPartai)

			(*a)[idx].calon.nama = newNama
			(*a)[idx].calon.partai = newPartai
			return true
		}
	case *arrProvB:
		idx = findCalonByName(*a, n, nama)
		if idx != -1 {
			var newNama, newPartai string
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&newNama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&newPartai)

			(*a)[idx].calon.nama = newNama
			(*a)[idx].calon.partai = newPartai
			return true
		}
	case *arrProvC:
		idx = findCalonByName(*a, n, nama)
		if idx != -1 {
			var newNama, newPartai string
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&newNama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&newPartai)

			(*a)[idx].calon.nama = newNama
			(*a)[idx].calon.partai = newPartai
			return true
		}
	}
	return false
}

// Fungsi untuk menghapus calon berdasarkan nama
func hapusCalonByName(arr interface{}, nama string, n *int) bool {
	var idx int = -1

	switch a := arr.(type) {
	case *arrProvA:
		idx = findCalonByName(*a, *n, nama)
		if idx != -1 {
			for i := idx; i < *n-1; i++ {
				(*a)[i].calon = (*a)[i+1].calon
			}
			*n--
			return true
		}
	case *arrProvB:
		idx = findCalonByName(*a, *n, nama)
		if idx != -1 {
			for i := idx; i < *n-1; i++ {
				(*a)[i].calon = (*a)[i+1].calon
			}
			*n--
			return true
		}
	case *arrProvC:
		idx = findCalonByName(*a, *n, nama)
		if idx != -1 {
			for i := idx; i < *n-1; i++ {
				(*a)[i].calon = (*a)[i+1].calon
			}
			*n--
			return true
		}
	}
	return false
}

// Menampilkan menu data pemilih
func menuDataPemilih(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int

	for {
		fmt.Println("\n===============")
		fmt.Println(" Data Pemilih  ")
		fmt.Println("===============")
		fmt.Println("1. Tambah Data Pemilih")
		fmt.Println("2. Edit Data Pemilih")
		fmt.Println("3. Hapus Data Pemilih")
		fmt.Println("4. Cari Pemilih")
		fmt.Println("5. Kembali")
		fmt.Println()

		fmt.Print("Pilihan anda: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahDataPemilih(TA, TB, TC, nPA, nPB, nPC)

		case 2:
			var provinsi int
			fmt.Println("\nPilih Provinsi:")
			fmt.Println("1. Provinsi A")
			fmt.Println("2. Provinsi B")
			fmt.Println("3. Provinsi C")
			fmt.Print("Pilihan: ")
			fmt.Scan(&provinsi)

			switch provinsi {
			case 1:
				editPemilihA(TA, nPA)
			case 2:
				editPemilihB(TB, nPB)
			case 3:
				editPemilihC(TC, nPC)
			}

		case 3:
			var provinsi int
			fmt.Println("\nPilih Provinsi:")
			fmt.Println("1. Provinsi A")
			fmt.Println("2. Provinsi B")
			fmt.Println("3. Provinsi C")
			fmt.Print("Pilihan: ")
			fmt.Scan(&provinsi)

			switch provinsi {
			case 1:
				hapusPemilihA(TA, nPA)
			case 2:
				hapusPemilihB(TB, nPB)
			case 3:
				hapusPemilihC(TC, nPC)
			}

		case 4:
			menuCariPemilih()
			var pilihanProvinsi int
			fmt.Print("Pilihan anda : ")
			fmt.Scan(&pilihanProvinsi)

			switch pilihanProvinsi {
			case 1:
				searchPemilihA(*TA, *nPA)
			case 2:
				searchPemilihB(*TB, *nPB)
			case 3:
				searchPemilihC(*TC, *nPC)
			}

		case 5:
			menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
			return
		}
	}
}

// Menampilkan menu untuk KPU
func menuKPU(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihan int

	fmt.Println()
	fmt.Println("======================")
	fmt.Println("Anda Masuk Sebagai KPU")
	fmt.Println("======================")
	fmt.Println()
	fmt.Println("Daftar Menu : ")
	fmt.Println("1. Data Calon")
	fmt.Println("2. Data Pemilih")
	fmt.Println("3. Atur Tanggal Pemilu")
	fmt.Println("4. Hasil Pemilu")
	fmt.Println("5. Tampilkan Data Terurut")
	fmt.Println("6. Atur Ambang Batas")
	fmt.Println("7. Keluar")

	fmt.Println()
	fmt.Print("Pilihan anda : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1: // data calon
		dataCALON(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	case 2: // data pemilih
		menuDataPemilih(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC) // Ganti dataPEMILIH menjadi menuDataPemilih
	case 3: // Atur tanggal pemilu
		aturTanggal(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	case 4: // Hasil Pemilu
		hasilPEMILU(*TA, *TB, *TC, *nA, *nB, *nC, *nPA, *nPB, *nPC)
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	case 5:
		tampilkanDataTerurut(*TA, *TB, *TC, *nA, *nB, *nC)
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	case 6:
		aturAmbangBatas()
		menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	case 7:
		menuUtama(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	}
}

// Fungsi sorting utama yang menangani semua jenis pengurutan
func sortData[T arrProvA | arrProvB | arrProvC](arr *T, n int, by string, ascending bool, method string) {
	if n <= 0 {
		return
	}

	if method == "selection" {
		// Selection sort
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				var shouldSwap bool
				switch by {
				case "nama":
					if ascending {
						shouldSwap = (*arr)[j].calon.nama < (*arr)[minIdx].calon.nama
					} else {
						shouldSwap = (*arr)[j].calon.nama > (*arr)[minIdx].calon.nama
					}
				case "partai":
					if ascending {
						shouldSwap = (*arr)[j].calon.partai < (*arr)[minIdx].calon.partai
					} else {
						shouldSwap = (*arr)[j].calon.partai > (*arr)[minIdx].calon.partai
					}
				case "suara":
					if ascending {
						shouldSwap = (*arr)[j].totalSuaraCalon < (*arr)[minIdx].totalSuaraCalon
					} else {
						shouldSwap = (*arr)[j].totalSuaraCalon > (*arr)[minIdx].totalSuaraCalon
					}
				}
				if shouldSwap {
					minIdx = j
				}
			}
			if minIdx != i {
				(*arr)[i], (*arr)[minIdx] = (*arr)[minIdx], (*arr)[i]
			}
		}
	} else {
		// Insertion sort
		for i := 1; i < n; i++ {
			key := (*arr)[i]
			j := i - 1
			shouldMove := true
			for j >= 0 && shouldMove {
				switch by {
				case "nama":
					if ascending {
						shouldMove = (*arr)[j].calon.nama > key.calon.nama
					} else {
						shouldMove = (*arr)[j].calon.nama < key.calon.nama
					}
				case "partai":
					if ascending {
						shouldMove = (*arr)[j].calon.partai > key.calon.partai
					} else {
						shouldMove = (*arr)[j].calon.partai < key.calon.partai
					}
				case "suara":
					if ascending {
						shouldMove = (*arr)[j].totalSuaraCalon > key.totalSuaraCalon
					} else {
						shouldMove = (*arr)[j].totalSuaraCalon < key.totalSuaraCalon
					}
				}

				if shouldMove {
					(*arr)[j+1] = (*arr)[j]
					j--
				}
			}
			(*arr)[j+1] = key
		}
	}
}

// Fungsi tampilkan data terurut yang menggunakan fungsi sorting baru
func tampilkanDataTerurut(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC int) {
	var urutan, metode, kategori int

	// Validasi input kategori
	valid := false
	for !valid {
		fmt.Println("\nPilih kategori pengurutan:")
		fmt.Println("1. Nama")
		fmt.Println("2. Partai")
		fmt.Println("3. Total Suara")
		fmt.Print("Pilihan: ")
		fmt.Scan(&kategori)
		if kategori >= 1 && kategori <= 3 {
			valid = true
		}
	}

	// Validasi input metode
	valid = false
	for !valid {
		fmt.Println("\nPilih metode pengurutan:")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Pilihan: ")
		fmt.Scan(&metode)
		if metode >= 1 && metode <= 2 {
			valid = true
		}
	}

	// Validasi input urutan
	valid = false
	for !valid {
		fmt.Println("\nPilih urutan:")
		fmt.Println("1. Ascending (A-Z / Kecil ke Besar)")
		fmt.Println("2. Descending (Z-A / Besar ke Kecil)")
		fmt.Print("Pilihan: ")
		fmt.Scan(&urutan)
		if urutan >= 1 && urutan <= 2 {
			valid = true
		}
	}

	// Tentukan kategori pengurutan
	by := "nama"
	switch kategori {
	case 1:
		by = "nama"
	case 2:
		by = "partai"
	case 3:
		by = "suara"
	}

	ascending := urutan == 1
	sortMethod := map[int]string{1: "selection", 2: "insertion"}[metode]

	// Buat salinan array
	tempTA := TA
	tempTB := TB
	tempTC := TC

	// Lakukan pengurutan menggunakan fungsi sorting baru
	sortData(&tempTA, nA, by, ascending, sortMethod)
	sortData(&tempTB, nB, by, ascending, sortMethod)
	sortData(&tempTC, nC, by, ascending, sortMethod)

	// Tampilkan header hasil pengurutan
	fmt.Println("\nHasil Pengurutan:")
	fmt.Printf("Kategori: %s\n", by)
	fmt.Printf("Metode: %s\n", map[int]string{1: "Selection Sort", 2: "Insertion Sort"}[metode])
	fmt.Printf("Urutan: %s\n", map[bool]string{true: "Ascending (A-Z / Kecil ke Besar)", false: "Descending (Z-A / Besar ke Kecil)"}[ascending])
	fmt.Println()

	// Tampilkan data terurut
	if nA > 0 {
		fmt.Println("Provinsi A:")
		for i := 0; i < nA; i++ {
			fmt.Printf("%d. %-20s - %-15s (%d suara)\n",
				i+1, tempTA[i].calon.nama, tempTA[i].calon.partai, tempTA[i].totalSuaraCalon)
		}
	} else {
		fmt.Println("Provinsi A: Tidak ada data")
	}

	if nB > 0 {
		fmt.Println("\nProvinsi B:")
		for i := 0; i < nB; i++ {
			fmt.Printf("%d. %-20s - %-15s (%d suara)\n",
				i+1, tempTB[i].calon.nama, tempTB[i].calon.partai, tempTB[i].totalSuaraCalon)
		}
	} else {
		fmt.Println("\nProvinsi B: Tidak ada data")
	}

	if nC > 0 {
		fmt.Println("\nProvinsi C:")
		for i := 0; i < nC; i++ {
			fmt.Printf("%d. %-20s - %-15s (%d suara)\n",
				i+1, tempTC[i].calon.nama, tempTC[i].calon.partai, tempTC[i].totalSuaraCalon)
		}
	} else {
		fmt.Println("\nProvinsi C: Tidak ada data")
	}
}

// Mengatur tanggal pemilu
func aturTanggal(A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var konfirmasi string

	konfirmasi = "n"
	for konfirmasi != "y" {
		tanggalMulai(A)
		tanggalSelesai(A)
		fmt.Println("Pemilu berlangsung selama : ", selisihTanggal(*A), "hari")
		fmt.Println("Tanggal pemilu dimulai    : ", A.dd1, "-", A.mm1, "-", A.yy1)
		fmt.Println("Tanggal pemilu selesai    : ", A.dd2, "-", A.mm2, "-", A.yy2)
		fmt.Print("Konfirmasi (y/n) : ")
		fmt.Scan(&konfirmasi)

		fmt.Println()
		if konfirmasi == "y" {
			fmt.Println("------------------------------------")
			fmt.Println("* *Tanggal pemilu berhasil diatur* *")
			fmt.Println("------------------------------------")
			fmt.Println()
			menuKPU(A, TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
		}
	}
}

// Mengatur tanggal mulai pemilu
func tanggalMulai(A *date) {
	fmt.Println("Masukkan tanggal pemilu dimulai : (dd/mm/yyyy)")
	fmt.Scan(&A.dd1, &A.mm1, &A.yy1)
}

// Mengatur tanggal selesai pemilu
func tanggalSelesai(A *date) {
	fmt.Println("Masukkan tanggal pemilu selesai : (dd/mm/yyyy)")
	fmt.Scan(&A.dd2, &A.mm2, &A.yy2)
}

// Menghitung selisih tanggal pemilu
func selisihTanggal(A date) int {
	var selisih int
	if A.mm1 <= A.mm2 {
		selisih = A.dd2 - A.dd1
	} else {
		selisih = A.dd2 - A.dd1
		selisih += 30
	}
	return selisih
}

// Mengecek apakah tanggal pemilu valid
func cekTanggal(A date, tanggal, bulan int) bool {
	var bener bool
	if tanggal >= A.dd1 && tanggal <= A.dd2 {
		bener = true
	} else {
		bener = false
	}
	return bener
}

// Mengisi data calon untuk semua provinsi
func isiCalon(TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC *int) {
	var namaA, partaiA string
	var lanjut string

	fmt.Println("\n=================================")
	fmt.Println("    PENDAFTARAN CALON PEMILU    ")
	fmt.Println("=================================")

	// Input untuk Provinsi A
	fmt.Println("\nProvinsi A:")
	fmt.Println("---------------------------------")
	for {
		fmt.Printf("\nDaftar Calon ke-%d:\n", *nA+1)
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&namaA)
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&partaiA)

		TA[*nA].calon.nama = namaA
		TA[*nA].calon.partai = partaiA
		*nA++

		fmt.Println("\n>>> Data calon telah tercatat! <<<")
		fmt.Print("\nIngin menambah calon lagi? (y/n): ")
		fmt.Scan(&lanjut)
		if lanjut != "y" {
			break
		}
	}

	// Input untuk Provinsi B
	fmt.Println("\nProvinsi B:")
	fmt.Println("---------------------------------")
	for {
		fmt.Printf("\nDaftar Calon ke-%d:\n", *nB+1)
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&namaA)
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&partaiA)

		TB[*nB].calon.nama = namaA
		TB[*nB].calon.partai = partaiA
		*nB++

		fmt.Println("\n>>> Data calon telah tercatat! <<<")
		fmt.Print("\nIngin menambah calon lagi? (y/n): ")
		fmt.Scan(&lanjut)
		if lanjut != "y" {
			break
		}
	}

	// Input untuk Provinsi C
	fmt.Println("\nProvinsi C:")
	fmt.Println("---------------------------------")
	for {
		fmt.Printf("\nDaftar Calon ke-%d:\n", *nC+1)
		fmt.Print("Masukkan nama calon: ")
		fmt.Scan(&namaA)
		fmt.Print("Masukkan nama partai: ")
		fmt.Scan(&partaiA)

		TC[*nC].calon.nama = namaA
		TC[*nC].calon.partai = partaiA
		*nC++

		fmt.Println("\n>>> Data calon telah tercatat! <<<")
		fmt.Print("\nIngin menambah calon lagi? (y/n): ")
		fmt.Scan(&lanjut)
		if lanjut != "y" {
			break
		}
	}

	// Tampilkan ringkasan
	fmt.Println("\n=================================")
	fmt.Println("    RINGKASAN DATA CALON PEMILU    ")
	fmt.Println("=================================")

	if *nA > 0 {
		fmt.Println("\nProvinsi A:")
		cetakCalonA(*TA, *nA)
	}

	if *nB > 0 {
		fmt.Println("\nProvinsi B:")
		cetakCalonB(*TB, *nB)
	}

	if *nC > 0 {
		fmt.Println("\nProvinsi C:")
		cetakCalonC(*TC, *nC)
	}

	fmt.Println("\nPendaftaran calon selesai!")
	fmt.Println("=================================")
}

// Mengedit data calon untuk Provinsi A
func editCalonA(TA *arrProvA, noUrut, nA int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TA[noUrut-1].calon.nama = nama
	TA[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TA[noUrut-1].calon.nama, TA[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi A : ")
		cetakCalonA(*TA, nA)
	} else {
		editCalonA(TA, noUrut, nA)
	}
}

// Mengedit data calon untuk Provinsi B
func editCalonB(TB *arrProvB, noUrut, nB int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TB[noUrut-1].calon.nama = nama
	TB[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TB[noUrut-1].calon.nama, TB[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi B : ")
		cetakCalonB(*TB, nB)
	} else {
		editCalonB(TB, noUrut, nB)
	}
}

// Mengedit data calon untuk Provinsi C
func editCalonC(TC *arrProvC, noUrut, nC int) {
	var nama, partai, konfirmasi string
	fmt.Println("Masukkan data yang di edit: ")
	fmt.Scan(&nama, &partai)
	TC[noUrut-1].calon.nama = nama
	TC[noUrut-1].calon.partai = partai
	fmt.Println("======== Data yang di edit =========")
	fmt.Println(TC[noUrut-1].calon.nama, TC[noUrut-1].calon.partai)
	fmt.Print("Konfirmasi? (y/n) ")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		fmt.Println("Data berhasil di edit!")
		fmt.Println("Data Calon Provinsi C : ")
		cetakCalonC(*TC, nC)
	} else {
		editCalonC(TC, noUrut, nC)
	}
}

// Menghapus data calon untuk Provinsi A
func hapusCalonA(TA *arrProvA, noUrut int, nA *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TA[noUrut-1].calon.nama, TA[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nA-(noUrut-1); i++ {
			TA[i].calon = TA[i+1].calon
		}
		*nA = *nA - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonA(*TA, *nA)
	} else {
		fmt.Println("Data tidak dihapus")
	}
}

// Menghapus data calon untuk Provinsi B
func hapusCalonB(TB *arrProvB, noUrut int, nB *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TB[noUrut-1].calon.nama, TB[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nB-(noUrut-1); i++ {
			TB[i].calon = TB[i+1].calon
		}
		*nB = *nB - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonB(*TB, *nB)
	} else {
		fmt.Println("Data tidak dihapus")
	}
}

// Menghapus data calon untuk Provinsi C
func hapusCalonC(TC *arrProvC, noUrut int, nC *int) {
	var konfirmasi string
	fmt.Println("Apakah anda yakin akan menghapus calon nomor", noUrut, TC[noUrut-1].calon.nama, TC[noUrut-1].calon.partai, " ?")
	fmt.Print("Konfirmasi? (y/n)")
	fmt.Scan(&konfirmasi)
	if konfirmasi == "y" {
		for i := noUrut - 1; i < *nC-(noUrut-1); i++ {
			TC[i].calon = TC[i+1].calon
		}
		*nC = *nC - 1
		fmt.Println("Data berhasil dihapus!")
		cetakCalonC(*TC, *nC)
	} else {
		fmt.Println("Data tidak dihapus")
	}
}

// Mencetak data calon untuk Provinsi A
func cetakCalonA(TA arrProvA, nA int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi A : ")
	fmt.Println("--------------")
	for i := 0; i < nA; i++ {
		fmt.Printf("%d. %s %s \n", i+1, TA[i].calon.nama, TA[i].calon.partai)
	}
	fmt.Println()
}

// Mencetak data calon untuk Provinsi B
func cetakCalonB(TB arrProvB, nB int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi B : ")
	fmt.Println("--------------")
	for j := 0; j < nB; j++ {
		fmt.Printf("%d. %s %s \n", j+1, TB[j].calon.nama, TB[j].calon.partai)
	}
	fmt.Println()
}

// Mencetak data calon untuk Provinsi C
func cetakCalonC(TC arrProvC, nC int) {
	fmt.Println("--------------")
	fmt.Println("Provinsi C : ")
	fmt.Println("--------------")
	for k := 0; k < nC; k++ {
		fmt.Printf("%d. %s %s \n", k+1, TC[k].calon.nama, TC[k].calon.partai)
	}
	fmt.Println()
}

// Mencari calon berdasarkan nama untuk Provinsi A
func searchNamaCalonA(TA arrProvA, nA int) {
	var nama string
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)

	idx := sequentialSearch(TA, nA, nama, "nama")
	if idx != -1 {
		fmt.Println("Data ditemukan!")
		fmt.Printf("%d. %s %s\n", idx+1, TA[idx].calon.nama, TA[idx].calon.partai)
	} else {
		fmt.Println("Data tidak ditemukan!")
	}
}

// Mencari calon berdasarkan nama untuk Provinsi B
func searchNamaCalonB(TB arrProvB, nB int) {
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	i := 0
	for i < nB && !ketemu {
		if TB[i].calon.nama == nama {
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. %s %s\n", i+1, TB[i].calon.nama, TB[i].calon.partai)
			fmt.Println("============================================")
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

// Mencari calon berdasarkan nama untuk Provinsi C
func searchNamaCalonC(TC arrProvC, nC int) {
	var nama string
	var ketemu bool
	fmt.Print("Masukkan nama yang dicari: ")
	fmt.Scan(&nama)
	ketemu = false
	i := 0
	for i < nC && !ketemu {
		if TC[i].calon.nama == nama {
			fmt.Println("============================================")
			fmt.Println("              ~Data ditemukan!~             ")
			fmt.Println("--------------------------------------------")
			fmt.Printf("%d. %s %s\n", i+1, TC[i].calon.nama, TC[i].calon.partai)
			fmt.Println("============================================")
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

// Mencari calon berdasarkan partai untuk Provinsi A
func searchPartaiCalonA(TA arrProvA, nA int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nA; i++ {
		if TA[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TA[i].calon.nama, TA[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

// Mencari calon berdasarkan partai untuk Provinsi B
func searchPartaiCalonB(TB arrProvB, nB int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nB; i++ {
		if TB[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TB[i].calon.nama, TB[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

// Mencari calon berdasarkan partai untuk Provinsi C
func searchPartaiCalonC(TC arrProvC, nC int) {
	var partai string
	var ketemu bool
	fmt.Print("Masukkan partai yang dicari: ")
	fmt.Scan(&partai)
	ketemu = false
	for i := 0; i < nC; i++ {
		if TC[i].calon.partai == partai {
			fmt.Printf("%d. %s %s\n", i+1, TC[i].calon.nama, TC[i].calon.partai)
			ketemu = true
		}
	}
	if ketemu == true {
		fmt.Println("============================================")
		fmt.Println("              ~Data ditemukan!~             ")
		fmt.Println("============================================")
	} else if ketemu == false {
		fmt.Println("============================================")
		fmt.Println("            ~Data tidak ditemukan!~         ")
		fmt.Println("============================================")
	}
}

// Mencetak data pemilih untuk Provinsi A
func cetakPemilihA(TA arrProvA, nPA int) {
	var calon int
	fmt.Println("Provinsi A : ")
	if nPA > 0 {
		for i := 0; i < nPA; i++ {
			calon = TA[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TA[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TA[i].pemilih.pilihan, TA[calon-1].calon.nama, TA[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

// Mencetak data pemilih untuk Provinsi B
func cetakPemilihB(TB arrProvB, nPB int) {
	var calon int
	fmt.Println("Provinsi B : ")
	if nPB > 0 {
		for i := 0; i < nPB; i++ {
			calon = TB[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TB[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TB[i].pemilih.pilihan, TB[calon-1].calon.nama, TB[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

// Mencetak data pemilih untuk Provinsi C
func cetakPemilihC(TC arrProvC, nPC int) {
	var calon int
	fmt.Println("Provinsi C : ")
	if nPC > 0 {
		for i := 0; i < nPC; i++ {
			calon = TC[i].pemilih.pilihan
			fmt.Printf("%d. Nama: %s \n", i+1, TC[i].pemilih.nama)
			fmt.Printf("   Pilihan: %d. %s - %s \n", TC[i].pemilih.pilihan, TC[calon-1].calon.nama, TC[calon-1].calon.partai)
		}
		fmt.Println("")
	}
}

// Mencari pemilih berdasarkan nama untuk Provinsi A
func searchPemilihA(TA arrProvA, nPA int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang dicari: ")
	fmt.Scan(&nama)

	found := false
	for i := 0; i < nPA; i++ {
		if TA[i].pemilih.nama == nama {
			fmt.Println("\nData Pemilih ditemukan!")
			fmt.Println("Nama:", TA[i].pemilih.nama)
			fmt.Println("Pilihan:", TA[i].pemilih.pilihan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Mencari pemilih berdasarkan nama untuk Provinsi B
func searchPemilihB(TB arrProvB, nPB int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang dicari: ")
	fmt.Scan(&nama)

	found := false
	for i := 0; i < nPB; i++ {
		if TB[i].pemilih.nama == nama {
			fmt.Println("\nData Pemilih ditemukan!")
			fmt.Println("Nama:", TB[i].pemilih.nama)
			fmt.Println("Pilihan:", TB[i].pemilih.pilihan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Mencari pemilih berdasarkan nama untuk Provinsi C
func searchPemilihC(TC arrProvC, nPC int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang dicari: ")
	fmt.Scan(&nama)

	found := false
	for i := 0; i < nPC; i++ {
		if TC[i].pemilih.nama == nama {
			fmt.Println("\nData Pemilih ditemukan!")
			fmt.Println("Nama:", TC[i].pemilih.nama)
			fmt.Println("Pilihan:", TC[i].pemilih.pilihan)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Menghitung total suara untuk semua calon
func hitungSemuaCALON(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC, nPA, nPB, nPC int) {
	// Reset total suara terlebih dahulu
	for i := 0; i < nA; i++ {
		TA[i].totalSuaraCalon = 0
	}
	for i := 0; i < nB; i++ {
		TB[i].totalSuaraCalon = 0
	}
	for i := 0; i < nC; i++ {
		TC[i].totalSuaraCalon = 0
	}

	// Hitung suara untuk setiap calon berdasarkan pilihan pemilih
	for i := 0; i < nPA; i++ {
		pilihan := TA[i].pemilih.pilihan
		if pilihan > 0 && pilihan <= nA {
			TA[pilihan-1].totalSuaraCalon++
		}
	}

	for i := 0; i < nPB; i++ {
		pilihan := TB[i].pemilih.pilihan
		if pilihan > 0 && pilihan <= nB {
			TB[pilihan-1].totalSuaraCalon++
		}
	}

	for i := 0; i < nPC; i++ {
		pilihan := TC[i].pemilih.pilihan
		if pilihan > 0 && pilihan <= nC {
			TC[pilihan-1].totalSuaraCalon++
		}
	}

	// Lanjutkan dengan sorting dan menampilkan hasil
	sortingTOTAL(TA, TB, TC, nPA, nPB, nPC)

	// Hitung dan tampilkan total suara partai
	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi A: ")
	hitungSuaraPartai(TA, nA)

	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi B: ")
	hitungSuaraPartai(TB, nB)

	fmt.Println()
	fmt.Println("----------------------------------")
	fmt.Println("Total suara PARTAI di Provinsi C: ")
	hitungSuaraPartai(TC, nC)
}

// Fungsi helper untuk menghitung suara partai per provinsi
func hitungSuaraPartai[T arrProvA | arrProvB | arrProvC](p T, n int) {
	// Map untuk menyimpan total suara per partai
	partaiSuara := make(map[string]int)

	// Hitung total suara untuk setiap partai
	for i := 0; i < n; i++ {
		partai := p[i].calon.partai
		suara := p[i].totalSuaraCalon
		partaiSuara[partai] += suara
	}

	// Tampilkan hasil
	for partai, suara := range partaiSuara {
		if partai != "" {
			fmt.Printf("%s - %d suara\n", partai, suara)
		}
	}
}

// Mengurutkan total suara untuk semua calon
func sortingTOTAL(TA arrProvA, TB arrProvB, TC arrProvC, nPA, nPB, nPC int) {
	var tempA, tempB, tempC provinsi

	for i := 0; i < nPA; i++ {
		for j := i + 1; j < nPA; j++ {
			if TA[i].totalSuaraCalon < TA[j].totalSuaraCalon {
				tempA = TA[i]
				TA[i] = TA[j]
				TA[j] = tempA
			}
		}
	}

	for i := 0; i < nPB; i++ {
		for j := i + 1; j < nPB; j++ {
			if TB[i].totalSuaraCalon < TB[j].totalSuaraCalon {
				tempB = TB[i]
				TB[i] = TB[j]
				TB[j] = tempB
			}
		}
	}

	for i := 0; i < nPC; i++ {
		for j := i + 1; j < nPC; j++ {
			if TC[i].totalSuaraCalon < TC[j].totalSuaraCalon {
				tempC = TC[i]
				TC[i] = TC[j]
				TC[j] = tempC
			}
		}
	}

	// Menampilkan hasil pemilu 3 suara terbanyak untuk Provinsi A
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi A: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TA[i].calon.nama, " - ", TA[i].calon.partai, " - ", TA[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")

	// Menampilkan hasil pemilu 3 suara terbanyak untuk Provinsi B
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi B: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TB[i].calon.nama, " - ", TB[i].calon.partai, " - ", TB[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")

	// Menampilkan hasil pemilu 3 suara terbanyak untuk Provinsi C
	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("Hasil PEMILU 3 Suara terbanyak Provinsi C: ")
	for i := 0; i < 3; i++ {
		fmt.Println(TC[i].calon.nama, " - ", TC[i].calon.partai, " - ", TC[i].totalSuaraCalon, "Suara ")
	}
	fmt.Println("===========================================")
	fmt.Println()
}

// Mengatur ambang batas suara
func aturAmbangBatas() {
	fmt.Println("=====================================")
	fmt.Println("      PENGATURAN AMBANG BATAS      ")
	fmt.Println("=====================================")

	var input int
	fmt.Print("Masukkan nilai ambang batas suara: ")
	fmt.Scan(&input)

	if input > 0 {
		threshold = input
		fmt.Println("----------------------------------------")
		fmt.Printf("Ambang batas berhasil diatur: %d suara\n", threshold)
		fmt.Println("----------------------------------------")
	} else {
		fmt.Println("Error: Ambang batas harus lebih dari 0")
		aturAmbangBatas()
	}
}

// Menampilkan kandidat yang melewati ambang batas suara
func tampilkanAmbangBatas(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC int) {
	fmt.Printf("\nKandidat yang melewati ambang batas suara (%d):\n\n", threshold)

	fmt.Println("Provinsi A:")
	for i := 0; i < nA; i++ {
		if TA[i].totalSuaraCalon >= threshold {
			fmt.Printf("%s - %s (%d suara) ✓\n", TA[i].calon.nama, TA[i].calon.partai, TA[i].totalSuaraCalon)
		} else {
			fmt.Printf("%s - %s (%d suara) ✗\n", TA[i].calon.nama, TA[i].calon.partai, TA[i].totalSuaraCalon)
		}
	}

	fmt.Println("\nProvinsi B:")
	for i := 0; i < nB; i++ {
		if TB[i].totalSuaraCalon >= threshold {
			fmt.Printf("%s - %s (%d suara) ✓\n", TB[i].calon.nama, TB[i].calon.partai, TB[i].totalSuaraCalon)
		} else {
			fmt.Printf("%s - %s (%d suara) ✗\n", TB[i].calon.nama, TB[i].calon.partai, TB[i].totalSuaraCalon)
		}
	}

	fmt.Println("\nProvinsi C:")
	for i := 0; i < nC; i++ {
		if TC[i].totalSuaraCalon >= threshold {
			fmt.Printf("%s - %s (%d suara) ✓\n", TC[i].calon.nama, TC[i].calon.partai, TC[i].totalSuaraCalon)
		} else {
			fmt.Printf("%s - %s (%d suara) ✗\n", TC[i].calon.nama, TC[i].calon.partai, TC[i].totalSuaraCalon)
		}
	}

	fmt.Println("\nKeterangan:")
	fmt.Println("✓ : Melewati ambang batas")
	fmt.Println("✗ : Tidak melewati ambang batas")
}

// Menampilkan hasil pemilu
func hasilPEMILU(TA arrProvA, TB arrProvB, TC arrProvC, nA, nB, nC, nPA, nPB, nPC int) {
	fmt.Println("============================================")
	fmt.Println("              ~HASIL PEMILU~                ")
	fmt.Println("--------------------------------------------")
	hitungSemuaCALON(TA, TB, TC, nA, nB, nC, nPA, nPB, nPC)
	tampilkanAmbangBatas(TA, TB, TC, nA, nB, nC) // Menampilkan ambang batas
}

// Fungsi untuk mencari data dengan binary search (perlu dimodifikasi)
func binarySearch[T arrProvA | arrProvB | arrProvC](arr T, n int, target string, by string) int {
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		var current string

		if by == "nama" {
			current = arr[mid].calon.nama
		} else {
			current = arr[mid].calon.partai
		}

		if current == target {
			return mid
		} else if current < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Fungsi untuk input data dengan validasi
func inputData(prompt string) (string, bool) {
	var input string
	var valid bool = false

	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)
		if input != "" {
			valid = true
		}
	}
	return input, true
}

// Fungsi untuk konfirmasi tindakan
func konfirmasi(prompt string) bool {
	var jawaban string
	var selesai bool = false
	var hasil bool = false

	for !selesai {
		fmt.Print(prompt + " (y/n): ")
		fmt.Scan(&jawaban)
		if jawaban == "y" {
			hasil = true
			selesai = true
		} else if jawaban == "n" {
			hasil = false
			selesai = true
		}
	}
	return hasil
}

// Fungsi untuk menu dengan loop
func tampilkanMenu(pilihan *int, maxPilihan int) bool {
	var valid bool = false

	for !valid {
		fmt.Print("Pilihan anda: ")
		fmt.Scan(pilihan)
		if *pilihan >= 1 && *pilihan <= maxPilihan {
			valid = true
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
	return true
}

// Fungsi untuk validasi input angka
func inputAngka(prompt string, min, max int) int {
	var input int
	var valid bool = false

	for !valid {
		fmt.Print(prompt)
		fmt.Scan(&input)
		if input >= min && input <= max {
			valid = true
		} else {
			fmt.Printf("Masukkan angka antara %d dan %d!\n", min, max)
		}
	}
	return input
}

// Fungsi untuk mencari data dengan sequential search
func searchSequential(arr []string, n int, target string) (bool, int) {
	var i int = 0
	var ketemu bool = false
	var idx int = -1

	for i < n && !ketemu {
		if arr[i] == target {
			ketemu = true
			idx = i
		}
		i++
	}
	return ketemu, idx
}

func validasiInput(input string, tipe string) bool {
	switch tipe {
	case "nama":
		return len(input) >= 3 && len(input) <= 50
	case "partai":
		return len(input) >= 2 && len(input) <= 20
	case "tanggal":
		// Validasi format tanggal
		return true
	default:
		return false
	}
}

func tampilkanStatistik(TA arrProvA, TB arrProvB, TC arrProvC, nPA, nPB, nPC int) {
	totalPemilih := nPA + nPB + nPC
	fmt.Println("\nStatistik Pemilu:")
	fmt.Printf("Total Pemilih: %d\n", totalPemilih)
	fmt.Printf("Provinsi A: %d pemilih (%.2f%%)\n", nPA, float64(nPA)/float64(totalPemilih)*100)
	fmt.Printf("Provinsi B: %d pemilih (%.2f%%)\n", nPB, float64(nPB)/float64(totalPemilih)*100)
	fmt.Printf("Provinsi C: %d pemilih (%.2f%%)\n", nPC, float64(nPC)/float64(totalPemilih)*100)
}

// Memisahkan fungsi sorting menjadi lebih modular
func selectionSort[T arrProvA | arrProvB | arrProvC](arr T, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ascending {
				if arr[j].totalSuaraCalon < arr[idx].totalSuaraCalon {
					idx = j
				}
			} else {
				if arr[j].totalSuaraCalon > arr[idx].totalSuaraCalon {
					idx = j
				}
			}
		}
		if idx != i {
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
}

// Implementasi insertion sort
func insertionSort[T arrProvA | arrProvB | arrProvC](arr T, n int, by string, ascending bool) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		shouldMove := true
		for j >= 0 && shouldMove {
			switch by {
			case "nama":
				if ascending {
					shouldMove = arr[j].calon.nama > key.calon.nama
				} else {
					shouldMove = arr[j].calon.nama < key.calon.nama
				}
			case "partai":
				if ascending {
					shouldMove = arr[j].calon.partai > key.calon.partai
				} else {
					shouldMove = arr[j].calon.partai < key.calon.partai
				}
			case "suara":
				if ascending {
					shouldMove = arr[j].totalSuaraCalon > key.totalSuaraCalon
				} else {
					shouldMove = arr[j].totalSuaraCalon < key.totalSuaraCalon
				}
			}

			if shouldMove {
				arr[j+1] = arr[j]
				j--
			}
		}
		arr[j+1] = key
	}
}

// Fungsi untuk mengurutkan data dengan pilihan metode
func urutkanData[T arrProvA | arrProvB | arrProvC](arr T, n int, metode string, ascending bool) {
	if metode == "selection" {
		selectionSort(arr, n, ascending)
	} else if metode == "insertion" {
		insertionSort(arr, n, "suara", ascending) // Menambahkan parameter "suara" sebagai default
	}
}

// Fungsi untuk selection sort
func selectionSortCalon[T arrProvA | arrProvB | arrProvC](arr T, n int, by string, ascending bool) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			var shouldSwap bool
			if by == "nama" {
				if ascending {
					shouldSwap = arr[j].calon.nama < arr[minIdx].calon.nama
				} else {
					shouldSwap = arr[j].calon.nama > arr[minIdx].calon.nama
				}
			} else { // by == "suara"
				if ascending {
					shouldSwap = arr[j].totalSuaraCalon < arr[minIdx].totalSuaraCalon
				} else {
					shouldSwap = arr[j].totalSuaraCalon > arr[minIdx].totalSuaraCalon
				}
			}

			if shouldSwap {
				minIdx = j
			}
		}

		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// Tambahkan fungsi sequential search yang lebih umum
func sequentialSearch[T arrProvA | arrProvB | arrProvC](arr T, n int, target string, by string) int {
	for i := 0; i < n; i++ {
		var current string
		if by == "nama" {
			current = arr[i].calon.nama
		} else {
			current = arr[i].calon.partai
		}
		if current == target {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengedit data pemilih Provinsi A
func editPemilihA(TA *arrProvA, nPA *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan diedit: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPA && !found {
		if TA[i].pemilih.nama == nama {
			fmt.Println("Data pemilih ditemukan!")
			fmt.Println("Pilihan saat ini:", TA[i].pemilih.pilihan)

			fmt.Print("Masukkan pilihan baru: ")
			fmt.Scan(&TA[i].pemilih.pilihan)

			fmt.Println("Data berhasil diubah!")
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk menghapus data pemilih Provinsi A
func hapusPemilihA(TA *arrProvA, nPA *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan dihapus: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPA && !found {
		if TA[i].pemilih.nama == nama {
			if konfirmasi("Yakin ingin menghapus data pemilih ini?") {
				// Geser semua data setelahnya
				j := i
				for j < *nPA-1 {
					TA[j].pemilih = TA[j+1].pemilih
					j++
				}
				*nPA--
				fmt.Println("Data pemilih berhasil dihapus!")
			}
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk mengedit data pemilih Provinsi B
func editPemilihB(TB *arrProvB, nPB *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan diedit: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPB && !found {
		if TB[i].pemilih.nama == nama {
			fmt.Println("Data pemilih ditemukan!")
			fmt.Println("Pilihan saat ini:", TB[i].pemilih.pilihan)

			fmt.Print("Masukkan pilihan baru: ")
			fmt.Scan(&TB[i].pemilih.pilihan)

			fmt.Println("Data berhasil diubah!")
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk mengedit data pemilih Provinsi C
func editPemilihC(TC *arrProvC, nPC *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan diedit: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPC && !found {
		if TC[i].pemilih.nama == nama {
			fmt.Println("Data pemilih ditemukan!")
			fmt.Println("Pilihan saat ini:", TC[i].pemilih.pilihan)

			fmt.Print("Masukkan pilihan baru: ")
			fmt.Scan(&TC[i].pemilih.pilihan)

			fmt.Println("Data berhasil diubah!")
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk menghapus data pemilih Provinsi B
func hapusPemilihB(TB *arrProvB, nPB *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan dihapus: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPB && !found {
		if TB[i].pemilih.nama == nama {
			if konfirmasi("Yakin ingin menghapus data pemilih ini?") {
				j := i
				for j < *nPB-1 {
					TB[j].pemilih = TB[j+1].pemilih
					j++
				}
				*nPB--
				fmt.Println("Data pemilih berhasil dihapus!")
			}
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk menghapus data pemilih Provinsi C
func hapusPemilihC(TC *arrProvC, nPC *int) {
	var nama string
	fmt.Print("Masukkan nama pemilih yang akan dihapus: ")
	fmt.Scan(&nama)

	found := false
	i := 0
	for i < *nPC && !found {
		if TC[i].pemilih.nama == nama {
			if konfirmasi("Yakin ingin menghapus data pemilih ini?") {
				j := i
				for j < *nPC-1 {
					TC[j].pemilih = TC[j+1].pemilih
					j++
				}
				*nPC--
				fmt.Println("Data pemilih berhasil dihapus!")
			}
			found = true
		}
		i++
	}

	if !found {
		fmt.Println("Pemilih tidak ditemukan!")
	}
}

// Fungsi untuk menangani pencarian calon
func handlePencarianCalon(pilihan int, A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var pilihanCari int
	fmt.Println("\n1. Cari berdasarkan nama")
	fmt.Println("2. Cari berdasarkan partai")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanCari)

	if pilihanCari == 1 {
		switch pilihan {
		case 1:
			searchNamaCalonA(*TA, *nA)
		case 2:
			searchNamaCalonB(*TB, *nB)
		case 3:
			searchNamaCalonC(*TC, *nC)
		}
	} else if pilihanCari == 2 {
		switch pilihan {
		case 1:
			searchPartaiCalonA(*TA, *nA)
		case 2:
			searchPartaiCalonB(*TB, *nB)
		case 3:
			searchPartaiCalonC(*TC, *nC)
		}
	}
}

// Fungsi untuk menangani proses pemilihan
func handlePemilihan(pilihan int, A *date, TA *arrProvA, TB *arrProvB, TC *arrProvC, nA, nB, nC, nPA, nPB, nPC *int) {
	var nama string
	var pilihanCalon int

	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&nama)

	// Cek duplikasi pemilih
	if cekDuplikasiPemilih(nama, *TA, *TB, *TC, *nPA, *nPB, *nPC) {
		fmt.Println("Maaf, Anda sudah melakukan pemilihan!")
		return
	}

	// Tampilkan daftar calon sesuai provinsi
	switch pilihan {
	case 1:
		cetakCalonA(*TA, *nA)
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pilihanCalon)
		if validasiPemilih(nama, pilihanCalon, *nA) {
			TA[*nPA].pemilih.nama = nama
			TA[*nPA].pemilih.pilihan = pilihanCalon
			// Tambahkan suara ke calon yang dipilih
			TA[pilihanCalon-1].totalSuaraCalon++
			*nPA++
			fmt.Println("Pemilihan berhasil dicatat!")
		}
	case 2:
		cetakCalonB(*TB, *nB)
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pilihanCalon)
		if validasiPemilih(nama, pilihanCalon, *nB) {
			TB[*nPB].pemilih.nama = nama
			TB[*nPB].pemilih.pilihan = pilihanCalon
			// Tambahkan suara ke calon yang dipilih
			TB[pilihanCalon-1].totalSuaraCalon++
			*nPB++
			fmt.Println("Pemilihan berhasil dicatat!")
		}
	case 3:
		cetakCalonC(*TC, *nC)
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pilihanCalon)
		if validasiPemilih(nama, pilihanCalon, *nC) {
			TC[*nPC].pemilih.nama = nama
			TC[*nPC].pemilih.pilihan = pilihanCalon
			// Tambahkan suara ke calon yang dipilih
			TC[pilihanCalon-1].totalSuaraCalon++
			*nPC++
			fmt.Println("Pemilihan berhasil dicatat!")
		}
	}
}

func tambahDataPemilih(TA *arrProvA, TB *arrProvB, TC *arrProvC, nPA, nPB, nPC *int) {
	var provinsi int
	fmt.Println("\n=================================")
	fmt.Println("    PENDAFTARAN PEMILIH BARU     ")
	fmt.Println("=================================")

	fmt.Println("\nPilih Provinsi:")
	fmt.Println("1. Provinsi A")
	fmt.Println("2. Provinsi B")
	fmt.Println("3. Provinsi C")
	fmt.Print("Pilihan: ")
	fmt.Scan(&provinsi)

	switch provinsi {
	case 1:
		isiPemilihA(TA, nPA)
	case 2:
		isiPemilihB(TB, nPB)
	case 3:
		isiPemilihC(TC, nPC)
	default:
		fmt.Println("Pilihan provinsi tidak valid!")
	}
}
