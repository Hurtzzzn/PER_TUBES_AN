package main
func hitungHarga(dataKendaraan *tabDataKendaraan, idxData int) {
	/* IS : Array dataKendaraan dengan indeks idxData terdefinisi (kecuali field hargaParkir)
	      FS : hargaParkir terdefinisi dengann ketentuan :
	           - Jika kendaraan bertipe Mobil akan dikenakan harga Rp.10.000 untuk jam pertama
	   		  dan Rp.5.000 untuk jam berikutnya. Jika parkir di tempat vip akan dikenakan tambahan
	   		  harga Rp.20.000. Jika menggunakan jasa valet akan dikenakan tambahan harga Rp.50.000
	   		- Jika kendaraan bertipe Motor akan dikenakan harga Rp.5.000 untuk jam pertama
	   		  dan Rp.2.000 untuk jam berikutnya. Jika parkir di tempat vip akan dikenakan tambahan
	   		  harga Rp.10.000. Jika menggunakan jasa valet akan dikenakan tambahan harga Rp.20.000
	   		- Jika kendaraan parkir di hari Jumat akan mendapatkan diskon sebesar 20%
	*/
	var harga, totalMenit, totalJam int
	totalMenit = (dataKendaraan[idxData].waktu.jamKeluar*60 + dataKendaraan[idxData].waktu.menitKeluar) - (dataKendaraan[idxData].waktu.jamMasuk*60 + dataKendaraan[idxData].waktu.menitMasuk)
	totalJam = totalMenit / 60
	if totalMenit%60 != 0 {
		totalJam++
	}
	if dataKendaraan[idxData].tipe == "Mobil" {
		harga = 10000 + (totalJam-1)*5000
		if dataKendaraan[idxData].vip {
			harga += 20000
		}
		if dataKendaraan[idxData].valet {
			harga += 50000
		}
	} else if dataKendaraan[idxData].tipe == "Motor" {
		harga = 5000 + (totalJam-1)*2000
		if dataKendaraan[idxData].vip {
			harga += 10000
		}
		if dataKendaraan[idxData].valet {
			harga += 20000
		}
	}
	if dataKendaraan[idxData].hari == "Jumat" {
		harga = harga - (harga * 20 / 100)
	}
	dataKendaraan[idxData].hargaParkir = harga
}