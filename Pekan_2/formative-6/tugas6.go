package main

import "fmt"

//func Soal 1
func introduce(sentence *string, nama, gender, profesi, umur string) {
	*sentence = fmt.Sprintln("Pak", nama, "adalah seorang", profesi, "yang berusia", umur, "tahun")
	if gender == "perempuan" {
		*sentence = fmt.Sprintln("Bu", nama, "adalah seorang", profesi, "yang berusia", umur, "tahun")
	}

}

//func Soal 2
func TambahBuah(buah *[]string, tambah string) {
	*buah = append(*buah, tambah)
}

//func Soal 3
func tambahDataFilm(judul, durasi, genre, tahun string, daftar *[]map[string]string) {
	film := map[string]string{}
	film["judul"] = judul
	film["durasi"] = durasi
	film["genre"] = genre
	film["tahun"] = tahun

	*daftar = append(*daftar, film)
}

//func Soal 4
func genap(bil *[4]int) {
	for i := 0; i < len(*bil); i++ {
		(*bil)[i] = (*bil)[i] * 2
	}
}

//struct Soal 5
type inibuah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func main() {
	//Jawaban Soal 1
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	//Jawaban Soal 2
	var buah = []string{}
	TambahBuah(&buah, "Jeruk")
	TambahBuah(&buah, "Semangka")
	TambahBuah(&buah, "Mangga")
	TambahBuah(&buah, "Strawberry")
	TambahBuah(&buah, "Durian")
	TambahBuah(&buah, "Manggis")
	TambahBuah(&buah, "Alpukat")

	angka := 1
	for _, buahs := range buah {
		fmt.Println(angka, ".", buahs)
		angka++
	}
	fmt.Println("")
	//Jawaban Soal 3
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	angkas := 1
	for _, daftars := range dataFilm {
		fmt.Println(angkas, ".", "title:", daftars["judul"])
		fmt.Println("    duration:", daftars["durasi"])
		fmt.Println("    genre:", daftars["genre"])
		fmt.Println("    year:", daftars["tahun"])
		angkas++
	}
	fmt.Println("")
	//Jawaban Soal 4
	bil := [4]int{3, 5, 7, 9}
	fmt.Println("Sebelum:", bil)
	genap(&bil)
	fmt.Println("Sesudah:", bil)

	//Jawaban Soal 5
	var nanas = inibuah{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000}
	var jeruk = inibuah{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000}
	var semangka = inibuah{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: false, harga: 10000}
	var pisang = inibuah{nama: "Pisang", warna: "Kuning", adaBijinya: true, harga: 5000}

	fmt.Println("nama        :", nanas.nama)
	fmt.Println("warna       :", nanas.warna)
	fmt.Println("ada bijinya :", nanas.adaBijinya)
	fmt.Println("harga       :", nanas.harga)
	fmt.Println("")
	fmt.Println("nama        :", jeruk.nama)
	fmt.Println("warna       :", jeruk.warna)
	fmt.Println("ada bijinya :", jeruk.adaBijinya)
	fmt.Println("harga       :", jeruk.harga)
	fmt.Println("")
	fmt.Println("nama        :", semangka.nama)
	fmt.Println("warna       :", semangka.warna)
	fmt.Println("ada bijinya :", semangka.adaBijinya)
	fmt.Println("harga       :", semangka.harga)
	fmt.Println("")
	fmt.Println("nama        :", pisang.nama)
	fmt.Println("warna       :", pisang.warna)
	fmt.Println("ada bijinya :", pisang.adaBijinya)
	fmt.Println("harga       :", pisang.harga)
}
