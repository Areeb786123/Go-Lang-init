package main

import "fmt"

func main() {
	//creating new user
	areeb := User{"areeb", 122, true}
	fmt.Println(areeb)

	// output
	// {areeb 122 true}

}

type User struct {
	Name   string
	Age    int
	Status bool
}
