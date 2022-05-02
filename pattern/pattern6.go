//To print Floyd Triangle
package main

import "fmt"

func main() {
	var r, c, n int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	fmt.Print("Enter the Start Number:")
	fmt.Scan(&n)
	for i := 1; i <= r; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", n)
			n++
		}
		fmt.Println()
	}
}
