package main

import (
	"fmt"
	"strings"
)

func main() {
	//Jawaban Soal 1
	for i := 1; i <= 7; i++ {
		fmt.Println(strings.Repeat("#", i))
	}

	//Jawaban Soal 2
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2], kalimat[3], kalimat[4], kalimat[5], kalimat[6])

	//Jawaban Soal 3
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	for i, sayuran := range sayuran {
		i++
		fmt.Println(i, sayuran)
	}

	//Jawaban Soal 4
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for kata, nilai := range satuan {
		fmt.Println(kata, "=", nilai)
	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	fmt.Println(volume)

	//Jawaban Soal 5
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volumeB := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volumeB)
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}
