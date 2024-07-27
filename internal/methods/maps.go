package methods

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	fmt.Println(m["one"]) // 1

	delete(m, "one")
	fmt.Println(m["one"]) // 0, так как ключ "one" больше не существует
}
