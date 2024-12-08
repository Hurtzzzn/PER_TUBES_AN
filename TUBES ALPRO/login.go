package main
import "fmt"
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
		if input == 1 {
			statusLogin = false
			for !statusLogin {
				fmt.Print("Masukkan Username : ")
				fmt.Scan(&username)
				fmt.Print("Masukkan Password : ")
				fmt.Scan(&password)
				idx = searchIdxUsn(petugas, nPetugas, username)
				if cekLogin(petugas, nPetugas, username, password) {
					fmt.Println("Log In berhasil! Menuju ke menu petugas!")
					fmt.Println("--------------------------------------------------------------------------------------------")
					petugasMenu(petugas, idx)
					statusLogin = true
				} else {
					fmt.Println("Username atau Password Anda Salah! Silahkan input ulang!")
					fmt.Println("--------------------------------------------------------------------------------------------")
				}
			}
		}
	}
}