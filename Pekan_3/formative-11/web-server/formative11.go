package main

import (
	"fmt"
	"math"
	"net/http"
)

var jariJari = 7.0
var tinggi = 10.0

func hitungTabung(w http.ResponseWriter, r *http.Request) {
	var volume = math.Pi * jariJari * jariJari * tinggi
	var luasAlas = math.Pi * jariJari * jariJari
	var kelilingAlas = 2 * math.Pi * jariJari

	fmt.Fprintln(w, "jariJari :", jariJari)
	fmt.Fprintln(w, "tinggi :", tinggi)
	fmt.Fprintln(w, "volume :", volume)
	fmt.Fprintln(w, "luas alas :", luasAlas)
	fmt.Fprintln(w, "keliling alas :", kelilingAlas)
}

func main() {
	http.HandleFunc("/", hitungTabung)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
