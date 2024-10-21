package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
func main(){
	fmt.Println("Enter comma separated numbers : ")
	var scanner = bufio.NewScanner(os.Stdin)
	var input string = ""
	if scanner.Scan() {
		input = scanner.Text()
	}

	var numStr []string = strings.Split(input, ",")
	var numbers []string
	for _,num := range numStr {
		sqrt, _ := strconv.Atoi(num)
		numbers = append(numbers, fmt.Sprint(math.Sqrt(float64((2*50*sqrt)/30))))
	}
	fmt.Println("Output : ", strings.Join(numbers, ","))
}