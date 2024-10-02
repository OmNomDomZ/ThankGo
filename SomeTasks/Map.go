package SomeTasks

//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	var i int
//	_, err := fmt.Scan(&i)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	counter := make(map[int]int)
//
//	for i != 0 {
//		counter[i%10]++
//		i /= 10
//	}
//
//	printCounter(counter)
//
//}
//
//func printCounter(counter map[int]int) {
//	d := make([]int, 0)
//	for ind := range counter {
//		d = append(d, ind)
//	}
//
//	sort.Ints(d)
//
//	for _, ind := range d {
//		fmt.Printf("%v:%v ", ind, counter[ind])
//	}
//
//}
