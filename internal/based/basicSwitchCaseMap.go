package based

import "fmt"

/*
1. Использует switch case для проверки типа переменной value.
2. Создает карту (map), добавляет в неё ключ "key" с значением 10 и выводит значение по этому ключу.
*/

func main() {
	var a interface{} = "example"

	switch v := a.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}

	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"])
}
