package main

import "fmt"

func main() {
	var i, j, k, r, c int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	for i = 1; i <= r; i++ {
		for j = i; j <= r; j++ {
			fmt.Print(" ")
		}
		for k = 1; k < i; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
	for i = r; i > 0; i-- {
		for j = i; j < r; j++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
}
