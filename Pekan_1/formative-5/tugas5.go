package main

import (
	"fmt"
)

// fuc soal 1
func introduce(nama, gender, profesi, umur string) string {
	sapa := fmt.Sprintln("Pak", nama, "adalah seorang", profesi, "yang berusia", umur, "tahun")
	if gender == "perempuan" {
		return fmt.Sprintln("Bu", nama, "adalah seorang", profesi, "yang berusia", umur, "tahun")
	}
	return sapa
}

// func soal 2
func buahFavorit(nama string, buah ...string) string {
	hasil := "halo nama saya " + nama + " dan buah favorit saya adalah "
	for i, b := range buah {
		if i == len(buah)-1 {
			hasil += `"` + b + `"`
		} else {
			hasil += `"` + b + `", `
		}
	}
	return hasil
}

// func soal 3
var dataFilm = []map[string]string{}
var tambahDataFilm = func(tahun string, genre string, jam string, title string) {
	film := map[string]string{}
	film["tahun"] = tahun
	film["genre"] = genre
	film["jam"] = jam
	film["title"] = title
	dataFilm = append(dataFilm, film)
}

// func soal 4
func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

// func soal 5
func hitungLingkaran(luas *float32, keliling *float32, jariJari float32) {
	*luas = 3.14 * jariJari * jariJari
	*keliling = 2 * 3.14 * jariJari
}

func main() {
	//Jawaban Soal 1
	john := introduce("John", "laki-laki", "penulis", "30")
	sarah := introduce("Sarah", "perempuan", "model", "28")

	fmt.Println(john)
	fmt.Println(sarah)

	//Jawaban Soal 2
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)
	fmt.Println(buahFavoritJohn)

	//Jawaban Soal 3
	tambahDataFilm("1999", "action", "2 jam", "LOTR")
	tambahDataFilm("2019", "action", "2 jam", "avenger")
	tambahDataFilm("2004", "action", "2 jam", "spiderman")
	tambahDataFilm("2004", "horror", "2 jam", "juon")
	for _, item := range dataFilm {
		fmt.Println(item)
	}

	//Jawaban Soal 4
	angka := 5
	fmt.Println("faktorial dari", angka)
	fmt.Println("menghasilkan", factorial(angka))

	//Jawaban Soal 5
	var luasLingkaran float32
	var kelilingLingkaran float32

	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, 7)

	fmt.Println("Luas Lingkaran    :", luasLingkaran)
	fmt.Println("Keliling Lingkaran:", kelilingLingkaran)
}
