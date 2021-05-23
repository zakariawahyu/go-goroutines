package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
Timer
- Timer adalah representasi satu kejadian
- Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat timer kita bisa menggunakan time.NewTimer(duration)
 */

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Waktu runnning pertama : ", time.Now())

	// setelah menunggu 5 detik maka akan kirim ke channel
	// biasanya digunakan untuk jobs yang akan datang
	afterTimer := <- timer.C
	fmt.Println("Waktu setelah timer : ", afterTimer)
}

/**
time.After()
- Kadang kita hanya butuh chanel nya saja, tidak membutuhkan data Timer nya
- Untuk melakukan itu kita bisa menggunakan function time.After(duration)
 */

func TestAfterTimer(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Waktu runnning pertama : ", time.Now())

	afterTimer := <- channel
	fmt.Println("Waktu setelah timer : ", afterTimer)
}

/**
time.AfterFunc()
- Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
- Kita bisa memanfaatkan Timer dengan menggunakan function time.AterFunc()
- Kita tidak perlu lagi menggunakan channelnya, cukup kirimkan function yang akan dipanggil ketika Timer mengirim kejadian
 */

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println("Waktu setelah timer : ", time.Now())
		fmt.Println("Excude after 5 second")
		group.Done()
	})
	fmt.Println("Waktu runnning pertama : ", time.Now())
	group.Wait()
}