package main

import (
	"fmt"
)

func main(){
	//var wg sync.WaitGroup
	for inx := 0; inx < 101; inx++ {
		if inx%3 == 0{
			//wg.Add(1)
			go func(){
				//defer wg.Done()
				fmt.Println("====> ", inx)
			}()
		} else {
			fmt.Println("=> ", inx);
		}
	}
	//wg.Wait()
	fmt.Println("Exiting Main Goroutine !!")
}