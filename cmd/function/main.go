package main

import (
	"flag"
	"fmt"
)

func main() {
	name := "unknown"
	flag.Func("name", "`name` of the user", func(s string) error {
		if len(s) < 5 {
			return fmt.Errorf("Length of name %s is less than 5", s)
		}

		name = s
		return nil
	})

	flag.Parse()

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("N args: %d\n", flag.NArg())
	fmt.Printf("N Flags: %d\n", flag.NFlag())
	fmt.Printf("Args: %s\n", flag.Args())
}
