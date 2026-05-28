package belajarpackage

import "fmt"

// interface, struct, dan method Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}
func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}
func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

const pi = 3.14

func (t Tabung) Volume() float64 {
	return pi * t.JariJari * t.JariJari * t.Tinggi
}
func (t Tabung) LuasPermukaan() float64 {
	return 2 * pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}
func (b Balok) LuasPermukaan() float64 {
	return float64(2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi))
}

// struct, interface, dan method soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type InfoHP interface {
	Tampilkan()
}

func (hp Phone) Tampilkan() {
	fmt.Println("name:", hp.Name)
	fmt.Println("brand:", hp.Brand)
	fmt.Println("year:", hp.Year)
	fmt.Println("colors:", hp.Colors)
}

// func Soal 3
func LuasPersegi(sisi int, status bool) interface{} {
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

//func Soal 5
func Tambah(a int, b int) int {
	return a + b
}

func Kali(a int, b int) int {
	return a * b
}
