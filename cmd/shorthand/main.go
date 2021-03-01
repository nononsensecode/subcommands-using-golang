package main

import (
	"flag"
	"fmt"
)

var name string

const (
	defaultName = "unknown"
	description = "Name of the user"
)

func main() {
	flag.StringVar(&name, "name", defaultName, description)
	flag.StringVar(&name, "n", defaultName, description)

	flag.Parse()

	fmt.Printf("Name: %s\n", name)
}
