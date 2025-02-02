package based

import "fmt" // Импортируем пакет fmt для работы с форматированным вводом/выводом.

func main() { // Основная функция, с которой начинается выполнение программы.
	var input int                // Объявляем переменную input типа int для хранения введенного числа.
	fmt.Print("Введите число: ") // Выводим сообщение на экран с просьбой ввести число.
	fmt.Scanln(&input)           // Считываем введенное пользователем число и сохраняем его в переменную input.

	switch input { // Начинаем оператор switch для проверки значения переменной input.
	case 0: // Если input равен 0,
		fmt.Println("ноль") // выводим на экран "ноль".
	default: // Если input не равен 0,
		fmt.Println("Ошибка: введите число от 0 до 9") // выводим сообщение об ошибке.
	}
}
