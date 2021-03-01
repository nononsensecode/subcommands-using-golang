package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	name := flag.NewFlagSet("name", flag.ContinueOnError)
	contact := flag.NewFlagSet("contact", flag.ContinueOnError)

	firstname := name.String("firstname", "unknown", "first name of the user")
	lastname := name.String("lastname", "unknown", "last name of the user")

	mobile := contact.String("mobile", "", "mobile number of the user")
	email := contact.String("email", "", "email id of the user")

	// Get the subcommand which is the 1st argument
	if cap(os.Args) < 2 {
		os.Exit(0)
	}

	subCommand := os.Args[1]

	switch subCommand {
	case name.Name():
		name.Parse(os.Args[2:])
		fmt.Printf("Name: %s %s\n", *firstname, *lastname)
	case contact.Name():
		contact.Parse(os.Args[2:])
		fmt.Printf("Mobile: %s, Email: %s\n", *mobile, *email)
	}
}
