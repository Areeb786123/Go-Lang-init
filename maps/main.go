package main

import "fmt"

func main() {

	ma := make(map[int]string)
	ma[1] = "areeb"
	ma[2] = "wish"
	ma[3] = "dua"
	fmt.Println(ma)
	delete(ma, 3)
	fmt.Println(ma)

	//output
	// 	map[1:areeb 2:wish 3:dua]
	// map[1:areeb 2:wish]
}
