package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH") + "\n")
	
	t := 10;
	fmt.Printf("The value of the T is %d\n", t)
	
	switch {
	case t == 10:
		fmt.Printf("T is %d\n", t)
	}
}
