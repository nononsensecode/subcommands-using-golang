package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	c := flag.CommandLine

	name := c.String("name", "unknown", "Name of the person")

	c.Parse(os.Args[1:])

	fmt.Printf("Name: %s\n", *name)
}
