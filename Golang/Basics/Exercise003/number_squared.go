// Create a map with numbers squared
package main

import (
	"fmt"
)

func main(){
	var number int

	fmt.Print("Please input a number : ");
	fmt.Scan(&number);

	var m = make(map[int]int)
	for inx := 1; inx <= number; inx++{
		m[inx] = inx*inx
	}
	fmt.Println(m);
}