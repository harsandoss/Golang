//To find Rectangle Number Pattern
package main

import "fmt"

func main() {
	var r, c int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
