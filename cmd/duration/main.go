package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	duration := flag.Duration("sleep", time.Duration(50*time.Millisecond), "Time to sleep before exiting")

	flag.Parse()

	fmt.Printf("Sleeping for %v\n", *duration)

	time.Sleep(*duration)

	fmt.Printf("Exiting....\n")
}
