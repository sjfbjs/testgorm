package main

import "fmt"
import "time"

func main(){
	sum := 0
	var start = time.Now()
	for i := 1; i<= 10000000; i++ {
		sum += i
	}
	fmt.Println(sum)
    var cost = time.Since(start)
    fmt.Println(cost)
}