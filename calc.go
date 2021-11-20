package main

import (
	"ecalc/logic"
	"fmt"
	"os"
)

func printErrExit() {
	fmt.Println("Expected 'ohm' or 'vdiv' subcommands")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printErrExit()
	}
	switch os.Args[1] {
	case "ohm":
		if len(os.Args) < 4 || os.Args[3] == "help" {
			logic.Ohm_help()
			os.Exit(1)
		}
		logic.Ohm_exec(os.Args[2], os.Args[3])
	case "vdiv":
		printErrExit()
	default:
		printErrExit()
	}

}
