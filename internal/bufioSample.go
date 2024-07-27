package internal

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You entered:", line)
	}
}
