package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "Hello world!"
	b := strings.Contains(s1,"wor")
	fmt.Printf("%v\n",b)  // true
}
