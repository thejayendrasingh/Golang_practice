// Write a program that accepts a comma separated sequence of words as input and prints the words in a comma-separated sequence
// after sorting them alphabetically.

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	var input string
	fmt.Println("Insert comma separated sequence of words to sort : ")
	fmt.Scan(&input)

	var words []string = strings.Split(input, ",")
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	fmt.Println(strings.Join(words, ","))
}