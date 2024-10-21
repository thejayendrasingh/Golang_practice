// Find numbers devisable by 7 but not by 5
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []string
	for inx := 2000; inx <= 3200; inx++ {
		if inx%7 == 0 && inx%5 != 0 {
			numbers = append(numbers, strconv.Itoa(inx))
		}
	}
	fmt.Println(strings.Join(numbers, ","))
}
