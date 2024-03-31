package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("Seven is even ")
	} else {
		fmt.Println("Seven is odd")
	}
	if 8%4 == 0 {
		fmt.Println("eight is divisble by 4")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even")
	}
	if num := 9; num < 0 {
		fmt.Println(num, " is negative ")
	} else if num < 10 {
		fmt.Println(num, "num has 1 digit")
	} else {
		fmt.Println("it has multiple digits")
	}
}
