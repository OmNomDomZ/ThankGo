package SomeTasks

//
//import (
//	"fmt"
//	"strings"
//	"unicode"
//)
//
//func main() {
//	phrase := readString()
//
//	var words []string
//
//	words = strings.Fields(phrase)
//
//	abbr := make([]string, 0)
//
//	for _, word := range words {
//		runes := []rune(word)
//		if unicode.IsLetter(runes[0]) || unicode.Is(unicode.Cyrillic, runes[0]) {
//			abbr = append(abbr, strings.ToUpper(string(runes[0])))
//		}
//	}
//
//	fmt.Println(strings.Join(abbr, ""))
//}
//
////func readString() string {
////	rdr := bufio.NewReader(os.Stdin)
////	str, _ := rdr.ReadString('\n')
////
////	return str
////}
