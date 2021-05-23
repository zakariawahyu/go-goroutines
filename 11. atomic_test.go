package go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/**
Atomic
- Golang memiliki package yang bernama sync/Atomic
- Atomic adalah package yang digunakan untuk menggunakan data primitive secara aman pada prosess cocurrent
- Contoh data primitiv : Int, Float, String, Boolean
- Conohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di coounter.
Hal ini sebenarnya bisa menggunakan aatomic package
- Ada banyak func di atomic package, kita bisa eksplore sendiri di dokumentasinya
- golang.org/pkg/sync/atomic
 */

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", x)
}
