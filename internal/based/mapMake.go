package based

import "fmt"

/*
Напиши код, который создает карту (map) с использованием функции make,
добавляет в неё ключ "one" с значением 1
и выводит значение по этому ключу.
*/

func main() {
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m["one"])
}
