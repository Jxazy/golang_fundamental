package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Sprintln("Total Goroutine:", totalGoroutine)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

}

func TestChangeGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Sprintln("Total Goroutine:", totalGoroutine)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

}
