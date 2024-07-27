package methods

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"
	fmt.Println(strings.Contains(str, "World")) // true
	fmt.Println(strings.Split(str, ","))        // ["Hello" " World!"]
}
