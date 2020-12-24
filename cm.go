package main

import (
	"flag"
	"fmt"
)

var helpFlag bool

func init() {
	flag.BoolVar(&helpFlag, "help", false, "print help message")
	flag.Parse()
}

func main() {
	command := flag.Arg(0)

	if helpFlag {
		help()
		return
	}

	handleCommand(command)
}

func handleCommand(cmd string) {
	switch cmd {
	case "help":
		help()
	default:
		help()
	}
}

func help() {
	fmt.Fprintln(flag.CommandLine.Output(), "Usage of cm")
	flag.PrintDefaults()
}
