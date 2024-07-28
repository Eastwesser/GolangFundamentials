package StepikGo

import "fmt"

func main() {
	var currentNum int

	// Считываем число с клавиатуры
	fmt.Scanln(&currentNum)

	mult := 1

	for i := 1; i <= currentNum; i++ {
		mult *= i
	}
	fmt.Println(mult)
}
