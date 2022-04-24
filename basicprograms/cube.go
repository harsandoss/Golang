// to find a cube number
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number:")
	fmt.Scan(&n)
	cube := n * n * n
	fmt.Printf("cube of %d is %d", n, cube)
}
