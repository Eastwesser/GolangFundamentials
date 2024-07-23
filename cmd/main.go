package main

import (
	"GolangFundamentials/internal"
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
	totalPoints := internal.CalculatePoints(matchResults)

	// Вывод общего количества очков
	fmt.Println(totalPoints)
}
