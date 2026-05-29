package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// func Soal 1
func munculkan(kalimat string, tahun int) {
	fmt.Println(kalimat)
	fmt.Println(tahun)
}

func coba() {
	defer munculkan("Golang Backend Development", 2021)
	fmt.Println("Hello Golang")
}

// func Soal 2
func kelilingSegitigaSamaSisi(sisi int, status bool) (string, error) {
	if sisi == 0 && status {
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	if sisi == 0 && !status {
		defer func() {
			message := recover()
			fmt.Println("Terjadi Error", message)
		}()
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	keliling := sisi * 3

	if status {
		return fmt.Sprintln("keliling segitiga sama sisinya dengan sisi", sisi, "cm adalah", keliling, "cm"), nil
	}

	return fmt.Sprintln(sisi), nil
}

// func Soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

// func Soal 4
func tambahPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

// func Soal 5
func tampilkanPhone(urutan int, phone string, wg *sync.WaitGroup) {
	if urutan > 2 {
		time.Sleep(time.Duration(urutan-2) * time.Second)
	}
	fmt.Println(urutan, phone)
	wg.Done()
}

func main() {
	// Jawaban Soal 1
	coba()
	fmt.Println("")

	//Jawaban Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
	fmt.Println("")

	//Jawaban Soal 3
	angka := 1

	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//Jawaban Soal 4
	var phones = []string{}
	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")

	sort.Strings(phones)
	for i, phone := range phones {
		time.Sleep(1 * time.Second)
		fmt.Println(i+1, phone)
	}
	fmt.Println("")

	//Jawaban Soal 5
	var phones2 = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones2)

	var wg sync.WaitGroup
	for i, phone2 := range phones2 {
		wg.Add(1)
		go tampilkanPhone(i+1, phone2, &wg)
	}
	wg.Wait()
}
