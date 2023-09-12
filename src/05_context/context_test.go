package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "f")
	contextG := context.WithValue(contextC, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Hanya dapat mengambil value ke atas
	fmt.Println(contextF.Value("f")) // bisa, karena memiliki value
	fmt.Println(contextF.Value("c")) // bisa, karena mengambil value dari parent
	fmt.Println(contextF.Value("b")) // tidak bisa, karena berbeda parent
	fmt.Println(contextA.Value("b")) // tidak bisa, karena tidak mengambil value dari child
}

func GoroutineLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestLeak(t *testing.T) {
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	destination := GoroutineLeak()

	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}

	}

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}

func GoroutineLeakFix(ctx context.Context) chan int {

	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}

		}

	}()

	return destination
}

func TestContextCancel(t *testing.T) {

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := GoroutineLeakFix(ctx)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter :", n)
		if n == 10 {
			break
		}

	}

	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}

func TestContextTimeout(t *testing.T) {
	// waktu duration menggunakan cancel
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := GoroutineLeakFix(ctx)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter :", n)

	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}

func TestContextDeadline(t *testing.T) {
	// waktu fix menggunakan deadline
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := GoroutineLeakFix(ctx)

	fmt.Println("Total Goroutine :", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter :", n)

	}

	time.Sleep(5 * time.Millisecond)
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}
