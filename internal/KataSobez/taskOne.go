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

//func newCalc(numA, numB, operation int) int {
//	var result int
//
//	switch operation {
//	case 0: // Умножение
//		result = numA * numB
//	case 1: // Деление
//		if numB != 0 {
//			result = numA / numB
//		} else {
//			fmt.Println("Ошибка: Деление на ноль")
//			return 0 // Возвращаем 0 в случае ошибки
//		}
//	case 2: // Вычитание
//		result = numA - numB
//	default: // Сложение
//		result = numA + numB
//	}
//	return result
//}

func main() {
	var a, b int
	//var operation int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	res := compare(a, b)
	fmt.Println(res)

	//fmt.Println("Выберите операцию (0 - умножение, 1 - деление, 2 - вычитание, любое другое число - сложение):")
	//fmt.Scanln(&operation)
	//
	//calcResult := newCalc(a, b, operation)
	//fmt.Printf("Результат операции: %d\n", calcResult)
}

// Необходимо написать функцию, которая принимает два числа в качестве аргументов и
// возвращает строку с информацией о том, какое число больше - первое, второе или они равны.
//
// Затем в main считать из консоли два числа от пользователя
// и вызвать функцию с этими значениями и вывести результат в консоль.
