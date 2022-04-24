package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter the Number:")
	fmt.Scan(&a)
	if a >= 0 {
		fmt.Printf("%d is a Positive Number", a)
	} else {
		fmt.Printf("%d is a Negative Number", a)
	}
}
