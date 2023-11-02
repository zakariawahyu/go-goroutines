package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Deadlock
- Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi adalah Deadlock
- Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehinggs tidak ada satupun goroutine yang bisa jalan
*/

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Tranfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1 : ", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2 : ", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Zakaria",
		Balance: 100000,
	}

	user2 := UserBalance{
		Name:    "Wahyu",
		Balance: 100000,
	}

	// dari transaksi pertama, seharusnya setelah user1 lock maka akan lock user2
	// tetapi pada transaksi pertama ini, user2 sudah keburu di lock pada transaksi yang ke dua
	go Tranfer(&user1, &user2, 10000)
	// pada transaksi kedua, seharusnya setelah user2 lock maka akan lock user1
	//tetapi pada transaksi kedua ini, user1 keburu di lock pada transaksi yang pertama
	// jadi antara transaksi pertama dan kedua saling tunggu dan terjadi deadlock
	go Tranfer(&user2, &user1, 20000)

	time.Sleep(3 * time.Second)

	// ketika sleep sudah habis maka akan di print dengan apa adanya
	fmt.Println("User : ", user1.Name, "Balance : ", user1.Balance)
	fmt.Println("User : ", user2.Name, "Balance : ", user2.Balance)
}
