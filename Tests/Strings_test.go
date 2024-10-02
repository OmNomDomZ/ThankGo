package Tests

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func check(r rune) bool {
	if ("a" <= string(r) && string(r) <= "z") || unicode.IsNumber(r) || string(r) == "-" {
		return true
	}
	return false
}

func slugify(src string) string {

	src = strings.ToLower(src)
	l := len(src)

	prevWasWord := true

	for i := 0; i < l; i++ {
		if check(rune(src[i])) {
			prevWasWord = true
			continue
		} else {
			letter := string(src[i])
			fmt.Println(letter)
			if prevWasWord {
				src = strings.Replace(src, string(rune(src[i])), " ", 1)
				prevWasWord = false
			} else {
				src = strings.Replace(src, string(rune(src[i])), "", 1)
			}
			l = len(src)
		}
	}

	src = strings.TrimLeft(src, " ")
	src = strings.TrimRight(src, " ")
	src = strings.ReplaceAll(src, " ", "-")

	return src
}

// конец решения

func Test(t *testing.T) {
	const phrase = "Arrays, slices (and strings): The mechanics of 'append'"
	const want = "arrays-slices-and-strings-the-mechanics-of-append"
	got := slugify(phrase)
	if got != want {
		t.Errorf("%s: got %#v, want %#v", phrase, got, want)
	}
}
