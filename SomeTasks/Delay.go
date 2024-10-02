package SomeTasks

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancel := func() {
	}

	return cancel
}

// конец решения

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
