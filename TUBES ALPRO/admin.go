package main
import "fmt"
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
		if input == 1 {
			tambahPetugas(petugas, nPetugas)
		} else if input == 2 {
			editPetugas(petugas, *nPetugas)
		} else if input == 3 {
			hapusPetugas(petugas, nPetugas)
		} else if input == 4 {
			cetakPetugas(*petugas, *nPetugas)
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
	if n > MAXPETUGAS-*nPetugas {
		fmt.Println("Banyaknya petugas melebihi batasan maksimum!")
		fmt.Printf("Banyaknya petugas baru yang ditambahkan berubah dari %d menjadi %d\n", n, MAXPETUGAS-*nPetugas)
		n = MAXPETUGAS - *nPetugas
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
	}
	*nPetugas = i
	urutID(petugas, *nPetugas)
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
		if input == 1 || input == 2 || input == 3 || input == 4 {
			editDataPetugas(petugas, nPetugas, input)
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
		fmt.Println("ID tersebut tidak terdapat di dalam data Petugas!")
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
			urutID(petugas, nPetugas)
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
		fmt.Println("ID tersebut tidak terdapat di dalam data Petugas!")
	} else {
		i = idx
		for i < *nPetugas-1 {
			petugas[i] = petugas[i+1]
			i++
		}
		*nPetugas--
		fmt.Printf("Semua data Petugas dengan id %d berhasil dihapus :)\n", idHapus)
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
}