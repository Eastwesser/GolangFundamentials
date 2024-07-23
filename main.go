package GolangFundamentials

import (
	"fmt"
)

func main() {
	// Срез с результатами матчей
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	// Считываем результат последней игры
	var lastGame string
	fmt.Scanln(&lastGame)

	// Добавляем результат в срез
	matchResults := append(results, lastGame)

	// Считаем очки
	maxScore := 0
	for _, result := range matchResults {
		switch result {
		case "w":
			maxScore += 3
		case "l":
			maxScore += 0
		case "d":
			maxScore += 1
		}
	}
	// Вывод общего количества очков
	fmt.Println(maxScore)
}
