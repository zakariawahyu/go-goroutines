package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
*
Mutex (Mutual Exclusion)
- Untuk mengatasi masalah race condition, di golang terdapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika locking terhadap mutex, maka tidak ada yang bisa melakukan locking sampai kita melakukan unlock
- Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut melakukan unlock,
baru goroutine selanjutnya diperbolehkan melakukan lock lagi
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/
func TestMutex(t *testing.T) {
	x := 0
	mutex := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock() // kunci dulu jangan ada go routine yang mengakses bersamaan
				x += 1
				mutex.Unlock() // setelah unlock baru go routine lainnya bisa lock
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter : ", x)
}

// Novalagung version
type Increment struct {
	sync.Mutex
	val int
}

func (c *Increment) Add(int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *Increment) Value() int {
	return c.val
}

func TestMutex2(t *testing.T) {
	var count Increment
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				count.Add(1)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter : ", count.Value())
}

/**
sync.RWMutex (Read Write Mutex)
- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga untuk membaca data sekaligus
- Kita sebenernya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebuatan antara proses read dan writenya
- Di golang telah disediakan struct bernama RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock, lock untuk Read dan lock untuk Write
*/

type BankAcount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAcount) addBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAcount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}
func TestReadWriteMutex(t *testing.T) {
	account := BankAcount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}

	time.Sleep(7 * time.Second)
	fmt.Println("Final Balance : ", account.getBalance())
}
