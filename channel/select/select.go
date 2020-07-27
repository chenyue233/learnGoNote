package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int,c chan int)  {
	for n := range c{
		fmt.Printf("Worker %d recevied %d\n",id,n)
	}
}

func createWorker(id int,) chan<- int {
	c := make(chan int)
	go worker(id,c)
	return c
}

func main() {
	var c1, c2= generator(), generator()
	var worker = createWorker(0)
	n := 0
	hasValue := false
	for {
		var activeworker chan <- int
		if hasValue{
			activeworker = worker
		}
		select {
			case n = <-c1:
				hasValue = true
			case n = <-c2:
				hasValue = true
			case activeworker <- n:
				hasValue = false

			// default:
			// 	fmt.Println("No value received")
		}
}
}
