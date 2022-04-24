package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter the number:")
	fmt.Scan(&a)
	num := a
	for num >= 10 {
		num = num / 10
	}
	fmt.Printf("First digit of %d is %d", a, num)
}
