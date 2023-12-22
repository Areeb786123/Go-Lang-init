package main

import "fmt"

func main() {

	var list = []int{1, 2, 3, 4, 54, 5, 6}
	fmt.Println(list)
	list = append(list, 234)
	fmt.Println(list)

	list = append(list[1:])
	fmt.Println(list)

	// to remove the value

}
