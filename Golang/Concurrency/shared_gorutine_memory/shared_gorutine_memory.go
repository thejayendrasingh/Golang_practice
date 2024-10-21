package main

import (
	"fmt"
	"sync"
)

func main(){
	var str string = "Main Goroutine"
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("Inside child goroutine");
		str = "Child Goroutine"
	}()
	wg.Wait()
	fmt.Println("After Wait group wait : ", str)
}