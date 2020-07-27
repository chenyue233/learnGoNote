package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "sadasds123"
	b := strings.HasSuffix(s1,"3")
	fmt.Printf("%v\n",b) // true
}
