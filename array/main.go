package main

import "fmt"

//Note array  is not used much in  go lang instead slices is used here

func main() {
	var arr [2]string
	arr[0] = "areeb"
	arr[1] = "areeb"
	fmt.Println(arr)
	fmt.Println(len(arr)) // print length

	var arr2 = [6]int{1, 2, 3, 4, 4, 5}
	fmt.Println(arr2)
}
