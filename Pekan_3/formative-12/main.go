package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// struct mahasiswa
type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Nilai       uint   `json:"nilai"`
	IndeksNilai string `json:"indeks_nilai"`
}

// struct untuk input nilai
type InputNilai struct {
	Nama       string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai      uint   `json:"nilai"`
}

// database
var (
	nilaiNilaiMahasiswa = []NilaiMahasiswa{}
	mu                  sync.Mutex
	idCounter           uint = 1
)

// Syarat pemberian nilai
func hitungIndeks(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

// middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="NilaiMahasiswa"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		//username dan password
		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

// post nilai
func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(`{"error": "Content-Type harus application/json"}`))
			return
		}

		var input InputNilai
		decodeJSON := json.NewDecoder(r.Body)
		decodeJSON.DisallowUnknownFields()
		if err := decodeJSON.Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Format JSON tidak valid atau terdapat field yang tidak diizinkan (id/indeks_nilai tidak boleh dikirim)"}`))
			return
		}

		if input.Nama == "" || input.MataKuliah == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Field nama dan mata_kuliah wajib diisi"}`))
			return
		}

		if input.Nilai > 100 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Nilai tidak boleh lebih dari 100"}`))
			return
		}

		mu.Lock()
		data := NilaiMahasiswa{
			ID:          idCounter,
			Nama:        input.Nama,
			MataKuliah:  input.MataKuliah,
			Nilai:       input.Nilai,
			IndeksNilai: hitungIndeks(input.Nilai),
		}
		idCounter++
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, data)
		mu.Unlock()

		response, _ := json.Marshal(map[string]interface{}{
			"message": "Data berhasil ditambahkan",
			"data":    data,
		})
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "ERROR.... Method tidak diizinkan"}`))
}

// get nilai
func GetNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		mu.Lock()
		defer mu.Unlock()

		response, err := json.Marshal(map[string]interface{}{
			"total": len(nilaiNilaiMahasiswa),
			"data":  nilaiNilaiMahasiswa,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error": "ERROR.... Method tidak diizinkan"}`))
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	//alamatnya
	http.Handle("/nilai/post", Auth(http.HandlerFunc(PostNilai)))
	http.HandleFunc("/nilai", GetNilai)

	fmt.Println("server running at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
