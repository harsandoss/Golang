// to concat two string
package main

import "fmt"

func main() {
	var v string
	v = addstr()
	fmt.Println(v)
}
func addstr() string {
	var x, y string
	fmt.Print("Enter the 1st String:")
	fmt.Scan(&x)
	fmt.Print("Enter the 2nd String:")
	fmt.Scan(&y)
	z := x + " " + y

	return z
}
