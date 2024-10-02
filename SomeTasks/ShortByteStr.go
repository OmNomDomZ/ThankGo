package SomeTasks

import "fmt"

func main() {
	var text string
	var width int

	_, err := fmt.Scan(&text, &width)
	if err != nil {
		fmt.Println(err)
	}

	if len(text) <= width {
		fmt.Println(text)
		return
	}

	s := []byte(text)
	slice := s[:width]

	res := string(slice) + "..."

	fmt.Println(res)
}
