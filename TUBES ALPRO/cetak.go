package main
import "fmt"
func cetakParkirMobil(pMobil tabParkirMobil, nMobil int) {
	/* IS : Array pMobil dengan banyak elemen nMobil terdefinsi
	   FS : Array pMobil dengan banyak elemen nMobil tercetak di layar
	*/
	var i int
	fmt.Println("Data Parkiran Mobil : ")
	for i = 0; i < nMobil; i++ {
		fmt.Println(pMobil[i].tipe, pMobil[i].hari, pMobil[i].noPol, pMobil[i].vip, pMobil[i].valet, pMobil[i].waktu.jamMasuk, pMobil[i].waktu.menitMasuk, pMobil[i].waktu.jamKeluar, pMobil[i].waktu.menitKeluar)
	}
}

func cetakParkirMotor(pMotor tabParkirMotor, nMotor int) {
	/* IS : Array pMotor dengan banyak elemen nMotor terdefinsi
	   FS : Array pMotor dengan banyak elemen nMotor tercetak di layar
	*/
	var i int
	fmt.Println("Data Parkiran Motor : ")
	for i = 0; i < nMotor; i++ {
		fmt.Println(pMotor[i].tipe, pMotor[i].hari, pMotor[i].noPol, pMotor[i].vip, pMotor[i].valet, pMotor[i].waktu.jamMasuk, pMotor[i].waktu.menitMasuk, pMotor[i].waktu.jamKeluar, pMotor[i].waktu.menitKeluar)
	}
}

func cetakParkirKeseluruhan(dataKendaraan tabDataKendaraan, nData, totalPendapatan int) {
	/* IS : Array dataKendaraan dengan banyak elemen nMotor terdefinsi
	   FS : Array dataKendaraan dengan banyak elemen nMotor tercetak di layar berdasarkan urutan tertentu
	        dan tercetak total pendapatan di akhir
	*/
	var i int
	fmt.Println("Data Parkiran Keseluruhan : ")
	for i = 0; i < nData; i++ {
		fmt.Printf("%s %s %s %t %t %02d:%02d %02d:%02d Rp.%d\n", dataKendaraan[i].tipe, dataKendaraan[i].hari, dataKendaraan[i].noPol, dataKendaraan[i].vip, dataKendaraan[i].valet, dataKendaraan[i].waktu.jamMasuk, dataKendaraan[i].waktu.menitMasuk, dataKendaraan[i].waktu.jamKeluar, dataKendaraan[i].waktu.menitKeluar, dataKendaraan[i].hargaParkir)
	}
	fmt.Println(totalPendapatan)
}
