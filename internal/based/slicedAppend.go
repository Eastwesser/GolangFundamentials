package based

import "fmt"

/*
Напиши код, который создает слайс из трех чисел [1, 2, 3],
добавляет в него числа 4, 5, 6 с помощью append
и выводит конечный слайс.
*/

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)
}
