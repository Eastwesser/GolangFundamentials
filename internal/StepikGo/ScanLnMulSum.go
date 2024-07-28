package StepikGo

import "fmt"

func main() {
	var a, b, c int

	// fmt.Println("Ввод значений a, b, c:")
	fmt.Scanln(&a, &b, &c)

	// Вывод суммы и произведения
	fmt.Println(a + b + c)
	fmt.Println(a * b * c)
}
