// Write a program which takes 2 digits, X,Y as input and generates a 2-dimensional array.
// The element value in the i-th row and j-th column of the array should be i*j
package main

import "fmt"

func main(){
	fmt.Println("Please enter 2 dims : ")
	var row,col int
	fmt.Scan(&row)
	fmt.Scan(&col)

	var mat = make([][]int, row)
	for inx := 0; inx < row; inx++ {
		mat[inx] = make([]int, col)
		for jnx := 0; jnx < col; jnx++{
			mat[inx][jnx] = inx*jnx;
		}
	}
	fmt.Println("Matrix : ", mat)
}