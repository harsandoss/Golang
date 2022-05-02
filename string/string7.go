//To find the Length of the String
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	str = "Hello Golang"
	fmt.Println(str)
	length := utf8.RuneCountInString(str)
	fmt.Println("The Length of a given string =", length)
}
