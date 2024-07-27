package KataSobez

import (
	"fmt"
)

func compare(a, b int) string {
	var stringBigger string

	if a < b {
		stringBigger = "Число Б больше"
	} else if a > b {
		stringBigger = "Число А больше"
	} else {
		stringBigger = "Числа равны"
	}
	return stringBigger
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	res := (compare(a, b))
	fmt.Println(res)
}

// Необходимо написать функцию, которая принимает два числа в качестве аргументов и
// возвращает строку с информацией о том, какое число больше - первое, второе или они равны.
//
// Затем в main считать из консоли два числа от пользователя
// и вызвать функцию с этими значениями и вывести результат в консоль.
