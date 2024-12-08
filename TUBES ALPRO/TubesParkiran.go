package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const MAXDATA int = 1000

const MAXPARKIRMOBIL int = 400

const MAXPARKIRMOTOR int = 100

const MAXPETUGAS int = 5

type tPetugas struct {
	nama, username, password string
	id                       int
}

type tKendaraan struct {
	tipe, noPol, hari string
	waktu             tWaktuParkir
	vip, valet        bool
	hargaParkir       int
}

type tWaktuParkir struct {
	jamMasuk, jamKeluar, menitMasuk, menitKeluar int
}

type tabPetugas [MAXPETUGAS]tPetugas

type tabDataKendaraan [MAXDATA]tKendaraan

type tabParkirMobil [MAXPARKIRMOBIL]tKendaraan

type tabParkirMotor [MAXPARKIRMOTOR]tKendaraan

func main() {
	var inputMenu, nPetugas int
	var petugas tabPetugas
	for inputMenu != 3 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                  *SELAMAT DATANG*                                  *||*")
		fmt.Println("*||*                                    DI PARKIRAN                                     *||*")
		fmt.Println("*||*                                  *PARIS VAN JOWO*                                  *||*")
		fmt.Println("*||*                         GENA DARMA    //  FACHRI MUTHAWWA                          *||*")
		fmt.Println("*||*                       (103032330095)  //  (103032330141)                           *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*** Menu Utama ***")
		fmt.Println("1. Admin")
		fmt.Println("2. Login Petugas Parkir")
		fmt.Println("3. Keluar")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&inputMenu)
		ClearScreen()
		if inputMenu == 1 {
			adminMenu(&petugas, &nPetugas)
			ClearScreen()
		} else if inputMenu == 2 {
			loginMenu(petugas, nPetugas)
			ClearScreen()
		}
	}
}

func adminMenu(petugas *tabPetugas, nPetugas *int) {
	var input int
	for input != 5 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                    *MENU ADMIN*                                    *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Tambah Data Petugas Parkir")
		fmt.Println("2. Edit Data Petugas Parkir")
		fmt.Println("3. Hapus Data Petugas Parkir")
		fmt.Println("4. Cetak Data Petugas Parkir")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			tambahPetugas(&*petugas, &*nPetugas)
			ClearScreen()
		} else if input == 2 {
			editPetugas(&*petugas, *nPetugas)
			ClearScreen()
		} else if input == 3 {
			hapusPetugas(&*petugas, &*nPetugas)
			ClearScreen()
		} else if input == 4 {
			cetakPetugas(*petugas, *nPetugas)
			ClearScreen()
		}
	}
}

func tambahPetugas(petugas *tabPetugas, nPetugas *int) {
	/* IS : Array petugas dengan banyak elemen n terdefinisi sembarang
	   FS : Petugas ditambahkan ke dalam array petugas sebanyak nPetugas dengan ID terurut ascending
	*/
	var i, n, id int
	var usn string
	fmt.Println("============================================================================================")
	fmt.Println("*||*                                *MENU TAMBAH PETUGAS*                               *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan Banyaknya Petugas Baru : ")
	fmt.Scan(&n)
	ClearScreen()
	if n > MAXPETUGAS-*nPetugas {
		fmt.Println("Banyaknya petugas melebihi batasan maksimum!")
		fmt.Printf("Banyaknya petugas baru yang ditambahkan berubah dari %d menjadi %d\n", n, MAXPETUGAS-*nPetugas)
		n = MAXPETUGAS - *nPetugas
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}
	i = *nPetugas
	for i < n+*nPetugas && i < MAXPETUGAS {
		fmt.Printf("Masukkan ID Petugas %d : ", i+1)
		fmt.Scan(&id)
		for !validId(*petugas, i, id) {
			fmt.Println("ID Petugas tidak valid!")
			fmt.Printf("Masukkan ID Petugas %d : ", i+1)
			fmt.Scan(&id)
		}
		petugas[i].id = id
		fmt.Printf("Masukkan Nama Petugas %d : ", i+1)
		fmt.Scan(&petugas[i].nama)
		fmt.Printf("Masukkan Username Petugas %d : ", i+1)
		fmt.Scan(&usn)
		for !validUsn(*petugas, i, usn) {
			fmt.Println("Username telah digunakan!")
			fmt.Printf("Masukkan Username Petugas %d : ", i+1)
			fmt.Scan(&usn)
		}
		petugas[i].username = usn
		fmt.Printf("Masukkan Password Petugas %d : ", i+1)
		fmt.Scan(&petugas[i].password)
		fmt.Println("Petugas berhasil ditambahkan :)")
		fmt.Println("--------------------------------------------------------------------------------------------")
		i++
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}
	*nPetugas = i
	urutID(&*petugas, *nPetugas)
	ClearScreen()
}

func editPetugas(petugas *tabPetugas, nPetugas int) {
	var input int
	for input != 5 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                *MENU EDIT PETUGAS*                                 *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Petugas yang ingin diedit :")
		fmt.Println("1. ID Petugas Parkir")
		fmt.Println("2. Nama Petugas Parkir")
		fmt.Println("3. Username Petugas Parkir")
		fmt.Println("4. Password Petugas Parkir")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 || input == 2 || input == 3 || input == 4 {
			editDataPetugas(&*petugas, nPetugas, input)
			ClearScreen()
		}
	}
}

