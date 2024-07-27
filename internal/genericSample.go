package internal

import "fmt"

func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(123)
	Print("hello")
	Print(3.14)
}
