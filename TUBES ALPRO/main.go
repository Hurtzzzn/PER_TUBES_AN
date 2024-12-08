package main

import "fmt"

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
		if inputMenu == 1 {
			adminMenu(&petugas, &nPetugas)
		} else if inputMenu == 2 {
			loginMenu(petugas, nPetugas)
		}
	}
}