func editDataPetugas(petugas *tabPetugas, nPetugas, input int) {
	/* IS : Array petugas dengan banyak elemen n terdefinisi
	   FS : Data Petugas berhasil diubah berdasarkan data petugas yang dipilih
	   		(contoh : ingin merubah username maka username petugas dengan ID tertentu akan berubah)
	*/
	var idx, id, idEdit int
	var usn string
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan ID Petugas yang ingin diedit : ")
	fmt.Scan(&idEdit)
	idx = searchID(*petugas, nPetugas, idEdit)
	if idx == -1 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*                 *ID TERSEBUT TIDAK TERDAPAT DI DALAM DATA PETUGAS!*                *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		if input == 1 {
			fmt.Print("Masukkan ID Petugas yang baru : ")
			fmt.Scan(&id)
			for !validId(*petugas, nPetugas, id) {
				fmt.Println("ID Petugas tidak valid!")
				fmt.Print("Masukkan ID Petugas yang baru : ")
				fmt.Scan(&id)
			}
			petugas[idx].id = id
			urutID(&*petugas, nPetugas)
			fmt.Println("ID Petugas berhasil diedit :)")
		} else if input == 2 {
			fmt.Print("Masukkan Nama Petugas yang baru : ")
			fmt.Scan(&petugas[idx].nama)
			fmt.Println("Nama Petugas berhasil diedit :)")
		} else if input == 3 {
			fmt.Print("Masukkan Username Petugas yang baru : ")
			fmt.Scan(&usn)
			for !validUsn(*petugas, nPetugas, usn) {
				fmt.Println("Username telah digunakan!")
				fmt.Printf("Masukkan Username Petugas yang baru : ")
				fmt.Scan(&usn)
			}
			petugas[idx].username = usn
			fmt.Println("Username Petugas berhasil diedit :)")
		} else if input == 4 {
			fmt.Print("Masukkan Password Petugas yang baru : ")
			fmt.Scan(&petugas[idx].password)
			fmt.Println("Password Petugas berhasil diedit :)")
		}
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}
}

func hapusPetugas(petugas *tabPetugas, nPetugas *int) {
	/* IS : Array petugas dengan banyak elemen n terdefinisi
	   FS : Data petugas dengan ID tertentu terhapus
	*/
	var idx, i, idHapus int
	fmt.Println("============================================================================================")
	fmt.Println("*||*                                *MENU HAPUS PETUGAS*                                *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan ID Petugas yang ingin dihapus : ")
	fmt.Scan(&idHapus)
	idx = searchID(*petugas, *nPetugas, idHapus)
	if idx == -1 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*                 *ID TERSEBUT TIDAK TERDAPAT DI DALAM DATA PETUGAS!*                *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		i = idx
		for i < *nPetugas-1 {
			petugas[i] = petugas[i+1]
			i++
		}
		*nPetugas--
		fmt.Printf("Semua data Petugas dengan id %d berhasil dihapus :)\n", idHapus)
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}
}

func cetakPetugas(petugas tabPetugas, nPetugas int) {
	var i int
	fmt.Println("Data Petugas Tiket Parkir Terurut Berdasarkan ID: ")
	for i = 0; i < nPetugas; i++ {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("ID Petugas : %d\n", petugas[i].id)
		fmt.Printf("Nama Petugas : %s\n", petugas[i].nama)
		fmt.Printf("Username Petugas : %s\n", petugas[i].username)
		fmt.Printf("Password Petugas : %s\n", petugas[i].password)
		fmt.Println("--------------------------------------------------------------------------------------------")
	}
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func loginMenu(petugas tabPetugas, nPetugas int) {
	/* IS : Array petugas dengan banyak elemen n terdefinisi
	   FS : Petugas dapat melakukan proses login dengan username dan password yang benar
	*/
	var input, idx int
	var username, password string
	var statusLogin bool
	fmt.Println("============================================================================================")
	fmt.Println("*||*                           *MENU LOG IN PETUGAS PARKIR*                             *||*")
	fmt.Println("============================================================================================")
	for input != 2 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Log In")
		fmt.Println("2. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			statusLogin = false
			for !statusLogin {
				fmt.Print("Masukkan Username : ")
				fmt.Scan(&username)
				fmt.Print("Masukkan Password : ")
				fmt.Scan(&password)
				idx = searchIdxUsn(petugas, nPetugas, username)
				if cekLogin(petugas, nPetugas, username, password) {
					fmt.Println("--------------------------------------------------------------------------------------------")
					fmt.Println("*||*                      *LOGIN BERHASIL! MENUJU KE MENU PETUGAS*                      *||*")
					fmt.Println("--------------------------------------------------------------------------------------------")
					duration := time.Duration(3) * time.Second
					time.Sleep(duration)
					ClearScreen()
					petugasMenu(petugas, idx)
					statusLogin = true
				} else {
					fmt.Println("--------------------------------------------------------------------------------------------")
					fmt.Println("*||*             *Username atau Password Anda Salah! Silahkan input ulang!*             *||*")
					fmt.Println("--------------------------------------------------------------------------------------------")
					duration := time.Duration(3) * time.Second
					time.Sleep(duration)
					ClearScreen()
				}
			}
		}
	}
}

func petugasMenu(petugas tabPetugas, idx int) {
	var input, nMobil, nMotor, nData, totalPendapatan int
	var pMobil tabParkirMobil
	var pMotor tabParkirMotor
	var dataKendaraan tabDataKendaraan
	for input != 6 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                  *MENU PETUGAS*                                    *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("*Selamat Datang Petugas %s, silahkan pilih perintah berikut. Selamat bekerja!*\n", petugas[idx].nama)
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Tambah Data Kendaraan Parkir")
		fmt.Println("2. Edit Data Kendaraan Parkir")
		fmt.Println("3. Hapus Data Kendaraan Parkir")
		fmt.Println("4. Cari Data Kendaraan Parkir")
		fmt.Println("5. Cetak Data Kendaraan Parkir")
		fmt.Println("6. Log Out")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5/6) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			inputHari(&pMobil, &pMotor, &dataKendaraan, &nMobil, &nMotor, &nData, &totalPendapatan) // 7.00 - 22.00
		} else if input == 2 {
			editKendaraan(&pMobil, &pMotor, &dataKendaraan, &totalPendapatan, nMobil, nMotor, nData)
		} else if input == 3 {
			hapusDataKendaraan(&pMobil, &pMotor, &dataKendaraan, &nMobil, &nMotor, &nData, &totalPendapatan)
		} else if input == 4 {
			cariKendaraan(dataKendaraan, nData)
		} else if input == 5 {
			cetakKendaraan(pMobil, pMotor, dataKendaraan, nMobil, nMotor, nData, totalPendapatan)
		}
		ClearScreen()
	}
}

