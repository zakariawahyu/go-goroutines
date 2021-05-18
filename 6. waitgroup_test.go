package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Wait Group
- Wait group adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
- Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine,
tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
- Kasus seperti ini bisa menggunakan Wait Group
- Untuk menendai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah proses goroutinw selesai kita bisa menggunakan method Dhone()
- Untuk menunggu semua proses selesai, kita bisa menggunakan method Wait()
 */

func RunAsynchronus(group *sync.WaitGroup)  {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
	}

	group.Wait()
	fmt.Println("Done")
}