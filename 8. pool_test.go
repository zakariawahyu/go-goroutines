package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Pool
- Pool adalah implementasi design pattern bernama object pool pattern
- Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya,
kita bisa mengambil dari Pool dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke pool
- Implementasi Pool di golang ini sudah aman dari problem race condition
 */

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New : func() interface{} { // func new ketika data nill maka akan return new
			return "NEW"
		},
	}

	pool.Put("Zakaria")
	pool.Put("Wahyu")
	pool.Put("Nur")
	pool.Put("Utomo")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
