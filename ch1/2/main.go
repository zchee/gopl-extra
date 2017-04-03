package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("n: %d, arg: %s\n", i, arg)
	}
}
