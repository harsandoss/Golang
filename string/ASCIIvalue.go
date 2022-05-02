package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any string to find ASCII values =")
	strData, _ := reader.ReadString('\n')

	for i := 0; i < len(strData); i++ {
		fmt.Printf("The ASCII Value of %c = %d\n", strData[i], strData[i])
	}
	fmt.Println()
}
