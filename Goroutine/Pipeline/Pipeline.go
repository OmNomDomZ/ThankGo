package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// начало решения

type WordPair struct {
	original string
	reversed string
}

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan WordPair {
	out := make(chan WordPair)
	go func() {
		defer close(out)
		for {
			select {
			case out <- WordPair{original: randomWord(5)}:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan WordPair) <-chan WordPair {
	out := make(chan WordPair)
	go func() {
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				m := make(map[int32]int)
				unique := true
				for _, letter := range word.original {
					m[letter]++
					if m[letter] == 2 {
						unique = false
						break
					}
				}
				if unique {
					select {
					case out <- word:
					case <-cancel:
						return
					}
				}
			case <-cancel:
				return
			}
		}
	}()
	return out
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan WordPair) <-chan WordPair {
	out := make(chan WordPair)
	go func() {
		defer close(out)
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				reversedRunes := []rune(word.original)
				for i, j := 0, len(reversedRunes)-1; i < j; i, j = i+1, j-1 {
					reversedRunes[i], reversedRunes[j] = reversedRunes[j], reversedRunes[i]
				}
				select {
				case out <- WordPair{word.original, string(reversedRunes)}:
				case <-cancel:
					return
				}

			case <-cancel:
				return
			}
		}
	}()
	return out
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan WordPair) <-chan WordPair {
	out := make(chan WordPair)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-cancel:
				return

			case pair, ok := <-c1:
				if !ok {
					return
				}
				select {
				case out <- pair:
				case <-cancel:
					return
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-cancel:
				return

			case pair, ok := <-c2:
				if !ok {
					return
				}
				select {
				case out <- pair:
				case <-cancel:
					return
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// печатает первые n результатов
// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan WordPair, n int) {
	for i := 0; i < n; i++ {
		select {
		case word, ok := <-in:
			if !ok {
				return
			}
			fmt.Printf("%v -> %v\n", word.original, word.reversed)
		case <-cancel:
			return
		}
	}
	return
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
