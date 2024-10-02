package main

import (
	"go.uber.org/goleak"
	"testing"
)

func TestMain(m *testing.M) {
	// Проверка утечек до и после выполнения всех тестов
	goleak.VerifyTestMain(m)
}

func TestPipeline(t *testing.T) {
	// Создаем канал для отмены
	cancel := make(chan struct{})
	defer close(cancel)

	// Вызываем функции, чтобы протестировать их работу
	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)

	// Проверяем, что результат не вызывает утечек и корректно завершается
	print(cancel, c4, 5)

	// Здесь происходит автоматическая проверка на утечки при завершении тестов
}
