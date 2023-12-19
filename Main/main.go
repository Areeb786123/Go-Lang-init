package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter something")

	input, _ := reader.ReadString('\n')
	fmt.Println("thankyou for the rating", input)

	//converting string to int
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // here we add 64 means the  the int type we want to add
	addNum := numRating + 1
	//here  we have to use if check to print because we are getting input and error both simultaniusly
	if err != nil {
		fmt.Println("Hey you go error --> ", err)
	} else {
		fmt.Println(addNum)
	}

}
