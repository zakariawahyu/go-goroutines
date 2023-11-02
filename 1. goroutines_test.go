package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
Membuat Go Routines
- Untuk membuat goroutines di golang sangatlah sederhana
- Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
- Saat sebuah function kita jalankan dalam goroutines, function tersebut akan berjalan secara asynchronus, artinya tidak akan ditunggu sampai function tersebut selesai
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

Go Routines sangat ringan
- Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
- Kita bisa membuat ribuan, bahkan sampai juataan go routines tanpa takut boros memory
- Tidak seperti Thread yang ukuran berat, goroutine sangat ringan
*/

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// akan runing secara asynchronus (tidak akan ditunggu sampai test selesai)
	// jika function return value, maka tidak akan bisa menamngkap valuenya
	go RunHelloWorld()

	fmt.Println("Ups")

	// untuk mengakali aplikasi bair menunggu go routine seelsai di eksekusi karena berjalan secara asynchronus
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	println("Disaplay number : ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

// Novalagung version
func cetak(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func TestGoroutine(t *testing.T) {
	go cetak(500, "Apakabar")
	cetak(500, "Hallo")

	time.Sleep(time.Second)
}
