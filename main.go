package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	fmt.Println("input word is \"" + text + "\"")
}
