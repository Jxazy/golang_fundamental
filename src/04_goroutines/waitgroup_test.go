package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// waitgrup digunakan untuk menunggu seluruh proses goroutines selesai
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Complete")
}
