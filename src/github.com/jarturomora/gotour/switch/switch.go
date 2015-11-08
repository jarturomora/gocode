package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	t := time.Date(2015, time.November, 6, 10, 57, 0, 0, time.UTC)
	today := t.Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}