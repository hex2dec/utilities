package main

import (
	"flag"
	"fmt"
)

var (
	hFlag  bool
	hValue = false
	hUsage = "print usage message"
)

func init() {
	flag.BoolVar(&hFlag, "help", hValue, hUsage)
	flag.BoolVar(&hFlag, "h", hValue, hUsage+"(shorthand)")
}

func main() {
	flag.Parse()
	command := flag.Arg(0)

	if hFlag {
		usage()
		return
	}

	handleCommand(command)
}

func handleCommand(cmd string) {
	switch cmd {
	case "help":
		usage()
	default:
		usage()
	}
}

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), "Usage of code-manager")
	flag.PrintDefaults()
}
