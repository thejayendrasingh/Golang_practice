// Create a slice from comma-seperated input string
package main

import (
	"fmt"
	"strings"
)

func main() {
	var nums string
	fmt.Println("Please enter comma separated numbers : ")

	fmt.Scan(&nums)
	fmt.Println(strings.Split(nums, ","))
}
