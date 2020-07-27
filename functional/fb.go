package main

import (
	"fmt"
	"strings"
)

func fibonacci() intGen {
	a,b :=0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(
	p []byte)(n int,err error){
		next := g()
		s := fmt.Sprintf("%d\n",next)
		return strings.NewReader(s).Read(p)
}
func main()  {
	f := fibonacci()

	for i:=0;i<=10;i++{
		fmt.Println(f())
	}

}