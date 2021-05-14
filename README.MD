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


