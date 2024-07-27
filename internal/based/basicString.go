package based

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"

	// 1. Выводит длину строки.
	fmt.Println(len(str))

	// 2. Разделяет строку по запятой.
	parts := strings.Split(str, ",")
	fmt.Println(parts)

	// 3. Проверяет, содержит ли строка подстроку "World".
	contains := strings.Contains(str, "World")
	fmt.Println(contains)
}
