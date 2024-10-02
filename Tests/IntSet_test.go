package Tests

// не удаляйте импорты, они используются при проверке
import (
	"fmt"
	"math/rand"
	"testing"
)

// IntSet реализует множество целых чисел
// (элементы множества уникальны).
type IntSet struct {
	m *map[int]bool
}

// MakeIntSet создает пустое множество.
func MakeIntSet() IntSet {
	m := make(map[int]bool)
	return IntSet{&m}
}

// Contains проверяет, содержится ли элемент в множестве.
func (s IntSet) Contains(elem int) bool {
	if _, ok := (*s.m)[elem]; !ok {
		return false
	}
	return true
}

// Add добавляет элемент в множество.
// Возвращает true, если элемент добавлен,
// иначе false (если элемент уже содержится в множестве).
func (s IntSet) Add(elem int) bool {
	if (*s.m)[elem] == true {
		return false
	}
	(*s.m)[elem] = true
	return true
}

func BenchmarkIntSet(b *testing.B) {
	for _, size := range []int{1, 10, 100, 1000, 10000} {
		set := randomSet(size)
		name := fmt.Sprintf("Contains-%d", size)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				elem := rand.Intn(100000)
				set.Contains(elem)
			}
		})

	}
}

var rnd = rand.New(rand.NewSource(42))

func randomSet(size int) IntSet {
	set := MakeIntSet()
	for i := 0; i < size; i++ {
		n := rnd.Intn(100000)
		set.Add(n)
	}
	return set
}
