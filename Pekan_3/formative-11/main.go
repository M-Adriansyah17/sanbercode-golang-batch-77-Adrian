package main

import (
	"encoding/json"
	"fmt"
)

// struct Soal 2
type Phone struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
	Tahun  string `json:"tahun"`
}

// struct Soal 3
type Book struct {
	Title, Desc, Author string
	ReleaseYear         int
}

func main() {
	var stringJson = `[
	{"vendor": "samsung", "model": "samsung j2 prime", "tahun": "2016"},
	{"vendor": "apple", "model": "iphone 11 pro", "tahun": "2019"},
	{"vendor": "xiaomi", "model": "redmi 15", "tahun": "2025"},
	{"vendor": "sony", "model": "sony xperia 10", "tahun": "2019"},
	{"vendor": "infinix", "model": "infinix hot 11", "tahun": "2021"}
	]`

	var data1 []map[string]string
	var Jsondata = []byte(stringJson)

	var err1 = json.Unmarshal(Jsondata, &data1)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	//jawaban Soal 1
	for _, item := range data1 {
		if item["vendor"] == "samsung" {
			fmt.Println("vendor:", item["vendor"])
			fmt.Println("model:", item["model"])
			fmt.Println("tahun:", item["tahun"])
		}
	}
	fmt.Println("")

	//jawaban Soal 2
	var data2 []Phone

	var err2 = json.Unmarshal(Jsondata, &data2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	for _, item2 := range data2 {
		if item2.Vendor == "sony" {
			fmt.Println("vendor:", item2.Vendor)
			fmt.Println("model:", item2.Model)
			fmt.Println("tahun:", item2.Tahun)
		}
	}
	fmt.Println("")

	//jawaban Soal 3
	var buku = []Book{
		{"Dune", "Perebutan rempah di gurun", "Frank Herbert", 1965},
		{"Steve Jobs", "Biografi resmi berdasarkan wawancara langsung dengan Steve Jobs dan kerabatnya", "Walter Isaacson", 2011},
		{"Atomic Habits", "Panduan praktis membangun kebiasaan baik melalui perubahan kecil yang dilakukan secara konsisten setiap hari", "James Clear", 2018},
		{"Cosmos", "Perjalanan puitis menjelajahi alam semesta, dari asal-usul kehidupan hingga evolusi bintang dan galaksi", "Carl Sagan", 1980},
		{"Thinking, Fast and Slow", "Eksplorasi dua sistem berpikir manusia, serta berbagai bias kognitif", "Daniel Kahneman", 2011},
	}

	var jsonData, err = json.Marshal(buku)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var bukuJson = string(jsonData)
	fmt.Println(bukuJson)
}
