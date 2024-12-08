package main
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