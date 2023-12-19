package main

import "fmt"

func main() {
	fmt.Println("hello")

	number := 23
	var ptr = &number
	fmt.Println(ptr)  // this will print  the pointer address where is it stored output 0x1400000e0b0
	fmt.Println(*ptr) // using * will tell the value of  the pointer output 23

}
