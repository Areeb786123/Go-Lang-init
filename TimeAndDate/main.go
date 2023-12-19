package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	//changing the formate of time
	/*
		to formate we have some specific reserved values
		1. For year 01-02-2006
		2. For time 15:04:05
		3.For Days  Moday
	*/

	//Here is the example

	//for Date
	fmt.Println(presentTime.Format("01 - 02 - 2006"))

	//for Time
	fmt.Println(presentTime.Format("15:04:05"))

	//for Days
	fmt.Println(presentTime.Format("Monday"))

	//You will get the output of todays date and time
}
