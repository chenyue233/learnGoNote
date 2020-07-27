package main

import (
	"os"
	"io"
	"fmt"
)

func Countdown(out io.Writer)  {
	for i := 3;i > 0;i--{
		fmt.Fprintln(out,i)
	}
	fmt.Fprint(out,"Go!")
}

func main()  {
	Countdown(os.Stdout)
}
