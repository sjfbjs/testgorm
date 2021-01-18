package main

import (
	"fmt"
	"sync"
)
import "time"




func main(){
	wg := sync.WaitGroup{}
	wg.Add(10000000)
	var start = time.Now()
	sum := 0
	var i int
	for i = 0; i< 10000000; i++ {
		sum += i
		wg.Done()
	}
	fmt.Print(sum)
	var cost = time.Since(start)
	fmt.Println(cost)
    wg.Wait()

}