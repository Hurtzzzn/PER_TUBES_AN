package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	const phi float64 = 3.14
	var a, b, c, r float64
	var alas, tinggi, panjang, lebar float64
	var LSegitiga, LPersegiPanjang, LLingkaran float64
	var kllSegitiga, kllPersegiPanjang, kllLingkaran float64
	var pilihan int
	for pilihan != 4 {
		fmt.Println("============================================================================================")
		fmt.Println("*||*                                                                                    *||*")
		fmt.Println("*||*                                   *MENU UTAMA*                                     *||*")
		fmt.Println("*||*                                                                                    *||*")
		fmt.Println("============================================================================================")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Println("*** Bangun Datar ***")
		fmt.Println("1. Segitiga")
		fmt.Println("2. Persegi Panjang")
		fmt.Println("3. Lingkaran")
		fmt.Println("4. Keluar")
		fmt.Println("--------------------------------------------------------------------------------------------")
		fmt.Print("Masukkan (1/2/3/4) : ")
		fmt.Scan(&pilihan)
		ClearScreen()
		if pilihan == 1 {
			pilihan = 0
			for pilihan != 3 {
				fmt.Println("============================================================================================")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("*||*                                *MENU SEGITIGA*                                     *||*")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("============================================================================================")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Println("*** Bangun Datar ***")
				fmt.Println("1. Keliling Segitiga")
				fmt.Println("2. Luas Segitiga")
				fmt.Println("3. Keluar")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan (1/2/3) : ")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Print("Masukkan sisi 1 : ")
					fmt.Scan(&a)
					fmt.Print("Masukkan sisi 2 : ")
					fmt.Scan(&b)
					fmt.Print("Masukkan sisi 3 : ")
					fmt.Scan(&c)
					kllSegitiga = a + b + c
					fmt.Println("Keliling Segitiga adalah :", kllSegitiga)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				} else if pilihan == 2 {
					fmt.Print("Masukkan Alas : ")
					fmt.Scan(&alas)
					fmt.Print("Masukkan Tinggi : ")
					fmt.Scan(&tinggi)
					LSegitiga = (alas * tinggi) / 2
					fmt.Println("Luas Segitiga adalah :", LSegitiga)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				}
			}
			ClearScreen()
		} else if pilihan == 2 {
			pilihan = 0
			for pilihan != 3 {
				fmt.Println("============================================================================================")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("*||*                             *MENU PERSEGI PANJANG*                                 *||*")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("============================================================================================")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Println("*** Bangun Datar ***")
				fmt.Println("1. Keliling Persegi Panjang")
				fmt.Println("2. Luas Persegi Panjang")
				fmt.Println("3. Keluar")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan (1/2/3) : ")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Print("Masukkan Panjang : ")
					fmt.Scan(&panjang)
					fmt.Print("Masukkan Lebar : ")
					fmt.Scan(&lebar)
					kllPersegiPanjang = 2 * (panjang + lebar)
					fmt.Println("Keliling Persegi Panjang adalah :", kllPersegiPanjang)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				} else if pilihan == 2 {
					fmt.Print("Masukkan Panjang : ")
					fmt.Scan(&panjang)
					fmt.Print("Masukkan Lebar : ")
					fmt.Scan(&lebar)
					LPersegiPanjang = panjang * lebar
					fmt.Println("Luas Persegi Panjang adalah :", LPersegiPanjang)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				}
			}
			ClearScreen()
		} else if pilihan == 3 {
			pilihan = 0
			for pilihan != 3 {
				fmt.Println("============================================================================================")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("*||*                                *MENU LINGKARAN*                                    *||*")
				fmt.Println("*||*                                                                                    *||*")
				fmt.Println("============================================================================================")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Println("*** Bangun Datar ***")
				fmt.Println("1. Keliling Lingkaran")
				fmt.Println("2. Luas Lingkaran")
				fmt.Println("3. Keluar")
				fmt.Println("--------------------------------------------------------------------------------------------")
				fmt.Print("Masukkan (1/2/3) : ")
				fmt.Scan(&pilihan)
				if pilihan == 1 {
					fmt.Print("Masukkan Jari-Jari : ")
					fmt.Scan(&r)
					kllLingkaran = 2 * phi * r
					fmt.Println("Keliling Lingkaran adalah :", kllLingkaran)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				} else if pilihan == 2 {
					fmt.Print("Masukkan Jari-Jari : ")
					fmt.Scan(&r)
					LLingkaran = phi * r * r
					fmt.Println("Luas Lingkaran adalah :", LLingkaran)
					fmt.Print("\nPress 'Enter' to continue...")
					fmt.Scanln()
					bufio.NewReader(os.Stdin).ReadBytes('\n')
					ClearScreen()
				}
			}
			ClearScreen()
		}
	}
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
