package go_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/**
Pengenalan Channel
- Channel adalah tempat komunikasi secara synchronus yang bisa dilakukan oleh go routine
- Di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang berbeda
- Saat melakukan pengiriman data ke channel, goroutine akan terblock sampai ada yang menerima data tersebut
- Maka dari itu, channel disebut sebagai alat komunikasi synchronus(blocking)
- Channel cocok sekali sebagai alternatif seperti mekanisme async wait yang terdapat di beberapa bahasa pemograman lainnya

Karakteristik Channel
- Secara default, channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi, harus menunggu data yang ada di channel diambil
- Channel hanya bisa meenerima satu jenis tipe data
- Channel bisa diambil lebih dari satu go routine
- Channel harus di close jika tidak digunakan, atau biasanya akan menyebabkan memory leak jika tidak di close

Mengirim dan menerima data dari Channel
- Seperti yang suudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data kita bisa menggunakan kode channel <- data
- Sedangkan untuk menerima data, bisa gunakan kode data <- channel
- Jika selesai jangan lupa untuk menutup channel dengan menggunakan function close
 */
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// mengirim data ke channel
		channel <- "Zakaria Wahyu Nur Utomo"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	//menerima data
	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// atau bisa dikirim langsung sebagai parameter
	//fmt.Print(<- channel)
}

/**
Channel sebagai parameter
- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahwal di Golang by default, parameter adalah pass by value artinya value akan diduplokasi lalu dikirim ke function parameter,
sehingga jika kita ingin mengirim data asli kita bisa gunakan pointer (pass by reference)
- Berbeda dengan channel, kita tidak perlu melakukan hal itu
 */
func GiveMeRespon(channel chan string)  {
	time.Sleep(2 * time.Second)

	channel <- "Zakaria Wahyu Nur Utomo"
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeRespon(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	close(channel)
}

/**
Channel In dan Out
- Saat kita mengirim channel sebagai parameter, isi fucntion tersbut bisa mengirim dan menerima data dari channel tersebut
- Terkadang kita ingin memberi tahu terhadap function, misal bahawa channel tersebut hanya digunakan untuk mengirim data atau hanya untuk menerima data saja
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerioma data)
 */
func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)

	channel <- "Zakaria Wahyu Nur Utomo"
}

func OnlyOut(channel <-chan string)  {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}

/**
Buffered Channel
- Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
- Artinya jika kita ingin menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengambil
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- Untungnya ada Buffered Channel yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel

Buffer Capacity
- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita ingin mengirim data ke-6, maka akan diminta untuk menunggu sampai buffer ada yang kosong
- Ini cocok sekali ketika memang go routine yang menerima data lebih lambat dari yang mengirim data
 */
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	// penulisan tanpa go routine
	channel <- "Zakaria"
	channel <- "Wahyu"

	fmt.Println(<- channel)
	fmt.Println(<- channel)

	// penulisan dengan goroutine
	go func() {
		channel <- "Zakaria"
		channel <- "Wahyu"
	}()

	go func() {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	time.Sleep(2 * time.Second)
	close(channel)
}

/**
Range Channel
- Kadang-kadang ada kasus sebuah channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan channel tersbut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
- Keika sebuah channel di close(), maka secara otomatis perulangan tersbeut akan berhenti
- ini lebih sederhana dari pada kita melakukan pengecekan channel secara manual
 */
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke : " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// perulangan for range untuk mengambil channel
	for data := range channel{
		fmt.Println("Menerima data : ", data)
	}
	fmt.Println("Selesai")
}

/**
Select Channel
- Kadang ada kasus dimana kita membuat beberapa channel dan menjalankan beberapa goroutine
- Lalu kita ingin mendapatkan data daru semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Golang
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika datang secara bersamaan di beberapa channel,
maka akan dipilih secara random
 */
func TestSelectCannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespon(channel1)
	go GiveMeRespon(channel2)

	counter := 0
	for {
		select {
			case data := <- channel1:
				fmt.Println("Data dari channel 1 : ", data)
				counter++
			case data := <-channel2:
				fmt.Println("Data dari channel 2 : ", data)
				counter++
		}
		if counter == 2 {
			break
		}
	}
}

/**
Default Select
- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
- Maka kita akan menunggu sampai data ada
- Kaadng mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select , kita bisa menambahkan default dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya
 */
func TestDefaultSelectCannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespon(channel1)
	go GiveMeRespon(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1 : ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 : ", data)
			counter++
		default:
			fmt.Println("Menungu data..")
		}
		if counter == 2 {
			break
		}
	}
}