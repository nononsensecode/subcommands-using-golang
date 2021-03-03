package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	person := flag.NewFlagSet("person", flag.ExitOnError)
	name := person.String("name", "unknown", "Name of the person")

	currentUsageFn := person.Usage
	person.Usage = func() {
		fmt.Printf("Before printing usage: \n")
		currentUsageFn()
		fmt.Printf("After the usage message :\n")
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case person.Name():
			person.Parse(os.Args[2:])
			fmt.Printf("Name: %s\n", *name)

		default:
			fmt.Printf("Invalid flag\n")
		}
	}
}
