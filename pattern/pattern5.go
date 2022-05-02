package main

import "fmt"

func main() {
	var r, c int
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	for i := 1; i <= r; i++ {
		for j := 1; j <= r-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
	for i := r - 1; i > 0; i-- {
		for j := 1; j <= r-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Printf("%d", k)
		}
		fmt.Println()
	}
}
