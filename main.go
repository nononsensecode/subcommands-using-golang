package main

import (
	"flag"
	"fmt"
	"os"
)

type Command struct {
	flagSet *flag.FlagSet
}

func (c Command) Execute(f func()) {
	f()
}

func (c Command) Name() string {
	return c.flagSet.Name()
}


func main() {

	subCommands := []Command{
		Command{flag.NewFlagSet("name", flag.ContinueOnError)},
		Command{flag.NewFlagSet("contact", flag.ContinueOnError)},
	}

	
	firstname := subCommands[0].flagSet.String("firstname", "unknown", "first name of the user")
	lastname := subCommands[0].flagSet.String("lastname", "unknown", "last name of the user")

	mobile := subCommands[1].flagSet.String("mobile", "", "mobile number of the user")
	email := subCommands[1].flagSet.String("email", "", "email id of the user")

	// Get the subcommand which is the 1st argument
	subcommand := os.Args[1]

	for _, cmd := range subcommands {
		if cmd.Name() == subcommand {
			cmd.Parse(os.Args[2:])
		}
	}
	
	fmt.Printf("Name: %s %s\n", *firstname, *lastname)
	fmt.Printf("Mobile: %s, Email: %s\n", *mobile, *email)
}