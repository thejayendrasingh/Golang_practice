// Compute factorial

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var count int
	var numbers []string

	fmt.Print("How many number to find factorial for : ")
	fmt.Scan(&count)
	fmt.Println("Insert all numbers below")
	for inx := 1; inx <= count; inx++{
		var input int
		fmt.Scan(&input);
		numbers = append(numbers, strconv.Itoa(fact(input)))
	}

	fmt.Println("Factorial of all numbers : ", strings.Join(numbers, ","))
}

func fact(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	if num < 0 {
		return 0
	}
	return num * fact(num-1)
}
