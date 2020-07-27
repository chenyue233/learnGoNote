package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "hello 世界!"
	s2 := "hello 世界!"
	b := strings.EqualFold(s1,s2)
	fmt.Printf("%v\n",b) // true
}
