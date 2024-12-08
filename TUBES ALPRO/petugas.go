package main

import "fmt"

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
		if input == 1 {
			fmt.Print("Hari (Minggu,..,Jumat,Sabtu) : ")
			fmt.Scan(&hari)
			tambahKendaraan(pMobil, pMotor, dataKendaraan, nMobil, nMotor, nData, totalPendapatan, hari)
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
		if input == 1 {
			masukParkir(pMobil, pMotor, dataKendaraan, nMobil, nMotor, nData, hari)
		} else if input == 2 {
			keluarParkir(pMobil, pMotor, dataKendaraan, nMobil, nMotor, nData, totalPendapatan)
		}
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
		if input == 1 {
			if *nMobil > MAXPARKIRMOBIL-1 {
				fmt.Println("Parkiran mobil sudah penuh!")
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
				//fmt.Print("Jam Masuk : ")
				for {
					fmt.Print("Jam Masuk (7 05) : ")
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
			}
		} else if input == 2 {
			if *nMotor > MAXPARKIRMOTOR-1 {
				fmt.Println("Parkiran motor sudah penuh!")
				*nMotor = MAXPARKIRMOTOR - 1
			} else {
				pMotor[*nMotor].tipe = "Motor"
				pMobil[*nMobil].hari = hari
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
					fmt.Print("Jam Masuk (7 05) : ")
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
		fmt.Println("Nomor Polisi tidak terdapat di dalam parkiran!")
	} else {
		idxData = searchIdxNoPol(*dataKendaraan, *nData, nopol)
		if !validNoPolMobil(*pMobil, *nMobil, nopol) {
			idx = searchIdxNoPolMobil(*pMobil, *nMobil, nopol)
			//fmt.Scanf("%02d %02d", &jamKeluar, &menitKeluar)
			//fmt.Print("Jam Keluar : ")
			//fmt.Scan(&jamKeluar, &menitKeluar)
			for {
				fmt.Print("Jam Keluar (07 05) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar > dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar > dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar == 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			pMobil[idx].waktu.jamKeluar = jamKeluar
			pMobil[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
			dataKendaraan[idxData] = pMobil[idx]
			hapusMobil(pMobil, nMobil, idx)
		} else if !validNoPolMotor(*pMotor, *nMotor, nopol) {
			idx = searchIdxNoPolMotor(*pMotor, *nMotor, nopol)
			//fmt.Scanf("%02d %02d", &jamKeluar, &menitKeluar)
			for {
				fmt.Print("Jam Keluar (07 05) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar > dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar > dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar == 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			pMotor[idx].waktu.jamKeluar = jamKeluar
			pMotor[idx].waktu.menitKeluar = menitKeluar //jam keluar maks 23.00
			dataKendaraan[idxData] = pMotor[idx]
			hapusMotor(pMotor, nMotor, idx)
		}
		hitungHarga(dataKendaraan, idxData)
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
		if input == 1 || input == 2 || input == 3 || input == 4 || input == 5 || input == 6 || input == 7 {
			editDataKendaraan(pMobil, pMotor, dataKendaraan, totalPendapatan, nMobil, nMotor, nData, input)
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
		fmt.Println("Nomor Polisi tersebut tidak terdapat di dalam data kendaraan!")
	} else {
		*totalPendapatan -= dataKendaraan[idxData].hargaParkir
		if input == 1 {
			fmt.Print("Hari : ")
			fmt.Scan(&dataKendaraan[idxData].hari)
			fmt.Println("Hari berhasil diedit :)")
		} else if input == 2 {
			fmt.Print("Tipe (Mobil/Motor) : ")
			fmt.Scan(&dataKendaraan[idxData].tipe)
			fmt.Println("Tipe kendaraan berhasil diedit :)")
		} else if input == 3 {
			fmt.Println("Nomor Polisi : ")
			fmt.Scan(&nopol)
			for !validNoPol(*dataKendaraan, nData, nopol) {
				fmt.Println("Nomor Polisi tidak valid!")
				fmt.Print("Nomor Polisi : ")
				fmt.Scan(&nopol)
				fmt.Println("Nomor Polisi berhasil diedit :)")
			}
			dataKendaraan[idxData].noPol = nopol
		} else if input == 4 {
			fmt.Print("Vip : ")
			fmt.Scan(&dataKendaraan[idxData].vip)
			fmt.Println("Vip berhasil diedit :)")
		} else if input == 5 {
			fmt.Print("Valet : ")
			fmt.Scan(&dataKendaraan[idxData].valet)
			fmt.Println("Valet berhasil diedit :)")
		} else if input == 6 {
			for {
				fmt.Print("Jam Masuk (7 05) : ")
				fmt.Scan(&jamMasuk, &menitMasuk)
				if jamMasuk >= 7 && jamMasuk <= 22 {
					break
				}
				fmt.Println("Jam masuk parkir tidak valid!")
			}
			dataKendaraan[idxData].waktu.jamMasuk = jamMasuk
			dataKendaraan[idxData].waktu.menitMasuk = menitMasuk
			fmt.Println("Jam Masuk berhasil diedit :)")
		} else if input == 7 {
			for {
				fmt.Print("Jam Keluar (7 05) : ")
				fmt.Scan(&jamKeluar, &menitKeluar)
				if (jamKeluar >= dataKendaraan[idxData].waktu.jamMasuk && jamKeluar >= 7 && jamKeluar <= 23) || (jamKeluar == dataKendaraan[idxData].waktu.jamMasuk && menitKeluar >= dataKendaraan[idxData].waktu.menitMasuk) || (jamKeluar == 23 && menitKeluar > 0) {
					break
				}
				fmt.Println("Jam keluar parkir tidak valid!")
			}
			dataKendaraan[idxData].waktu.jamKeluar = jamKeluar
			dataKendaraan[idxData].waktu.menitKeluar = menitKeluar
			fmt.Println("Jam Keluar berhasil diedit :)")
		}
		if !validNoPolMobil(*pMobil, nMobil, nopolEdit) {
			idxMobil = searchIdxNoPolMobil(*pMobil, nData, nopolEdit)
			pMobil[idxMobil] = dataKendaraan[idxData]
		} else if !validNoPolMotor(*pMotor, nMotor, nopolEdit) {
			idxMotor = searchIdxNoPolMotor(*pMotor, nData, nopolEdit)
			pMotor[idxMotor] = dataKendaraan[idxData]
		}
		hitungHarga(dataKendaraan, idxData)
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
		fmt.Println("Nomor Polisi tersebut tidak terdapat di dalam data kendaraan yang parkir!")
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
		fmt.Println("Nomor Polisi tidak terdapat di dalam data parkiran!")
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
		if input == 1 {
			cetakParkirMobil(pMobil, nMobil)
		} else if input == 2 {
			cetakParkirMotor(pMotor, nMotor)
		} else if input == 3 {
			cetakDataKeseluruhan(dataKendaraan, nData, totalPendapatan)
		}
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
	}
}
