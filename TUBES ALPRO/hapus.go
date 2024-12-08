package main
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