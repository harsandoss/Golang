// to find a square number
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number:")
	fmt.Scan(&n)
	square := n * n
	fmt.Printf("square of %d is %d", n, square)
}
