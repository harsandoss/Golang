//to find a Right angle Triangle
package main

import "fmt"

func main() {
	var i, j, r, c int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	for i = 1; i <= r; i++ {
		for j = 1; j <= i; j++ {
			if i%2 == 0 {
				fmt.Print("0 ")
			} else {
				fmt.Print("1 ")
			}
		}
		fmt.Println()
	}
}