func inputHari(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	var input int
	var hari string
	for input != 2 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                       *MENU TAMBAH KENDARAAN PARKIR*                               *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Input Hari")
		fmt.Println("2. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			fmt.Print("Hari (Minggu,..,Jumat,Sabtu) : ")
			fmt.Scan(&hari)
			tambahKendaraan(&*pMobil, &*pMotor, &*dataKendaraan, &*nMobil, &*nMotor, &*nData, &*totalPendapatan, hari)
			ClearScreen()
		}
	}

}

func tambahKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int, hari string) {
	var input int
	for input != 3 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("1. Kendaraan Masuk Parkir")
		fmt.Println("2. Kendaraan Keluar Parkir")
		fmt.Println("3. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			masukParkir(&*pMobil, &*pMotor, &*dataKendaraan, &*nMobil, &*nMotor, &*nData, hari)
		} else if input == 2 {
			keluarParkir(&*pMobil, &*pMotor, &*dataKendaraan, &*nMobil, &*nMotor, &*nData, &*totalPendapatan)
		}
		ClearScreen()
	}
}

func masukParkir(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData *int, hari string) {
	/* IS : Array pMobil, pMotor, dataKendaraan dengan banyak elemen nMobil, nMotor, nData terdefinisi sembarang
	      FS : Data mobil dan motor (kecuali data jam keluar parkir dan harga parkir) ditambahkan ke dalam array pMobil, pMotor, dan dataKendaraan sebanyak nMobil, nMotor, nData
	           (Jam masuk kendaraan 07.00 - 22.59)
	   	   	(array dataKendaraan berfungsi untuk menampung seluruh data baik mobil maupun motor yang masuk dan keluar parkiran)
	*/
	var input, jamMasuk, menitMasuk int
	var nopol string
	for input != 3 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Jenis Kendaraan :")
		fmt.Println("1. Mobil")
		fmt.Println("2. Motor")
		fmt.Println("3. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			if *nMobil > MAXPARKIRMOBIL-1 {
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Println("*||*                       *PARKIRAN MOBIL SUDAH PENUH!*                                *||*")
				fmt.Println("--------------------------------------------------------------------------------------------")
				duration := time.Duration(3) * time.Second
				time.Sleep(duration)
				ClearScreen()
				*nMobil = MAXPARKIRMOBIL - 1
			} else {
				pMobil[*nMobil].tipe = "Mobil"
				pMobil[*nMobil].hari = hari
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol)
				for !validNoPolMobil(*pMobil, *nMobil, nopol) || !validNoPolMotor(*pMotor, *nMotor, nopol) {
					fmt.Println("Nomor Polisi tidak valid!")
					fmt.Print("Nomor Polisi :")
					fmt.Scan(&nopol)
				}
				pMobil[*nMobil].noPol = nopol
				fmt.Print("Vip (true/false): ")
				fmt.Scan(&pMobil[*nMobil].vip)
				fmt.Print("Valet (true/false) : ")
				fmt.Scan(&pMobil[*nMobil].valet)
				for {
					fmt.Print("Jam Masuk (7 00 - 22 59) : ")
					fmt.Scan(&jamMasuk, &menitMasuk)
					if jamMasuk >= 7 && jamMasuk <= 22 {
						break
					}
					fmt.Println("Jam masuk parkir tidak valid!")
				}
				pMobil[*nMobil].waktu.jamMasuk = jamMasuk
				pMobil[*nMobil].waktu.menitMasuk = menitMasuk // jam masuk min 07.00 maks 22.59
				dataKendaraan[*nData] = pMobil[*nMobil]
				*nMobil++
				*nData++
				ClearScreen()
			}
		} else if input == 2 {
			if *nMotor > MAXPARKIRMOTOR-1 {
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Println("*||*                       *PARKIRAN MOTOR SUDAH PENUH!*                                *||*")
				fmt.Println("--------------------------------------------------------------------------------------------")
				duration := time.Duration(3) * time.Second
				time.Sleep(duration)
				ClearScreen()
				*nMotor = MAXPARKIRMOTOR - 1
			} else {
				pMotor[*nMotor].tipe = "Motor"
				pMotor[*nMotor].hari = hari
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol)
				for !validNoPolMobil(*pMobil, *nMobil, nopol) || !validNoPolMotor(*pMotor, *nMotor, nopol) {
					fmt.Println("Nomor Polisi tidak valid!")
					fmt.Print("Nomor Polisi : ")
					fmt.Scan(&nopol)
				}
				pMotor[*nMotor].noPol = nopol
				fmt.Print("Vip (true/false) : ")
				fmt.Scan(&pMotor[*nMotor].vip)
				fmt.Print("Valet (true/false) : ")
				fmt.Scan(&pMotor[*nMotor].valet)
				//fmt.Scanf("%02d %02d", &jamMasuk, &menitMasuk)
				for {
					fmt.Print("Jam Masuk (7 00 - 22 59) : ")
					fmt.Scan(&jamMasuk, &menitMasuk)
					if jamMasuk >= 7 && jamMasuk <= 22 {
						break
					}
					fmt.Println("Jam masuk parkir tidak valid!")
				}
				pMotor[*nMotor].waktu.jamMasuk = jamMasuk
				pMotor[*nMotor].waktu.menitMasuk = menitMasuk
				dataKendaraan[*nData] = pMotor[*nMotor]
				*nMotor++
				*nData++
				ClearScreen()
			}
		}
	}
}

