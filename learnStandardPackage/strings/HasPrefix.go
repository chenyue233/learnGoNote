package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "sadasds123"
	b := strings.HasPrefix(s1,"s")
	fmt.Printf("%v\n",b) // true
}