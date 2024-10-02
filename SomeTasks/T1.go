package SomeTasks

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	var text string
//
//	scanner := bufio.NewScanner(os.Stdin)
//
//	scanner.Scan()
//	text = scanner.Text()
//
//	res := getSlice(text)
//
//	fmt.Println(res)
//}
//
//func getSlice(str string) []int {
//	res := make([]int, 0)
//
//	parts := strings.Split(str, ",")
//
//	for _, part := range parts {
//		if strings.Contains(part, "-") {
//			bounds := strings.Split(part, "-")
//			start, _ := strconv.Atoi(bounds[0])
//			end, _ := strconv.Atoi(bounds[1])
//
//			for i := start; i <= end; i++ {
//				res = append(res, i)
//			}
//		} else {
//			num, _ := strconv.Atoi(part)
//			res = append(res, num)
//		}
//	}
//	return res
//}
