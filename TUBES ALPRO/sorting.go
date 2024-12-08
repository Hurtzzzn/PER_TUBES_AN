package main
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