package main

import (
	"errors"
	"fmt"
	"os"
)

type ErrorWithExitCode struct {
	error
	ExitCode int
}

var (
	ErrNoError = errors.New("")
)

func main() {
	command, err := ParseArgs(os.Args[1:])
	if err == ErrHelp {
		subcommand := ""
		if len(os.Args) > 1 {
			subcommand = os.Args[1]
		}
		err = displayHelp(HelpAliases[subcommand])
		if err != nil {
			println(err.Error())
		}
		return
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(255)
	}

	steward := &TTYSteward{}
	err = command.Run(steward)
	if err != nil {
		exiterr, ok := err.(ErrorWithExitCode)
		if !ok || exiterr.error != ErrNoError {
			fmt.Fprintln(os.Stderr, err)
		}
		if ok {
			os.Exit(exiterr.ExitCode)
		} else {
			os.Exit(1)
		}
	}
}
