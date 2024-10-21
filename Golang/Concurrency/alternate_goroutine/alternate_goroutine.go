package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	var firstLock, secondLock sync.Mutex
	secondLock.Lock()
	var num int
	wg.Add(2)
	go func(){
		defer wg.Done()
		for inx :=0; inx < 100; inx++{
			firstLock.Lock()
			fmt.Println("First Rountine : ", num);
			num++;
			secondLock.Unlock()
		}
	}()
	go func(){
		defer wg.Done()
		for inx :=0; inx < 100; inx++{
			secondLock.Lock()
			fmt.Println("Second Rountine : ", num);
			num++;
			firstLock.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println("Main Rountine exiting !!")

}