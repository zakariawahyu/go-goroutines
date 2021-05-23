package go_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/**
GOMAXPROCS
- Sebelumnya di awal kita sudah bahas bahwa gorouine itu sebenarnya dijalankan di dalam Thread
- Pertanyaannya, seberapa banyak Thread yang ada di Golang ketika aplikasi berjalan ?
- Untuk mengetahui berapa jumlah Thread, kita bisa menggunakan GOMAXPROCS, sebuah function di package runtime yang bisa
kita gunakan untuk mengubah jumlah Thread atau mengambil jumlah Threadnya
- Secara default, jumlah Thread di golang itu sebanyak jumlah CPU di komputer kita
- Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.numCPU()
 */

func TestGomacprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU : ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread : ", totalThread)

	// secara tidak langsnug pasti ada 2 goroutine yang sudah aktif
	// goroutine untuk menjalankan kode program dan go runtimenya
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}

/**
- Menambah jumlah Thread, tidak berati membuat aplikasi kita menjadi lebih cepat
- Karena pada saat yang sama, 1 CPU hanya akan menjalankan 1 goroutine dengan 1 Thread
- Oleh karena ini, jika ingin menambah throught aplikasi, disarankan melakukan vertical scaling (dengan menambah jumlah CPU) atau horizontal scaling (menambah node baru)
 */

func TestChangeThread(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU : ", totalCpu)

	// sebelum get data harus set dulu berapa jumlah thread yang akan dirubah
	runtime.GOMAXPROCS(20)
	// kemudian get data
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread : ", totalThread)

	// secara tidak langsnug pasti ada 2 goroutine yang sudah aktif
	// goroutine untuk menjalankan kode program dan go runtimenya
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine : ", totalGoroutine)

	group.Wait()
}
