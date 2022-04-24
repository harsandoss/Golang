//To find even number 1 to n numbers
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Emter the Number:")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf(" %d", i)
		}
	}
}
