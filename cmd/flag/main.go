package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "unknown", "Name of the person")

	flag.Parse()

	fmt.Printf("Name: %s\n", *name)

	fmt.Printf("1st argument: %s\n", flag.Arg(0))
	fmt.Printf("All arguments: %v\n", flag.Args())
}
