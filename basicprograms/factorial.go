package main

import "fmt"

func main() {
	var f int
	fmt.Print("Enter the number: ")
	fmt.Scan(&f)
	num := 1
	for i := 1; i <= f; i++ {
		num = num * i
	}
	fmt.Printf("The Factorial of %d is %d", f, num)
}
