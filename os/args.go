package main

import (
	"fmt"
	"os"
)

func main() {
	b := os.Args[1]
	e := os.Args[2]
	fmt.Println(b, e)
}
