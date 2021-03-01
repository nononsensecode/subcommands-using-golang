package main

import (
	"flag"
	"fmt"
)

var name string

type gender bool

func (g *gender) String() string {
	if *g {
		return "Mr."
	}

	return "Ms."
}

func (g *gender) Set(s string) error {
	switch s {
	case "m":
		*g = true
		return nil
	case "f":
		*g = false
		return nil
	default:
		panic("Invalid gender")
	}
}

func main() {
	name := flag.String("name", "unknown", "Name of the user")

	var isMale = gender(true)

	flag.Var(&isMale, "gender", "Gender of the user")

	flag.Parse()

	fmt.Printf("Name: %s %s\n", prefix, *name)
}
