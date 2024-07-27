package based

import "fmt"

func main() {
	str := "Hello"
	runes := []rune(str)
	runes[0] = 'h'
	newStr := string(runes)
	fmt.Println(newStr) // Output: "hello"
}
