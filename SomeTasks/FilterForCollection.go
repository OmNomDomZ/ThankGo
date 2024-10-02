package SomeTasks

//import (
//	"fmt"
//)
//
//func filter(predicate func(int) bool, iterable []int) []int {
//	res := make([]int, 0)
//
//	for _, i := range iterable {
//		if predicate(i) {
//			res = append(res, i)
//		}
//	}
//
//	return res
//}
//
//func main() {
//	src := readInput()
//
//	res := filter(func(i int) bool { return src[i]%2 == 0 }, src)
//
//	fmt.Println(res)
//}

//func readInput() []int {
//	var nums []int
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//	for scanner.Scan() {
//		num, _ := strconv.Atoi(scanner.Text())
//		nums = append(nums, num)
//	}
//	return nums
//}
