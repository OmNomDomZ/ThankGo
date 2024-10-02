package SomeTasks

import (
	"fmt"
	"math/rand"
)

func shuffle(nums []int) {
	rnd := rand.New(rand.NewSource(42))

	rnd.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

//func init() {
//rnd := rand.New(rand.NewSource(42))
//rand.Seed(42)
//}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
}

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
