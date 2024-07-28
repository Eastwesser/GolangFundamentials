package StepikGo

import "fmt"

func main() {
	var currentSum int

	// Считываем сумму денег с клавиатуры
	fmt.Scanln(&currentSum)

	if currentSum > 1000 {
		fmt.Println("Apple")
	} else if currentSum <= 1000 && currentSum >= 500 {
		fmt.Println("Samsung")
	} else if currentSum < 500 {
		fmt.Println("Nokia с фонариком")
	}
}
