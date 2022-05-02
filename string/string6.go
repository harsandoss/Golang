package main

import "fmt"

func main() {
	var str string

	str = "I like Golang"
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		strFirstChar := str[i]
		fmt.Printf("The Character in the String[%d] = %c\n", i, strFirstChar)
	}
}
