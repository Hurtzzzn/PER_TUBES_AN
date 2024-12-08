package main

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