package main

import (
	"flag"
	"fmt"
)

type gender string

func (g *gender) String() string {
	return string(*g)
}

func (g *gender) Set(s string) error {
	switch s {
	case "m":
		*g = "Mr."
		return nil
	case "f":
		*g = "Ms."
		return nil
	default:
		panic("Invalid gender")
	}
}

func main() {
	name := flag.String("name", "unknown", "Name of the user")

	var isMale = gender("Mr")


	flag.Var(&isMale, "gender", "Gender of the user")
	flag.Var(&isMale, "g", "Gender of the")

	flag.Parse()

	fmt.Printf("Name: %s %s\n", isMale.String(), *name)
}
