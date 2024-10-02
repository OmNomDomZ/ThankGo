package Tests

// не удаляйте импорты, они используются при проверке
import (
	"strings"
)

// Words работает со словами в строке.
type Words struct {
	m map[string]int
}

// MakeWords создает новый экземпляр Words.
func MakeWords(s string) Words {
	m := make(map[string]int)
	for ind, str := range strings.Fields(s) {
		if _, ok := m[str]; ok {
			continue
		}
		m[str] = ind
	}

	return Words{m}
}

// Index возвращает индекс первого вхождения слова в строке,
// или -1, если слово не найдено.
func (w Words) Index(word string) int {
	if val, ok := w.m[word]; ok {
		return val
	}
	return -1
}
