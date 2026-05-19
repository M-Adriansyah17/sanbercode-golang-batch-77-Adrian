package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Jawaban Soal 1
	//Bootcamp Digital Skill Sanbercode Golang
	var B = "Bootcamp "
	var D = "Digital "
	var S = "Skill "
	var C = "Sanbercode "
	var G = "Golang"
	var kalimat = (B + D + S + C + G)
	fmt.Println(kalimat)

	//Jawaban Soal 2
	halo := "Halo Dunia"
	lama := "Dunia"
	baru := "Golang"
	newhalo := strings.Replace(halo, lama, baru, -1)
	fmt.Println(newhalo)

	//Jawaban Soal 3
	var kataPertama = "saya"
	var kataKedua = "Senang"
	var kataKetiga = "belajar"
	var kataKeempat = strings.ToUpper("golang")
	var gabung = kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println(gabung)

	//Jawaban Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	b1, _ := strconv.Atoi(angkaPertama)
	b2, _ := strconv.Atoi(angkaKedua)
	b3, _ := strconv.Atoi(angkaKetiga)
	b4, _ := strconv.Atoi(angkaKeempat)
	total := b1 + b2 + b3 + b4
	fmt.Println(total)

	//Jawaban Soal 5
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	alasS, _ := strconv.Atoi(alasSegitiga)
	tinggiS, _ := strconv.Atoi(tinggiSegitiga)
	panjangP, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarP, _ := strconv.Atoi(lebarPersegiPanjang)

	var luasPersegiPanjang int = panjangP * lebarP
	var kelilingPersegiPanjang int = 2 * (panjangP + lebarP)
	var luasSegitiga int = (alasS * tinggiS) / 2

	fmt.Println("luas Persegi Panjang: ", luasPersegiPanjang)
	fmt.Println("keliling Persegi Panjang: ", kelilingPersegiPanjang)
	fmt.Println("luas Segitiga: ", luasSegitiga)
}
