package main

import (
	"fmt"
	"sanbercode-golang-batch-77-Adrian/Pekan_2/formative-8/matematika"
) //import package matematika

// interface, struct, dan method Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}
func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}
func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

const pi = 3.14

func (t tabung) volume() float64 {
	return pi * t.jariJari * t.jariJari * t.tinggi
}
func (t tabung) luasPermukaan() float64 {
	return 2 * pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}
func (b balok) luasPermukaan() float64 {
	return float64(2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi))
}

// struct, interface, dan method soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type infoHP interface {
	tampilkan()
}

func (hp phone) tampilkan() {
	fmt.Println("name:", hp.name)
	fmt.Println("brand:", hp.brand)
	fmt.Println("year:", hp.year)
	fmt.Println("colors:", hp.colors)
}

// func Soal 3
func luasPersegi(sisi int, status bool) interface{} {
	if sisi == 0 && status {
		return "Maaf anda belum menginput sisi dari persegi"
	}
	if sisi == 0 && !status {
		return nil
	}

	hasil := sisi * sisi

	if status {
		return fmt.Sprint("luas persegi dengan sisi ", sisi, " cm adalah ", hasil, " cm")
	}
	return hasil
}

func main() {
	//Jawaban Soal 1
	var datar hitungBangunDatar

	datar = segitigaSamaSisi{alas: 6, tinggi: 4}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("luas:", datar.luas())
	fmt.Println("keliling:", datar.keliling())
	fmt.Println("")

	datar = persegiPanjang{panjang: 8, lebar: 5}
	fmt.Println("Persegi Panjang:")
	fmt.Println("luas:", datar.luas())
	fmt.Println("keliling:", datar.keliling())
	fmt.Println("")

	var ruang hitungBangunRuang

	ruang = tabung{jariJari: 7, tinggi: 10}
	fmt.Println("Tabung:")
	fmt.Println("volume:", ruang.volume())
	fmt.Println("luas permukaan:", ruang.luasPermukaan())
	fmt.Println("")

	ruang = balok{panjang: 5, lebar: 4, tinggi: 3}
	fmt.Println("Balok:")
	fmt.Println("volume:", ruang.volume())
	fmt.Println("luas permukaan:", ruang.luasPermukaan())
	fmt.Println("")

	//Jawaban Soal 2
	var info infoHP

	info = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{`Mystic Bronze, Mystic White, Mystic Black`},
	}
	info.tampilkan()
	fmt.Println("")

	//Jawaban Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//Jawaban Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)

	total := 0
	for _, v := range angka1 {
		total += v
	}
	for _, v := range angka2 {
		total += v
	}

	fmt.Println(
		prefix.(string), angka1[0], "+", angka1[1], "+", angka2[0], "+", angka2[1], "=", total)
	fmt.Println("")

	//Jawaban Soal 5
	hasilTambah := matematika.Tambah(10, 2)
	fmt.Println("10 + 2 =", hasilTambah)

	hasilKali := matematika.Kali(10, 2)
	fmt.Println("10 x 2 =", hasilKali)
}
