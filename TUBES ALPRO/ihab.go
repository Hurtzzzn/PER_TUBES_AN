package main

import "fmt"

type sate struct {
	menu          string
	jumlah, harga int
}

const NMAX int = 20

type tabSate [NMAX]sate

func main() {
	var pilih int
	var stop bool = false
	for !stop {
		menu()
		var i int
		var s tabSate
		var a, d, p string

		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			pemesanan(&s, &i)
		case 2:
			PromoHariIni(d)
		case 3:
			listMenu(a)

		case 4:
			pembayaran(p)
		}
		if pilih == 5 {
			fmt.Println("Terima kasih telah melakukan pemesanan!")
			stop = true
		}
	}

}

func menu() {
	fmt.Println("---------------------------")
	fmt.Println("S E L A M A T   D A T A N G")
	fmt.Println("             DI            ")
	fmt.Println("    SATE PADANG ENJOYER    ")
	fmt.Println("---------------------------")
	fmt.Println("  1. Pemesanan📝          ")
	fmt.Println("  2. Promo Hari Ini!🙀    ")
	fmt.Println("  3. list Menu             ")
	fmt.Println("  4. Pembayaran            ")
	fmt.Println("  5. Exit                  ")
	fmt.Println("---------------------------")
	fmt.Println("  Pilih (1/2/3/4/5/)?      ")

}

func pemesanan(S *tabSate, n *int) {
	var i, total int

	fmt.Println("Berapa menu yang akan kamu pesan?")
	fmt.Scan(&*n)
	fmt.Printf("Tuliskan %d menu beserta jumlah makanan/minuman yang akan kamu pesan yaa (format: namaMenu jumlahPesanan hargaPerMenu) ", *n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&S[i].menu, &S[i].jumlah, &S[i].harga)
	}
	fmt.Println("Silahkan cek kembali pesananmu: ")
	for i = 0; i < *n; i++ {
		fmt.Println(S[i].menu, S[i].jumlah)
	}
	for i = 0; i < *n; i++ {
		total = total + S[i].jumlah*S[i].harga
	}
	fmt.Printf("Total: %d ", total)
}

func PromoHariIni(d string) {
	fmt.Println("maaf hari ini belum tersedia promo, kembali lagi ketika promo tersedia")
}

func listMenu(a string) {
	fmt.Println("------ DAFTAR MENU -------")
	fmt.Println("..........................")
	fmt.Println("  A. Sate                 ")
	fmt.Println("     - Daging ayam        ")
	fmt.Println("     - Daging sapi        ")
	fmt.Println("     - Lidah              ")
	fmt.Println("     - Usus               ")
	fmt.Println("     - Jando              ")
	fmt.Println("  B. Kuah                 ")
	fmt.Println("     - Pariaman           ")
	fmt.Println("     - Derek              ")
	fmt.Println("     - Dagung-Dangung     ")
	fmt.Println("  C. Tambahan             ")
	fmt.Println("     - Ketupat            ")
	fmt.Println("  D. Topping              ")
	fmt.Println("     - Kerupuk pedas      ")
	fmt.Println("     - Kerupuk kulit      ")
	fmt.Println("     - Bawang Goreng      ")
}

func pembayaran(p string) {

	fmt.Println("maaf untuk metode pembayaran banya tersedia untuk Transfer Bank🙏🏻")
	fmt.Println("Silakan transfer ke rekening berikut:")
	fmt.Println("- Bank: TelkomBank")
	fmt.Println("- Nomor Rekening: 123-456-7890")
	fmt.Println("- Atas Nama: tukang sate ngidam sate")
	fmt.Println("- jika sudah, silahkan kirim bukti transfer ke no WA ini:")
	fmt.Println("   - 0123456789")

}
