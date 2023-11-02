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

func RunAsynchronus(group *sync.WaitGroup, count int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello : ", count)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group, i)
	}

	group.Wait()
	fmt.Println("Done")
}

// Novalagung version
/*
- Kegunaan waitgroup adalah untuk sinkronisasi goroutine. Berbeda dengan channel, waitgroup memang dirancang khusus untuk maintain goroutine
- Channel untuk keperluan sharing data antar goroutine, sedangkan sync.WaitGroup untuk sinkronisasi goroutine

- WaitGroup digunakan untuk menunggu goroutine. Cara penggunaannya sangat mudah, tinggal masukan jumlah goroutine yang dieksekusi, sebagai parameter method Add()
- Dan pada akhir tiap-tiap goroutine, panggil method Done()
- Juga, pada baris kode setelah eksekusi goroutine, panggil method Wait()
*/

func TestWaitGroup2(t *testing.T) {
	doPrint := func(wg *sync.WaitGroup, message string) {
		defer wg.Done()
		fmt.Println(message)
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)
		wg.Add(1)
		go doPrint(wg, data)
	}

	wg.Wait()
}

/*
- Kode di atas merupakan contoh penerapan sync.WaitGroup untuk pengelolahan goroutine. Fungsi doPrint() akan dijalankan sebagai goroutine, dengan tugas menampilkan isi variabel message
- Variabel wg disiapkan bertipe sync.WaitGroup, dibuat untuk sinkronisasi goroutines yang dijalankan.
- Di tiap perulangan statement wg.Add(1) dipanggil. Kode tersebut akan memberikan informasi kepada wg bahwa jumlah goroutine yang sedang di proses ditambah 1 (karena dipanggil 5 kali, maka wg akan sadar bahwa terdapat 5 buah goroutine sedang berjalan).
- Di baris selanjutnya, fungsi doPrint() dieksekusi sebagai goroutine. Di dalam fungsi tersebut, sebuah method bernama Done() dipanggil. Method ini digunakan untuk memberikan informasi kepada wg bahwa goroutine di mana method itu dipanggil sudah selesai. Sejumlah 5 buah goroutine dijalankan, maka method tersebut harus dipanggil 5 kali.
- Statement wg.Wait() bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah 5 goroutine selesai. Jika Add(1) dipanggil 5 kali, maka Done() juga harus dipanggil 5 kali
*/
