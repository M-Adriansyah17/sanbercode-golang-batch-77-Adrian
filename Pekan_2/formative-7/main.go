package main

import "fmt"

//struct Soal 1
type Mahasiswa struct {
	Nama string
	NIM  string
	Usia int
}

//struct dan method Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (sg segitiga) luas() int {
	return (sg.alas * sg.tinggi) / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

//struct dan method Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (ph *phone) tambahwarna(warna string) {
	ph.colors = append(ph.colors, warna)
}

//struct dan func Soal 4
type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, daftar *[]movie) {
	var film movie
	film.title = title
	film.genre = genre
	film.duration = duration
	film.year = year

	*daftar = append(*daftar, film)
}

//interface dan struct Soal 5
type Hewan interface {
	Suara() string
}

type Kucing struct {
	nama string
}

type Anjing struct {
	nama string
}

func (k Kucing) Suara() string {
	return "Meong"
}

func (a Anjing) Suara() string {
	return "Guk Guk"
}

func main() {

	//Jawaban Soal 1
	var Budi = Mahasiswa{}
	Budi.Nama = "Budi"
	Budi.NIM = "26SK8821"
	Budi.Usia = 19

	fmt.Println("Nama Mahasiswa:", Budi.Nama)
	fmt.Println("NIM Mahasiswa:", Budi.NIM)
	fmt.Println("Usia Mahasiswa:", Budi.Usia)
	fmt.Println("")

	//Jawaban Soal 2
	var sg segitiga
	sg.alas = 10
	sg.tinggi = 5

	var p persegi
	p.sisi = 4

	var pp persegiPanjang
	pp.panjang = 8
	pp.lebar = 3

	fmt.Println("luas segitiga:", sg.luas())
	fmt.Println("luas persegi:", p.luas())
	fmt.Println("luas persegi panjang:", pp.luas())
	fmt.Println("")

	//Jawaban Soal 3
	var hp phone
	hp.name = "Infinix GT 50 Pro"
	hp.brand = "Infinix"
	hp.year = 2026

	hp.tambahwarna("Hitam")
	hp.tambahwarna("Merah")
	hp.tambahwarna("Silver")

	fmt.Println("nama:", hp.name)
	fmt.Println("brand:", hp.brand)
	fmt.Println("tahun:", hp.year)
	fmt.Println("warna:", hp.colors)
	fmt.Println("")

	//Jawaban Soal 4
	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	angkas := 1
	for _, daftars := range dataFilm {
		fmt.Println(angkas, ".", "title:", daftars.title)
		fmt.Println("    duration:", daftars.duration)
		fmt.Println("    genre:", daftars.genre)
		fmt.Println("    year:", daftars.year)
		angkas++
	}
	fmt.Println("")

	//Jawaban Soal 5
	var kucing Kucing
	kucing.nama = "Rick"

	var anjing Anjing
	anjing.nama = "Sven"

	var h Hewan

	h = kucing
	fmt.Println("(Kucing)", kucing.nama, "bersuara:", h.Suara())

	h = anjing
	fmt.Println("(Anjing)", anjing.nama, "bersuara:", h.Suara())

}
