//To find the given number is Even or Odd
package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number:")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Printf("%d is a Even Number", a)
	} else {
		fmt.Printf("%d is a odd Number", a)
	}
}
