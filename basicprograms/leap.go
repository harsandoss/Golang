package main

import "fmt"

func main() {
	var y int
	fmt.Print("Enter the Year:")
	fmt.Scan(&y)
	if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
		fmt.Println(y, "is a leap year")
	} else {
		fmt.Println(y, "is not a leap year")
	}
}
