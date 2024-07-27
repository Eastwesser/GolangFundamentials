package methods

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice) // [1 2 3 4 5 6]

	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println(newSlice) // [1 2 3 4 5 6]
}
