package main

import (
	"time"
	"fmt"
	"runtime"
)

func main()  {
	var a [10]int
	for i:=0; i<10;i++{
		go func(i int) {
			// fmt.Printf("Hello from" + "goroutine %d\n",i)
			for{
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}


