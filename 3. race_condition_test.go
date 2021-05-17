package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
Masalah dengan Goroutine
- Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurenct, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini sangat berbahaya ketika kita ingin melakukan manipulasi data variabel yang sama oleh beberapa goroutine secara bersamaan
- Hal ini bisa menyebabkan masalah yang namanya Race Condition
 */
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x+=1 // x = x+ 1
			}
		}()
	}

	// hasilnya seharusnya 100.000 data variabel x
	// race condition ini bisa jadi beberapa go routine mengakses variabel yang sama
	// maka dari itu otomatis nilainya hilang sebagian
	// race condition = go routine yang balapan buat merubah variabelnya
	// Cara mengatasi menggunakan mutex locking dan unlocking yang akan dibahas materi selanjutnya
	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
}
