package main

import (
	"fmt"
	"sanbercode-golang-batch-77-Adrian/Pekan_2/formative-9/belajarpackage"
)

func main() {
	//Jawaban Soal 1
	var datar belajarpackage.HitungBangunDatar

	datar = belajarpackage.SegitigaSamaSisi{Alas: 6, Tinggi: 4}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Println("luas:", datar.Luas())
	fmt.Println("keliling:", datar.Keliling())
	fmt.Println("")

	datar = belajarpackage.PersegiPanjang{Panjang: 8, Lebar: 5}
	fmt.Println("Persegi Panjang:")
	fmt.Println("luas:", datar.Luas())
	fmt.Println("keliling:", datar.Keliling())
	fmt.Println("")

	var ruang belajarpackage.HitungBangunRuang

	ruang = belajarpackage.Tabung{JariJari: 7, Tinggi: 10}
	fmt.Println("Tabung:")
	fmt.Println("volume:", ruang.Volume())
	fmt.Println("luas permukaan:", ruang.LuasPermukaan())
	fmt.Println("")

	ruang = belajarpackage.Balok{Panjang: 5, Lebar: 4, Tinggi: 3}
	fmt.Println("Balok:")
	fmt.Println("volume:", ruang.Volume())
	fmt.Println("luas permukaan:", ruang.LuasPermukaan())
	fmt.Println("")

	//Jawaban Soal 2
	var info belajarpackage.InfoHP

	info = belajarpackage.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{`Mystic Bronze, Mystic White, Mystic Black`},
	}
	info.Tampilkan()
	fmt.Println("")

	//Jawaban Soal 3
	LP := belajarpackage.LuasPersegi
	fmt.Println(LP(4, true))
	fmt.Println(LP(8, false))
	fmt.Println(LP(0, true))
	fmt.Println(LP(0, false))

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
	hasilTambah := belajarpackage.Tambah(10, 2)
	fmt.Println("10 + 2 =", hasilTambah)

	hasilKali := belajarpackage.Kali(10, 2)
	fmt.Println("10 x 2 =", hasilKali)
}
