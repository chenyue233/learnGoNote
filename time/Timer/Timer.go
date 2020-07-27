package main

import (
	"fmt"
	"time"
)

func main()  {
	// 1.timer基本使用
	// timer1 := time.NewTimer(2 * time.Second)
	// t1 := time.Now()
	// fmt.Printf("t1:%v\n", t1)
	// t2 := <-timer1.C
	// fmt.Printf("t2:%v\n", t2)
	// 2.验证timer只能响应1次
	// timer2 := time.NewTimer(time.Second)
	// for {
	// <-timer2.C
	// fmt.Println("时间到")
	// }
	
	// 3.timer实现延时的功能
	// (1)
	time.Sleep(time.Second)
	// (2)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到a")
	// (3)
	<-time.After(2*time.Second)
	fmt.Println("4秒到")
}
