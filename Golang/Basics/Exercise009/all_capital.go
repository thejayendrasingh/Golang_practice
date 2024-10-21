// Write a program that accepts sequence of lines as input and prints the lines after making all characters in the
// sentence capitalized.

package main

import "fmt"

/*
[Info  - 16:15:55] 2024/10/17 16:15:55 background imports cache refresh starting
[Info  - 16:15:55] 2024/10/17 16:15:55 background refresh finished after 4.997583ms
[Info  - 16:16:30] 2024/10/17 16:16:30 background imports cache refresh starting
[Info  - 16:16:30] 2024/10/17 16:16:30 background refresh finished after 3.960292ms
*/

func main(){
	var strArr = make([]string,4)
	var str string
	for inx := 0; inx < 4; inx++ {
		fmt.Scan(&str)
		strArr[inx] = str
	}

}

func toCapital1(str string) string{
	const diff int = 'a'-'A'
	var arr = []rune(str)
	for inx, ch := range arr {
		if ch >= 'a' && ch <= 'z'{
			arr[inx] = rune(int(ch)-diff)
		}
	}
	return string(arr)
}