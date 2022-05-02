// To find the K Shape
package main

import "fmt"

func main() {
	var i, j, r, c int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	for i = r; i > 0; i-- {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
	for i = 2; i <= r; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
