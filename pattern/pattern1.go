// print box letter pattern
package main

import "fmt"

func main() {
	var r, c int
	var x, y string
	fmt.Print("Enter the Number of Rows:")
	fmt.Scanln(&r)
	fmt.Print("Enter the Number of columns:")
	fmt.Scanln(&c)
	fmt.Print("Enter the Outer layer Letter:")
	fmt.Scanln(&x)
	fmt.Print("Enter the inner layer Letter:")
	fmt.Scanln(&y)
	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			if i == 1 || i == r || j == 1 || j == c {
				fmt.Printf("%s ", x)
			} else {
				fmt.Printf("%s ", y)
			}
		}
		fmt.Println()
	}

}
