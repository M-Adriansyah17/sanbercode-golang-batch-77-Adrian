package main

import (
	"fmt"
	"strings"
)

func main() {

	//Jawaban Soal 1
	kalimat := "halo halo bandung"
	angka := 2021
	awal := "halo"
	ganti := "Hi"

	kalimat2 := strings.Replace(kalimat, awal, ganti, -1)
	fmt.Println(`"`+kalimat2+`"`, "-", angka)

	//Jawaban Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("John dapat nilai A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("John dapat nilai B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("John dapat nilai C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("John dapat nilai D")
	} else {
		fmt.Println("John dapat nilai E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Doe dapat nilai A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Doe dapat nilai B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Doe dapat nilai C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Doe dapat nilai D")
	} else {
		fmt.Println("Doe dapat nilai E")
	}

	//Jawaban Soal 3
	var tanggal = 12
	var bulan = 9
	var tahun = 2006
	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, "Februari", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	default:
		fmt.Println("Tidak Valid")
	}

	//Jawaban Soal 4
	tahunlahir := 2006
	if tahunlahir >= 1995 && tahunlahir <= 2015 {
		fmt.Println("Kamu Generasi Z")
	} else if tahunlahir >= 1980 && tahunlahir <= 1994 {
		fmt.Println("Kamu Generasi Y(millenials)")
	} else if tahunlahir >= 1965 && tahunlahir <= 1979 {
		fmt.Println("Kamu Generasi X")
	} else if tahunlahir >= 1944 && tahunlahir <= 1964 {
		fmt.Println("Kamu Generasi Baby Boomer")
	}

	//Jawaban Soal 5
	for i := 1; i <= 20; i++ {
		if i%2 == 1 && i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i%2 == 1 {
			fmt.Println(i, "- Santai")
		} else if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		}
	}
}
