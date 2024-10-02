package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say печатает фразу от имени обработчика
func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

// начало решения

// makePool создает пул на n обработчиков
// возвращает функции handle и wait
func makePool(n int, handler func(int, string)) (func(string), func()) {
	// создайте пул на n обработчиков
	// используйте для канала имя pool и тип chan int
	// определите функции handle() и wait()

	// handle() выбирает токен из пула
	// и обрабатывает переданную фразу через handler()

	// wait() дожидается, пока все токены вернутся в пул

	pool := make(chan int, n)

	for i := 0; i < n; i++ {
		pool <- i
	}

	handle := func(phrase string) {
		id := <-pool
		go func() {
			handler(id, phrase)
			pool <- id
		}()
	}

	wait := func() {
		for i := 0; i < n; i++ {
			<-pool
		}
		close(pool)
	}

	return handle, wait
}

// конец решения

func main() {
	phrases := []string{
		"3",
		"1",
		"2",
	}

	handle, wait := makePool(4, say)
	for _, phrase := range phrases {
		handle(phrase)
	}
	wait()
}
