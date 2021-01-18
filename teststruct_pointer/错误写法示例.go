package main

import (
	"fmt"
	"sync"
)
import "time"




func main(){
	wg := sync.WaitGroup{}

	var start = time.Now()
	sum := 0

	for i := 0; i<= 10000000; i++ {
		wg.Add(1)
		go func(i int){
			sum += i
			wg.Done()
		}(i)

	}
	fmt.Println(sum)
	var cost = time.Since(start)
	fmt.Println(cost)
    wg.Wait()

}