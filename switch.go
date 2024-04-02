package main


import (
	"fmt"
	"time"
)


func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's Aefore Noon")
	default:
		fmt.Println("It's After Noon")
	}
	whatami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Don't know the type %T \n", t)
		}
	}
	whatami(true)
	whatami(1)
	whatami("hey")
}
