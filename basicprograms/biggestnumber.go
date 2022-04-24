package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter the Numbers:")
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Printf("%d,a is the largest number ", a)
	} else if b > a && b > c {
		fmt.Printf("%d,b is the largest number", b)
	} else {
		fmt.Printf("%d,c is the largest number", c)
	}
}
