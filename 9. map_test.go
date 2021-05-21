package go_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
sync.Map
- Golang memiliki sebuah struct bernama sync.Map{}
- Map ini mirip sekali dengan golang map, namun yang membedakan, map ini aman digunakan untuk menggunakan concurrent menggunakan goroutine

Ada beberapa function yang bisa kita gunakan didalam Map:
- Store(key, value)			= Untuk menyimpan data ke Map
- Load(key)					= Untuk mengambil data di Map menggunakan key
- Delete(key)				= Untuk menghapus data di Map menggunakan key
- Range(func(key. value))	= Dilakukan untuk melakukan iterasi seluruh data Map
 */

func addToMap(data *sync.Map ,value int, group *sync.WaitGroup)  {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)

	time.Sleep(1 * time.Nanosecond)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go addToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
