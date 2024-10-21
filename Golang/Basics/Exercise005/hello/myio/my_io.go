package myio1

import (
	"bufio"
	"fmt"
	"os"
)

var input string

func ReadString() {
	var scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		input = scanner.Text()
	}
} 

func PrintString(){
	fmt.Println("Last input string is : " + input)
}