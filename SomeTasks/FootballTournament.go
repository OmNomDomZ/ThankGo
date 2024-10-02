package SomeTasks

//
//import (
//	"fmt"
//	"sort"
//	"strings"
//)
//
//type result byte
//
//const (
//	win  result = 'W'
//	draw result = 'D'
//	loss result = 'L'
//)
//
//type team byte
//
//type match struct {
//	first  team
//	second team
//	result result
//}
//
//type rating map[team]int
//
//// tournament представляет турнир
//type tournament []match
//
//// calcRating считает и возвращает рейтинг турнира
//func (t *tournament) calcRating() rating {
//	r := make(rating)
//
//	for _, m := range *t {
//		if m.result == draw {
//			r[m.first]++
//			r[m.second]++
//		} else if m.result == win {
//			r[m.first] += 3
//			r[m.second] += 0
//		} else if m.result == loss {
//			r[m.first] += 0
//			r[m.second] += 3
//		}
//	}
//
//	return r
//}
//
//func main() {
//	src := readString()
//	trn := parseTournament(src)
//	rt := trn.calcRating()
//	rt.print()
//}
//
//// readString считывает и возвращает строку из os.Stdin
////func readString() string {
////	rdr := bufio.NewReader(os.Stdin)
////	str, err := rdr.ReadString('\n')
////	if err != nil && err != io.EOF {
////		log.Fatal(err)
////	}
////	return str
////}
//
//// parseTournament парсит турнир из исходной строки
//func parseTournament(s string) tournament {
//	descriptions := strings.Split(s, " ")
//	trn := tournament{}
//	for _, descr := range descriptions {
//		m := parseMatch(descr)
//		trn.addMatch(m)
//	}
//	return trn
//}
//
//// parseMatch парсит матч из фрагмента исходной строки
//func parseMatch(s string) match {
//	return match{
//		first:  team(s[0]),
//		second: team(s[1]),
//		result: result(s[2]),
//	}
//}
//
//// addMatch добавляет матч к турниру
//func (t *tournament) addMatch(m match) {
//	*t = append(*t, m)
//}
//
//// print печатает результаты турнира
//// в формате Aw Bx Cy Dz
//func (r *rating) print() {
//	var parts []string
//	for team, score := range *r {
//		part := fmt.Sprintf("%c%d", team, score)
//		parts = append(parts, part)
//	}
//	sort.Strings(parts)
//	fmt.Println(strings.Join(parts, " "))
//}
