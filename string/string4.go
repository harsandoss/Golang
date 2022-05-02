// To convert byte to String
package main

import "fmt"

func main() {
	byteArray := []byte{72, 101, 108, 108, 111, 32, 71, 111, 108, 97, 110, 103}
	var a string
	a = string(byteArray)
	fmt.Println(a)

}
