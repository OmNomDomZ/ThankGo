package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения

	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его

	// конец решения
	res := make([]any, len(funcs))
	done := make(chan any, len(funcs))

	for i, f := range funcs {
		go func(i int, f func() any) {
			res[i] = f()
			done <- struct{}{}
		}(i, f)
	}

	for i := 0; i < len(funcs); i++ {
		<-done
	}

	return res

}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(3), squared(4), squared(1)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
