package main

import (
    "fmt"
    "strconv"
    "time"
)

func Sample1( ch chan int)  {
    for i:= 0;i< 10;i++{
        ch <- i
        time.Sleep(1*time.Second)
    }
}

func Sample2(ch chan string)  {
    for i:= 0;i< 11;i++{
        ch <- "i an go" + strconv.Itoa(i)
        time.Sleep(1*time.Second)
    }
}
func main()  {
    ch1 := make(chan string)
    ch2 := make(chan int)
    go Sample1(ch2)
    go Sample2(ch1)
    
   for {
       select {
       case str, ok := <-ch1:
           if !ok {
               fmt.Println("ch1 fail")
           }
           fmt.Println(str)
       case int1, ok := <-ch2:
           if !ok {
               fmt.Println("ch2 fail")
           }
           fmt.Println(int1)
       }
   }
    time.Sleep(50*time.Second)
}