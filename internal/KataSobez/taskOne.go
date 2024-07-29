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

func newCalc(numA, numB int) int {
	var sumAB int

	switch i {
	case 0:

		sumAB = numA * numB
	case 1:

		sumAB = numA / numB
	case 2:
		sumAB = numA - numB
	default:
		sumAB = numA + numB
	}
}

func main() {
	var a, b, sumAB int
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