func keluarParkir(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	/* IS : Array pMobil, pMotor, dataKendaraan dengan banyak elemen nMobil, nMotor, nData terdefinisi (kecuali data jam keluar parkir dan harga parkir)
	      FS : Data jam keluar parkir dan harga parkir mobil atau motor ditambahkan ke dalam array pMobil, pMotor, dan dataKendaraan
	           berdasarkan nomor polisi jika ditemukan. Jika tidak ditemukan, maka tercetak "Nomor Polisi tidak terdapat di dalam parkiran!"
	   		(Jam keluar parkir 07.00 - 23.00)
	   	   	(array dataKendaraan berfungsi untuk menampung seluruh data baik mobil maupun motor yang masuk dan keluar parkiran)
	*/
	var nopol string
	var idx, idxData, jamKeluar, menitKeluar int
	fmt.Print("Masukkan Nomor Polisi : ")
	fmt.Scan(&nopol)
	if validNoPolMobil(*pMobil, *nMobil, nopol) && validNoPolMotor(*pMotor, *nMotor, nopol) {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*              *NOMOR POLISI TIDAK TERDAPAT DI DALAM PARKIRAN!*                      *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		idxData = searchIdxNoPol(*dataKendaraan, *nData, nopol)
		if !validNoPolMobil(*pMobil, *nMobil, nopol) {
			idx = searchIdxNoPolMobil(*pMobil, *nMobil, nopol)
			for {
				fmt.Print("Jam Keluar (7 01 - 23 00) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar > dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar > dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar == 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			pMobil[idx].waktu.jamKeluar = jamKeluar
			pMobil[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
			dataKendaraan[idxData] = pMobil[idx]
			hapusMobil(&*pMobil, &*nMobil, idx)
		} else if !validNoPolMotor(*pMotor, *nMotor, nopol) {
			idx = searchIdxNoPolMotor(*pMotor, *nMotor, nopol)
			//fmt.Scanf("%02d %02d", &jamKeluar, &menitKeluar)
			for {
				fmt.Print("Jam Keluar (7 01 - 23 00) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar > dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar > dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar == 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			pMotor[idx].waktu.jamKeluar = jamKeluar
			pMotor[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
			dataKendaraan[idxData] = pMotor[idx]
			hapusMotor(&*pMotor, &*nMotor, idx)
		}
		hitungHarga(&*dataKendaraan, idxData)
		*totalPendapatan += dataKendaraan[idxData].hargaParkir
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("Hari : %s\n", dataKendaraan[idxData].hari)
		fmt.Printf("Tipe Kendaraan : %s\n", dataKendaraan[idxData].tipe)
		fmt.Printf("Nomor Polisi : %s\n", dataKendaraan[idxData].noPol)
		fmt.Printf("Vip : %t\n", dataKendaraan[idxData].vip)
		fmt.Printf("Valet : %t\n", dataKendaraan[idxData].valet)
		fmt.Printf("Jam Masuk : %02d:%02d\n", dataKendaraan[idxData].waktu.jamMasuk, dataKendaraan[idxData].waktu.menitMasuk)
		fmt.Printf("Jam Keluar : %02d:%02d\n", dataKendaraan[idxData].waktu.jamKeluar, dataKendaraan[idxData].waktu.menitKeluar)
		fmt.Printf("Harga Parkir : Rp.%d\n", dataKendaraan[idxData].hargaParkir)
		fmt.Printf("Total Pendapatan : Rp.%d\n", *totalPendapatan)
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Press 'Enter' to continue...")
		fmt.Scanln()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		ClearScreen()
	}
}

func editKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, totalPendapatan *int, nMobil, nMotor, nData int) {
	var input int
	for input != 8 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                             *MENU EDIT DATA KENDARAAN*                             *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Kendaraan yang ingin diedit :")
		fmt.Println("1. Hari")
		fmt.Println("2. Tipe")
		fmt.Println("3. Nomor Polisi")
		fmt.Println("4. VIP")
		fmt.Println("5. Valet")
		fmt.Println("6. Jam Masuk")
		fmt.Println("7. Jam Keluar")
		fmt.Println("8. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5/6/7/8) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 || input == 7 {
			editDataKendaraan(&*pMobil, &*pMotor, &*dataKendaraan, &*totalPendapatan, nMobil, nMotor, nData, input)
			ClearScreen()
		}
	}
}

func editDataKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, totalPendapatan *int, nMobil, nMotor, nData, input int) {
	/* IS : Array pMobil, pMotor, dataKendaraan dengan banyak elemen nMobil, nMotor, nData terdefinisi
	   FS : Data kendaraan berhasil diubah berdasarkan data kendaraan yang dipilih
	   		(contoh : ingin merubah jam masuk maka jam masuk kendaraan dengan nopol tertentu akan berubah)
	*/
	var idxData, idxMobil, idxMotor, jamMasuk, jamKeluar, menitMasuk, menitKeluar int
	var nopol, nopolEdit string
	fmt.Print("Masukkan Nomor Polisi kendaraan yang akan diedit : ")
	fmt.Scan(&nopolEdit)
	idxData = searchIdxNoPol(*dataKendaraan, nData, nopolEdit)
	if idxData == -1 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*              *NOMOR POLISI TIDAK TERDAPAT DI DALAM PARKIRAN!*                      *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		*totalPendapatan -= dataKendaraan[idxData].hargaParkir
		if input == 1 {
			fmt.Print("Hari : ")
			fmt.Scan(&dataKendaraan[idxData].hari)
			fmt.Println("Hari berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 2 {
			fmt.Print("Tipe (Mobil/Motor) : ")
			fmt.Scan(&dataKendaraan[idxData].tipe)
			fmt.Println("Tipe kendaraan berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 3 {
			fmt.Print("Nomor Polisi : ")
			fmt.Scan(&nopol)
			for !validNoPol(*dataKendaraan, nData, nopol) {
				fmt.Println("Nomor Polisi tidak valid!")
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol)
			}
			dataKendaraan[idxData].noPol = nopol
			fmt.Println("Nomor Polisi berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 4 {
			fmt.Print("Vip : ")
			fmt.Scan(&dataKendaraan[idxData].vip)
			fmt.Println("Vip berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 5 {
			fmt.Print("Valet : ")
			fmt.Scan(&dataKendaraan[idxData].valet)
			fmt.Println("Valet berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 6 {
			for {
				fmt.Print("Jam Masuk (7 00 - 22 59) : ")
				fmt.Scan(&jamMasuk, &menitMasuk)
				if jamMasuk >= 7 && jamMasuk <= 22 {
					break
				}
				fmt.Println("Jam masuk parkir tidak valid!")
			}
			dataKendaraan[idxData].waktu.jamMasuk = jamMasuk
			dataKendaraan[idxData].waktu.menitMasuk = menitMasuk
			fmt.Println("Jam Masuk berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		} else if input == 7 {
			for {
				fmt.Print("Jam Keluar (7 01 - 23 00) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar >= dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar >= dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar > 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			dataKendaraan[idxData].waktu.jamKeluar = jamKeluar
			dataKendaraan[idxData].waktu.menitKeluar = menitKeluar
			fmt.Println("Jam Keluar berhasil diedit :)")
			duration := time.Duration(3) * time.Second
			time.Sleep(duration)
			ClearScreen()
		}
		if !validNoPolMobil(*pMobil, nMobil, nopolEdit) {
			idxMobil = searchIdxNoPolMobil(*pMobil, nData, nopolEdit)
			pMobil[idxMobil] = dataKendaraan[idxData]
		} else if !validNoPolMotor(*pMotor, nMotor, nopolEdit) {
			idxMotor = searchIdxNoPolMotor(*pMotor, nData, nopolEdit)
			pMotor[idxMotor] = dataKendaraan[idxData]
		}
		hitungHarga(&*dataKendaraan, idxData)
		*totalPendapatan += dataKendaraan[idxData].hargaParkir

	}
}

func hapusDataKendaraan(pMobil *tabParkirMobil, pMotor *tabParkirMotor, dataKendaraan *tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan *int) {
	/* IS : Array data kendaraan dengan banyak elemen nData terdefinisi
	   FS : Data kendaraan dengan nopol tertentu terhapus
	*/
	var nopol string
	var idxData, idxMobil, idxMotor, i int
	fmt.Println("============================================================================================")
	fmt.Println("*||*                             *MENU HAPUS DATA KENDARAAN*                            *||*")
	fmt.Println("============================================================================================")
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Masukkan Nomor Polisi dari kendaraan yang ingin dihapus : ")
	fmt.Scan(&nopol)
	idxData = searchIdxNoPol(*dataKendaraan, *nData, nopol)
	if idxData == -1 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*            *NOMOR POLISI TIDAK TERDAPAT DI DALAM DATA PARKIRAN!*                    *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		i = idxData
		*totalPendapatan -= dataKendaraan[i].hargaParkir
		for i < *nData-1 {
			dataKendaraan[i] = dataKendaraan[i+1]
			i++
		}
		*nData--
		if !validNoPolMobil(*pMobil, *nMobil, nopol) {
			idxMobil = searchIdxNoPolMobil(*pMobil, *nMobil, nopol)
			i = idxMobil
			for i < *nMobil-1 {
				pMobil[i] = pMobil[i+1]
				i++
			}
			*nMobil--
		} else if !validNoPolMotor(*pMotor, *nMotor, nopol) {
			idxMotor = searchIdxNoPolMotor(*pMotor, *nMotor, nopol)
			i = idxMotor
			for i < *nMotor-1 {
				pMotor[i] = pMotor[i+1]
				i++
			}
			*nMotor--
		}
		fmt.Printf("Semua data kendaraan dengan Nomor Polisi %s berhasil dihapus :)\n", nopol)
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	}
}

func cariKendaraan(dataKendaraan tabDataKendaraan, nData int) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinisi
	   FS : Data kendaraan dengan nopol tertentu tercetak, jika nopol tidak ditemukan maka data tidak tercetak
	*/
	var nopol string
	var idxData int
	fmt.Print("Masukkan Nomor Polisi : ")
	fmt.Scan(&nopol)
	if validNoPol(dataKendaraan, nData, nopol) {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*||*            *NOMOR POLISI TIDAK TERDAPAT DI DALAM DATA PARKIRAN!*                    *||*")
		fmt.Println("--------------------------------------------------------------------------------------------")
		duration := time.Duration(3) * time.Second
		time.Sleep(duration)
		ClearScreen()
	} else {
		idxData = searchIdxNoPol(dataKendaraan, nData, nopol)
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Printf("Hari : %s\n", dataKendaraan[idxData].hari)
		fmt.Printf("Tipe Kendaraan : %s\n", dataKendaraan[idxData].tipe)
		fmt.Printf("Nomor Polisi : %s\n", dataKendaraan[idxData].noPol)
		fmt.Printf("Vip : %t\n", dataKendaraan[idxData].vip)
		fmt.Printf("Valet : %t\n", dataKendaraan[idxData].valet)
		fmt.Printf("Jam Masuk : %02d:%02d\n", dataKendaraan[idxData].waktu.jamMasuk, dataKendaraan[idxData].waktu.menitMasuk)
		fmt.Printf("Jam Keluar : %02d:%02d\n", dataKendaraan[idxData].waktu.jamKeluar, dataKendaraan[idxData].waktu.menitKeluar)
		fmt.Printf("Harga Parkir : Rp.%d\n", dataKendaraan[idxData].hargaParkir)
		if dataKendaraan[idxData].hargaParkir == 0 {
			fmt.Println("Status : Masih dalam parkiran")
		} else {
			fmt.Println("Status : Sudah keluar parkiran")
		}
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Press 'Enter' to continue...")
		fmt.Scanln()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		ClearScreen()
	}
}

func cetakKendaraan(pMobil tabParkirMobil, pMotor tabParkirMotor, dataKendaraan tabDataKendaraan, nMobil, nMotor, nData, totalPendapatan int) {
	/* IS : Array pMobil, pMotor, dan dataKendaraan dengan banyak elemen nMobil, nMotor, dan nData terdefinisi
	   FS : Data mobil atau motor tercetak di layar berdasarkan kriteria tertentu
	*/
	var input int
	for input != 4 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                            *MENU CETAK DATA KENDARAAN*                             *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih Data Kendaraan yang ingin dicetak :")
		fmt.Println("1. Data di Parkiran Mobil")
		fmt.Println("2. Data di Parkiran Motor")
		fmt.Println("3. Data Keseluruhan")
		fmt.Println("4. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			cetakParkirMobil(pMobil, nMobil)
		} else if input == 2 {
			cetakParkirMotor(pMotor, nMotor)
		} else if input == 3 {
			cetakDataKeseluruhan(dataKendaraan, nData, totalPendapatan)
		}
		ClearScreen()
	}
}

func cetakDataKeseluruhan(dataKendaraan tabDataKendaraan, nData, totalPendapatan int) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinisi
	   FS : Data kendaraan dan totalPendapatan tercetak di layar berdasarkan kriteria tertentu
	*/
	var input int
	var tipe, hari string
	for input != 5 {
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("Pilih urutan data yang ingin dicetak :")
		fmt.Println("1. Urut berdasarkan harga tertinggi")
		fmt.Println("2. Urut berdasarkan harga terendah")
		fmt.Println("3. Urut berdasarkan tipe kendaraan")
		fmt.Println("4. Urut berdasarkan hari")
		fmt.Println("5. Kembali")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4/5) : ")
		fmt.Scan(&input)
		ClearScreen()
		if input == 1 {
			urutBerdasarkanHargaDescending(&dataKendaraan, nData)
			cetakParkirKeseluruhan(dataKendaraan, nData, totalPendapatan)
		} else if input == 2 {
			urutBerdasarkanHargaAscending(&dataKendaraan, nData)
			cetakParkirKeseluruhan(dataKendaraan, nData, totalPendapatan)
		} else if input == 3 {
			fmt.Print("Masukkan Tipe Kendaraan : ")
			fmt.Scan(&tipe)
			urutberdasarkanTipe(&dataKendaraan, nData, tipe)
			cetakParkirKeseluruhan(dataKendaraan, nData, totalPendapatan)
		} else if input == 4 {
			fmt.Print("Masukkan Hari : ")
			fmt.Scan(&hari)
			urutBerdasarkanHari(&dataKendaraan, nData, hari)
			cetakParkirKeseluruhan(dataKendaraan, nData, totalPendapatan)
		}
		ClearScreen()
	}
}

func searchID(petugas tabPetugas, nPetugas int, id int) int {
	/*mengembalikan indeks dari array petugas dengan ID tertentu, jika tidak ditemukan akan mengembalikan -1
	  (menggunakan algoritma binary search) */
	var idx, mid, right, left int
	left = 0
	right = nPetugas - 1
	mid = (left + right) / 2
	idx = -1
	for left <= right && idx == -1 {
		if petugas[mid].id == id {
			idx = mid
		} else if id > petugas[mid].id {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return idx
}

func validId(petugas tabPetugas, nPetugas int, id int) bool {
	/*mengembalikan nilai false apabila ditemukan ID petugas yang sama, jika tidak ditemukan akan mengembalikan true
	  (menggunakan algoritma binary search) */
	var mid, right, left int
	var ketemu bool = false
	left = 0
	right = nPetugas - 1
	mid = (left + right) / 2
	for left <= right && !ketemu {
		ketemu = petugas[mid].id == id
		if id > petugas[mid].id {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return !ketemu
}

func searchIdxUsn(petugas tabPetugas, nPetugas int, usn string) int {
	/*mengembalikan indeks dari array petugas dengan username tertentu, jika tidak ditemukan akan mengembalikan -1
	  (menggunakan algoritma sequential search) */
	var idx, i int
	idx = -1
	for i < nPetugas && idx == -1 {
		if petugas[i].username == usn {
			idx = i
		}
		i++
	}
	return idx
}

func validUsn(petugas tabPetugas, nPetugas int, usn string) bool {
	/*mengembalikan nilai false apabila ditemukan username petugas yang sama, jika tidak ditemukan akan mengembalikan true
	  (menggunakan algoritma sequential search) */
	var i int
	var ketemu bool = false
	for i < nPetugas && !ketemu {
		ketemu = petugas[i].username == usn
		i++
	}
	return !ketemu
}

func cekLogin(petugas tabPetugas, nPetugas int, username, password string) bool {
	/*mengembalikan nilai true apabila username dan password yang diinputkan terdapat di dalam array petugas,
	  jika tidak ditemukan akan mengembalikan false (menggunakan allgoritma sequential search)*/
	var i int
	var ketemu bool = false
	for i < nPetugas && !ketemu {
		ketemu = petugas[i].username == username && petugas[i].password == password
		i++
	}
	return ketemu
}

func searchIdxNoPolMobil(pMobil tabParkirMobil, nMobil int, nopol string) int {
	/*mengembalikan indeks dari array pMobil dengan nomor polisi tertentu, jika tidak ditemukan akan mengembalikan -1
	  (menggunakan algoritma sequential search) */
	var i, idx int
	idx = -1
	for i < nMobil && idx == -1 {
		if pMobil[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func validNoPolMobil(pMobil tabParkirMobil, nMobil int, nopol string) bool {
	/*mengembalikan nilai false apabila ditemukan nomor polisi yang sama dari array pMobil, jika tidak ditemukan akan mengembalikan true
	  (menggunakan algoritma sequential search) */
	var i int
	var ketemu bool = false
	for i < nMobil && !ketemu {
		ketemu = pMobil[i].noPol == nopol
		i++
	}
	return !ketemu
}

func searchIdxNoPolMotor(pMotor tabParkirMotor, nMotor int, nopol string) int {
	/*mengembalikan indeks dari array pMotor dengan nomor polisi tertentu, jika tidak ditemukan akan mengembalikan -1
	  (menggunakan algoritma sequential search) */
	var i, idx int
	idx = -1
	for i < nMotor && idx == -1 {
		if pMotor[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func validNoPolMotor(pMotor tabParkirMotor, nMotor int, nopol string) bool {
	/*mengembalikan nilai false apabila ditemukan nomor polisi yang sama dari array pMotor, jika tidak ditemukan akan mengembalikan true
	  (menggunakan algoritma sequential search) */
	var i int
	var ketemu bool = false
	for i < nMotor && !ketemu {
		ketemu = pMotor[i].noPol == nopol
		i++
	}
	return !ketemu
}

func searchIdxNoPol(dataKendaraan tabDataKendaraan, nData int, nopol string) int {
	/*mengembalikan indeks dari array dataKendaraan dengan nomor polisi tertentu, jika tidak ditemukan akan mengembalikan -1
	  (menggunakan algoritma sequential search) */
	var i, idx int
	idx = -1
	for i < nData && idx == -1 {
		if dataKendaraan[i].noPol == nopol {
			idx = i
		}
		i++
	}
	return idx
}

func validNoPol(dataKendaraan tabDataKendaraan, nData int, nopol string) bool {
	/*mengembalikan nilai false apabila ditemukan nomor polisi yang sama dari array dataKendaraan, jika tidak ditemukan akan mengembalikan true
	  (menggunakan algoritma sequential search) */
	var i int
	var ketemu bool = false
	for i < nData && !ketemu {
		ketemu = dataKendaraan[i].noPol == nopol
		i++
	}
	return !ketemu
}

func hitungHarga(dataKendaraan *tabDataKendaraan, idxData int) {
	/* IS : Array dataKendaraan dengan indeks idxData terdefinisi (kecuali field hargaParkir)
	      FS : hargaParkir terdefinisi dengann ketentuan :
	           - Jika kendaraan bertipe Mobil akan dikenakan harga Rp.10.000 untuk jam pertama
	   		  dan Rp.5.000 untuk jam berikutnya. Jika parkir di tempat vip akan dikenakan tambahan
	   		  harga Rp.20.000. Jika menggunakan jasa valet akan dikenakan tambahan harga Rp.50.000
	   		- Jika kendaraan bertipe Motor akan dikenakan harga Rp.5.000 untuk jam pertama
	   		  dan Rp.2.000 untuk jam berikutnya. Jika parkir di tempat vip akan dikenakan tambahan
	   		  harga Rp.10.000. Jika menggunakan jasa valet akan dikenakan tambahan harga Rp.20.000
	   		- Jika kendaraan parkir di hari Jumat akan mendapatkan diskon sebesar 20%
	*/
	var harga, totalMenit, totalJam int
	totalMenit = (dataKendaraan[idxData].waktu.jamKeluar*60 + dataKendaraan[idxData].waktu.menitKeluar) - (dataKendaraan[idxData].waktu.jamMasuk*60 + dataKendaraan[idxData].waktu.menitMasuk)
	totalJam = totalMenit / 60
	if totalMenit%60 != 0 {
		totalJam++
	}
	if dataKendaraan[idxData].tipe == "Mobil" {
		harga = 10000 + (totalJam-1)*5000
		if dataKendaraan[idxData].vip {
			harga += 20000
		}
		if dataKendaraan[idxData].valet {
			harga += 50000
		}
	} else if dataKendaraan[idxData].tipe == "Motor" {
		harga = 5000 + (totalJam-1)*2000
		if dataKendaraan[idxData].vip {
			harga += 10000
		}
		if dataKendaraan[idxData].valet {
			harga += 20000
		}
	}
	if dataKendaraan[idxData].hari == "Jumat" {
		harga = harga - (harga * 20 / 100)
	}
	if dataKendaraan[idxData].waktu.jamKeluar == 0 {
		harga = 0
	}
	dataKendaraan[idxData].hargaParkir = harga
}

func hapusMobil(pMobil *tabParkirMobil, nMobil *int, idx int) {
	/* IS : Array pMobil dengan banyak elemen nMobil terdefinsi
	   FS : Elemen dari array pMobil dengan indeks tertentu terhapus
	*/
	var i int
	i = idx
	for i < *nMobil-1 {
		pMobil[i] = pMobil[i+1]
		i++
	}
	*nMobil--
}

func hapusMotor(pMotor *tabParkirMotor, nMotor *int, idx int) {
	/* IS : Array pMotor dengan banyak elemen nMotor terdefinsi
	   FS : Elemen dari array pMotor dengan indeks tertentu terhapus
	*/
	var i int
	i = idx
	for i < *nMotor-1 {
		pMotor[i] = pMotor[i+1]
		i++
	}
	*nMotor--
}

func urutID(petugas *tabPetugas, nPetugas int) {
	/* IS : Array petugas dengan banyak elemen nPetugas terdefinsi
	   FS : Array petugas dengan banyak elemen nPetugas terdefinsi secara urut menaik (ascending) berdasarkan ID petugas
	        (menggunakan algoritma selection sort)
	*/
	var i, idx, pass int
	var temp tPetugas
	for pass = 1; pass <= nPetugas-1; pass++ {
		idx = pass - 1
		for i = pass; i <= nPetugas-1; i++ {
			if petugas[i].id < petugas[idx].id {
				idx = i
			}
		}
		temp = petugas[idx]
		petugas[idx] = petugas[pass-1]
		petugas[pass-1] = temp
	}
}

func urutBerdasarkanHargaDescending(dataKendaraan *tabDataKendaraan, nData int) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinsi
	   FS : Array dataKendaraan dengan banyak elemen nData terdefinsi secara urut menurun (descending) berdasarkan harga parkir
	        (menggunakan algortima insertion sort)
	*/
	var i, pass int
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		temp = dataKendaraan[pass]
		i = pass
		for i > 0 && dataKendaraan[i-1].hargaParkir < temp.hargaParkir {
			dataKendaraan[i] = dataKendaraan[i-1]
			i--
		}
		dataKendaraan[i] = temp
	}
}

func urutBerdasarkanHargaAscending(dataKendaraan *tabDataKendaraan, nData int) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinsi
	   FS : Array dataKendaraan dengan banyak elemen nData terdefinsi secara urut menaik (ascending) berdasarkan harga parkir
	        (menggunakan algoritma insertion sort)
	*/
	var i, pass int
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		temp = dataKendaraan[pass]
		i = pass
		for i > 0 && dataKendaraan[i-1].hargaParkir > temp.hargaParkir {
			dataKendaraan[i] = dataKendaraan[i-1]
			i--
		}
		dataKendaraan[i] = temp
	}
}

func urutBerdasarkanHari(dataKendaraan *tabDataKendaraan, nData int, hari string) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinsi
	      FS : Array dataKendaraan dengan banyak elemen nData terdefinsi secara urut berdasarkan hari yang diinput
	           (jika input hari Minggu, maka akan terurut mulai dari hari Minggu. Jika sudah tidak ada hari Minggu,
	   		 maka setelah hari Minggu terakhir akan tetap teracak)
	   		(menggunakan algoritma selection sort)
	*/
	var i, idx, pass int
	var ketemu bool
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		ketemu = false
		idx = pass - 1
		i = idx
		for i <= nData-1 && !ketemu {
			ketemu = dataKendaraan[i].hari == hari
			if ketemu {
				idx = i
			}
			i++
		}
		temp = dataKendaraan[idx]
		dataKendaraan[idx] = dataKendaraan[pass-1]
		dataKendaraan[pass-1] = temp
	}
}

func urutberdasarkanTipe(dataKendaraan *tabDataKendaraan, nData int, tipe string) {
	/* IS : Array dataKendaraan dengan banyak elemen nData terdefinsi
	   FS : Array dataKendaraan dengan banyak elemen nData terdefinsi secara berdasarkan tipe kendaraan (Mobil/Motor)
	        (menggunakan algoritma selection sort)
	*/
	var i, idx, pass int
	var ketemu bool
	var temp tKendaraan
	for pass = 1; pass <= nData-1; pass++ {
		ketemu = false
		idx = pass - 1
		i = idx
		for i <= nData-1 && !ketemu {
			ketemu = dataKendaraan[i].tipe == tipe
			if ketemu {
				idx = i
			}
			i++
		}
		temp = dataKendaraan[idx]
		dataKendaraan[idx] = dataKendaraan[pass-1]
		dataKendaraan[pass-1] = temp
	}
}

func cetakParkirMobil(pMobil tabParkirMobil, nMobil int) {
	/* IS : Array pMobil dengan banyak elemen nMobil terdefinsi
	   FS : Array pMobil dengan banyak elemen nMobil tercetak di layar
	*/
	var i int
	fmt.Println("Data Parkiran Mobil : ")
	fmt.Printf("%-6s %-7s %-10s %-6s %-6s %-10s %-10s %-10s\n", "Tipe", "Hari", "NoPol", "VIP", "Valet", "Jam Masuk", "Jam Keluar", "Harga")
	for i = 0; i < nMobil; i++ {
		fmt.Printf("%-6s %-7s %-10s %-6t %-6t %02d:%02d %4s %02d:%02d %4s Rp.%-10d\n", pMobil[i].tipe, pMobil[i].hari, pMobil[i].noPol, pMobil[i].vip, pMobil[i].valet, pMobil[i].waktu.jamMasuk, pMobil[i].waktu.menitMasuk, "", pMobil[i].waktu.jamKeluar, pMobil[i].waktu.menitKeluar, "", pMobil[i].hargaParkir)
	}
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func cetakParkirMotor(pMotor tabParkirMotor, nMotor int) {
	/* IS : Array pMotor dengan banyak elemen nMotor terdefinsi
	   FS : Array pMotor dengan banyak elemen nMotor tercetak di layar
	*/
	var i int
	fmt.Println("Data Parkiran Motor : ")
	fmt.Printf("%-6s %-7s %-10s %-6s %-6s %-10s %-10s %-10s\n", "Tipe", "Hari", "NoPol", "VIP", "Valet", "Jam Masuk", "Jam Keluar", "Harga")
	for i = 0; i < nMotor; i++ {
		fmt.Printf("%-6s %-7s %-10s %-6t %-6t %02d:%02d %4s %02d:%02d %4s Rp.%-10d\n", pMotor[i].tipe, pMotor[i].hari, pMotor[i].noPol, pMotor[i].vip, pMotor[i].valet, pMotor[i].waktu.jamMasuk, pMotor[i].waktu.menitMasuk, "", pMotor[i].waktu.jamKeluar, pMotor[i].waktu.menitKeluar, "", pMotor[i].hargaParkir)
	}
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func cetakParkirKeseluruhan(dataKendaraan tabDataKendaraan, nData, totalPendapatan int) {
	/* IS : Array dataKendaraan dengan banyak elemen nMotor terdefinsi
	   FS : Array dataKendaraan dengan banyak elemen nMotor tercetak di layar berdasarkan urutan tertentu
	        dan tercetak total pendapatan di akhir
	*/
	var i int
	fmt.Println("Data Parkiran Keseluruhan : ")
	fmt.Printf("%-6s %-7s %-10s %-6s %-6s %-10s %-10s %-10s\n", "Tipe", "Hari", "NoPol", "VIP", "Valet", "Jam Masuk", "Jam Keluar", "Harga")
	for i = 0; i < nData; i++ {
		fmt.Printf("%-6s %-7s %-10s %-6t %-6t %02d:%02d %4s %02d:%02d %4s Rp.%-10d\n", dataKendaraan[i].tipe, dataKendaraan[i].hari, dataKendaraan[i].noPol, dataKendaraan[i].vip, dataKendaraan[i].valet, dataKendaraan[i].waktu.jamMasuk, dataKendaraan[i].waktu.menitMasuk, "", dataKendaraan[i].waktu.jamKeluar, dataKendaraan[i].waktu.menitKeluar, "", dataKendaraan[i].hargaParkir)
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("Total Pendapatan : Rp.%d\n", totalPendapatan)
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Print("Press 'Enter' to continue...")
	fmt.Scanln()
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	ClearScreen()
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
