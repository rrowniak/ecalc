package main

import (
	"ecalc/logic"
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("Calculator for electronics.")
	fmt.Println()
	fmt.Println("usage: " + os.Args[0] + " <command> [<args>]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("\thelp\t\tPrints this help")
	fmt.Println("\tohm\t\tOhm's law calculator")
	fmt.Println("\tvdiv\t\tVoltage divider calculator")
	fmt.Println("\teseries\t\tResistor E-series calculator")
	fmt.Println()
	fmt.Println("See '" + os.Args[0] + " <command> help' to read about a specific command")
}

func printErrExit() {
	fmt.Println("Expected 'help', 'ohm', 'vdiv' or 'eseries' subcommands")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		printErrExit()
	}
	switch os.Args[1] {
	case "help":
		printHelp()
	case "ohm":
		if len(os.Args) < 4 || os.Args[3] == "help" {
			logic.Ohm_help()
			os.Exit(1)
		}
		logic.Ohm_exec(os.Args[2], os.Args[3])
	case "vdiv":
		if len(os.Args) < 5 || os.Args[3] == "help" {
			logic.Vdiv_help()
			os.Exit(1)
		}
		logic.Vdiv_exec(os.Args[2:])
	case "teseries":
		fmt.Println("Not supported yet")
	default:
		printErrExit()
	}
}
