package main

import (
	"fmt"
)

func main() {

	numbers := []int{10, 20, 30}
	// _は使用しない変数に使う
	for _, n := range numbers {
		fmt.Println(n)
	}
	for n := range numbers {
		fmt.Println(n)
	}
}
