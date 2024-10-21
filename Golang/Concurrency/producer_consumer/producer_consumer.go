package main

import (
	"fmt"
	"sync"
)

func main(){
	var limit int = 100
	var inx int
	var wg sync.WaitGroup
	const size int = 1
	var queue = make([]int, 0, size+1)
	var cond = sync.NewCond(&sync.Mutex{})
	var mu sync.Mutex
	
	var producerFunc = func(index int){
		defer wg.Done()
		for ;inx < limit;{
			cond.L.Lock()
			for;len(queue) >= size; {
				cond.Wait();
			}
			queue = append(queue, inx)
			fmt.Println(fmt.Sprintf("Produer %d : %d", index, inx));
			cond.L.Unlock()
			cond.Signal()
			inx++
		}
	}
	
	var consumerFunc = func(index int){
		defer wg.Done()
		for ;;{
			mu.Lock()
			cond.L.Lock()
			for;len(queue) == 0; {
				cond.Wait();
			}
			var inx = queue[0]
			fmt.Println(fmt.Sprintf("Consumer %d : %d", index, inx));
			queue = queue[1:]
			cond.L.Unlock()
			cond.Signal()
			mu.Unlock()
			if(inx == limit){
				break;
			}
		}
	}

	go producerFunc(1)
	go consumerFunc(1)

	wg.Wait()
	fmt.Println("Main thread exiting !")
}