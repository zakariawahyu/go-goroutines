package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/**
time.Ticker()
- Ticker adalah representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat ticker, kita bisa menggunakan time.NewTikcker(duration)
- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
 */

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	// untuk stop kita buat goroutine dan ticker.stop()
	// akan berhenti setelah 5 detik
	// hasilnya akan error deadlock karena setelah stop tidak ada data yang dikirim lagi padahal dibawah tetap melakukan for range
	// idealnya jangan melakukan for range, tapi menggunakan select
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

/**
time.Tick()
- Kadang kita tidak butuh data Tickernya, kita hanya butuh channel saja
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
mengembalikan ticker, hanya channel timernya saja
 */

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}
