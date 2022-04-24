package main

import "fmt"

func main() {
	var a, reverse int
	fmt.Print("Enter the number:")
	fmt.Scan(&a)
	reverse = 0
	temp := a
	for {
		reminder := a % 10
		reverse = reverse*10 + reminder
		a = a / 10
		if a == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d is a palindrome", temp)
	} else {
		fmt.Printf("%d is not a palindrome", temp)
	}
}
