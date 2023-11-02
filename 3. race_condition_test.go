package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/*
*
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
				x += 1 // x = x+ 1
			}
		}()
	}

	// hasilnya seharusnya 100.000 data variabel x
	// race condition ini bisa jadi beberapa go routine mengakses variabel yang sama
	// maka dari itu otomatis nilainya hilang sebagian
	// race condition = go routine yang balapan buat merubah variabelnya
	// Cara mengatasi menggunakan mutex locking dan unlocking yang akan dibahas materi selanjutnya
	time.Sleep(time.Second)
	fmt.Println("Counter : ", x)
}

// Novalagung version
/*
- Race condition adalah kondisi di mana lebih dari satu goroutine, mengakses data yang sama pada waktu yang bersamaan (benar-benar bersamaan)
- Dalam concurrency programming situasi seperti ini sering terjadi dan akan menjadi kacau
*/

type Counter struct {
	val int
}

func (c *Counter) Add(int) {
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}

func TestRaceCondition2(t *testing.T) {
	var meter Counter
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				meter.Add(1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter : ", meter.Value())
}
