**GO GOROUTINES**
- Concurenncy & Parallel Programming
- Goroutines
- Channel
- Buffered Channel
- Mutex
- WaitGroup
- Atomic
- Ticker

**Pengenalan Concurrency dan Parallel Programming**

**Pengenalan Paraller Programming :**
- Saat ini kita hidup di era multi-core, dimana jarang sekali kita menggunakan prosessor yang single core
- Semakin canggih perangkat keras, maka software pun akan mengikuti, dimana sekarang kita bisa dengan mudah membuat proses parallel di aplikasi
- Parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi lebih kecil dan dijalankan secara bersamaan dan pada waktu yang bersamaan juga

**Contoh Parallel :**
- Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (Office, Editor, Browser, Spotify dll)
- Beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanan masing-masing
- Antrian di bank, dimana tiap teller melayani nasabahnya masing-masing

**Process vs Thread**

| Process                                       | Thread                                                      |
|-----------------------------------------------|-------------------------------------------------------------|
| Process adalah sebuah eksekusi program        | Thread adalah segmen dari process                           |
| Process mengkonsumsi memory besar             | Thread menggunakan memory kecil                             |
| Process saling terisolasi dengan process lain | Thread bisa saling berhubungan jika dalam process yang sama |
| Process lama untuk dijalankan dan dihentikan  | Thread cepat untuk dijalankan dan dihentikan                |

**Parallel vs Concurrency**
- Berbeda dengan parallel (menjalankan beberapa pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara bergantian
- Dalam parallel kita bisa membutuhkan banyak Thread, sedangkan dalam Concurrency kita hanya membutuhkan sedikit Thread

**Contoh Concurrency**
- Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi dst. Tetapi kita tidak bisa pada saat bersamaan minum, makan dan ngobrol. hanya bisa melakukan satu hal pada waktu, namun bisa bergantian kapanpun kita mau.

**CPU-bound**
- Banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya. Algoritma jenis ini biasanya tergantung dengan kecepatan CPU
- Contoh yang paling populer adalah Machine Learning, oleh karena itu sekarang banyak sekali tekonologi Machine Learning yang banyak menggunakan GPU karena memiliki core yang lebih banyak dibandingkan CPU
- Jenis algoritma seperti ini tidak bisa ada benefit jika menggunakan Cocurrency Programming, namun bisa dibantu dengan implementasi Parallel Prgramming

**I/O-bound**
- I/O-bound dimana biasanya algoritma sangat tergantung dengan kecepatan input output devices yang digunakan
- Contoh aplikasi seperti membaca file, database dll
- Kebanyakan saat ini, biasanya kita akan membuat jenis aplikasi ini untuk backend
- Aplikasi jenis I/O-bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Cocurrency Programming
- Bayangkan jika kita membaca file dari database dan Thread harus menunggu 1 detik untuk mendapatkan respon dari database, padahal
waktu 1 detik itu jika menggunakan Cocurrency Programming, bisa digunakan untuk melakukan hal lain
  
**Pengenalan Goroutines**
- Goroutines adalah sebuah thread ringan yang dikelola oleh GO Rutines
- Ukuran Goroutines sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan thread yang bisa sampai 1MB atau 1000kb
- Namun tidak seperti thread yang berjalan parallel, goroutines berjalan secara concurennt

**Cara Kerja Goroutines**
- Sebenarnya, Goroutines dijalankan oleh GO Scheduler di dalam thread, dimana jumlah threadnya sebanyak GOMAXPROCS (biasanya sejumlah core CPU)
- Jadi sebenernya, tidak bisa dibilang Goroutines itu sebagai pengganti thread, karena Goroutines itu sendiri berjalan diatas Thread
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler


  



