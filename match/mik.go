package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(test(n))

}
func test(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	} else if n == 2 {
		return 3
	} else if n == 3 {
		return 4
	} else if n > 3 {
		return test(n) + test(n-3)
	} else {

		return 0
	}

}
