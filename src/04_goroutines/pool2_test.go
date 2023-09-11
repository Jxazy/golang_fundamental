package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func WaitPool(group *sync.WaitGroup) {
	defer group.Done()

	pool := sync.Pool{
		New: func() interface{} {
			return "Kosong"
		},
	}
	pool.Put("Fajar")
	pool.Put("Ahmad")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
}

func TestWaitPool(t *testing.T) {
	group := &sync.WaitGroup{}

	go WaitPool(group)

	group.Wait()
	fmt.Println("Complete")
}
