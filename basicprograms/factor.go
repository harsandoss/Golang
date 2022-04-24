package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the Number:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
