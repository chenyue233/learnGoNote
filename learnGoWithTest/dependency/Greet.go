package dependency

import (
	"fmt"
	"bytes"
)

func Greet(writer *bytes.Buffer,name string)  {
	fmt.Fprintf(writer,"Hello,%s",name)
}
