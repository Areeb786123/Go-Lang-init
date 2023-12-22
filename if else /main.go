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
	ans, _ := reader.ReadString('\n')
	fmt.Println("Tell us your age ")
	ansValue, _ := strconv.ParseFloat(strings.TrimSpace(ans), 64)

	if ansValue >= 18 {
		fmt.Println("get set go with traffic rules")
	} else if ansValue == 17 {
		fmt.Println("Ready to gear up")
	} else {
		fmt.Println("Under age for driving")
	}
}
