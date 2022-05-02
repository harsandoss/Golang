package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the 1st String:")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSuffix(str1, "\n")

	fmt.Print("Enter the 2st String:")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSuffix(str2, "\n")

	str3 := str1 + " " + str2
	fmt.Print(str3)
}
