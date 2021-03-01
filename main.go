package main

import (
	"flag"
	"fmt"
	"os"
)

// SubCommand describes name of the subcommand as well as the function to execute it
type SubCommand struct {
	FlagSet *flag.FlagSet
	Execute func(f *flag.FlagSet, args []string, vars ...interface{})
}

//Name returns the name of the subcommand
func (c SubCommand) Name() string {
	return c.FlagSet.Name()
}

const cFirstName = "firstname"
const cLastName = "lastname"
const cMobile = "mobile"
const cEmail = "email"

func main() {

	name := SubCommand{
		FlagSet: flag.NewFlagSet("name", flag.ContinueOnError),
		Execute: func(f *flag.FlagSet, args []string, vars ...interface{}) {
			f.Parse(args)

			if cap(vars) < 2 {
				panic("Invalid number of arguments")
			}

			fmt.Printf("Name: %s %s\n", *vars[0].(*string), *vars[1].(*string))
		},
	}

	contact := SubCommand{
		FlagSet: flag.NewFlagSet("contact", flag.ContinueOnError),
		Execute: func(f *flag.FlagSet, args []string, vars ...interface{}) {
			f.Parse(args)

			if cap(vars) < 2 {
				panic("Invalid number of arguments")
			}

			fmt.Printf("Mobile: %s, Email: %s\n", *vars[0].(*string), *vars[1].(*string))
		},
	}

	firstname := name.FlagSet.String("firstname", "unknown", "first name of the user")
	lastname := name.FlagSet.String("lastname", "unknown", "last name of the user")

	mobile := contact.FlagSet.String("mobile", "", "mobile number of the user")
	email := contact.FlagSet.String("email", "", "email id of the user")

	// Get the subcommand which is the 1st argument
	if cap(os.Args) < 2 {
		os.Exit(0)
	}

	subCommand := os.Args[1]

	switch subCommand {
	case name.Name():
		name.Execute(name.FlagSet, os.Args[2:], firstname, lastname)
	case contact.Name():
		contact.Execute(contact.FlagSet, os.Args[2:], mobile, email)
	}
}
