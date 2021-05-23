package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Cond
- Cond adalah implementasi locking berbasis kondisi
- Cond membutuhkan Locker(bisa menggunakan Mutex atau RWMutex) untuk implementasi lockingnya, namun berbeda
dengan Locker biasanya, di Cond terdapat function Wait untuk menunggu apakah perlu menungggu atau tidak
- Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi,
sedangkan fucntion Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan sync.NewCond(Locker)
 */

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func waitCondition(value int)  {
	defer group.Done()
	group.Add(1)

	// sama seperti lock(mutex) biasa
	// bedanya ada wait() yang artinya dia perlu mengunggu atau tidak
	cond.L.Lock()
	cond.Wait() // menunggu apakah setelah berhasil locking boleh lanjut atau tidak

	fmt.Println("Done : ", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go waitCondition(i)
	}
	
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // kasih arahan ke goroutine satu-satu
		}
	}()

	// Broadcast tanpa mengirim satu-satu ke goroutine
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast()
	//}()

	group.Wait()
	fmt.Println("Selesai")
}


