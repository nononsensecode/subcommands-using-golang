package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Func("usage", "Default message", func(s string) error {
		fmt.Printf("Testing usage function: %s\n", s)
		return nil
	})
}
