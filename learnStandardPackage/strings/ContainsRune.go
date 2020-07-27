package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "hello 世界!"
	b := strings.ContainsRune(s1,'世')
	fmt.Printf("%v\n",b) // true
}

