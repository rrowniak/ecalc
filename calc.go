package main

import (
    "ecalc/logic"
    "fmt"
    "os"
)

type Command struct {
    name          string
    description   string
    help_fun      func()
    min_args_exec int
    exec_fun      func([]string)
}

var commands []Command

func init() {
    commands = append(commands, Command{name: "help",
        description:   "Prints this help",
        min_args_exec: 0,
        help_fun:      nil, // to avoid infinite recursion
        exec_fun:      nil})

    commands = append(commands, Command{name: "ohm",
        description:   "Ohm's law calculator",
        min_args_exec: 2,
        help_fun:      logic.Ohm_help,
        exec_fun:      logic.Ohm_exec})

    commands = append(commands, Command{name: "vdiv",
        description:   "Voltage divider calculator",
        min_args_exec: 3,
        help_fun:      logic.Vdiv_help,
        exec_fun:      logic.Vdiv_exec})

    commands = append(commands, Command{name: "lc",
        description:   "LC resonant calculator",
        min_args_exec: 2,
        help_fun:      logic.Lc_help,
        exec_fun:      logic.Lc_exec})

    commands = append(commands, Command{name: "react",
        description:   "Reactance calculator",
        min_args_exec: 2,
        help_fun:      logic.React_help,
        exec_fun:      logic.React_exec})

    commands = append(commands, Command{name: "db",
        description:   "Decibel (dB) calculator",
        min_args_exec: 2,
        help_fun:      logic.Db_help,
        exec_fun:      logic.Db_exec})

    commands = append(commands, Command{name: "eseries",
        description:   "Resistor E-series calculator",
        min_args_exec: 1,
        help_fun:      logic.Eseries_help,
        exec_fun:      logic.Eseries_exec})
}

func printHelp() {
    fmt.Println("Calculator for electronics.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " <command> [<args>]")
    fmt.Println()
    fmt.Println("Commands:")
    for _, cmd := range commands {
        fmt.Printf("\t%s\t\t%s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    fmt.Println("See '" + os.Args[0] + " <command> help' to read about a specific command")
}

func printErrExit() {
    // fmt.Println("Expected 'help', 'ohm', 'vdiv' or 'eseries' subcommands")
    fmt.Printf("Expected ")
    for i, cmd := range commands {
        if i != len(commands)-1 {
            fmt.Printf("%s, ", cmd.name)
        } else {
            fmt.Printf("%s subcommands\n", cmd.name)
        }
    }
    os.Exit(1)
}

func main() {
    if len(os.Args) < 2 {
        printHelp()
        printErrExit()
    }

    handled := false
    for _, cmd := range commands {
        if os.Args[1] == cmd.name {

            if cmd.name == "help" {
                // a special case
                printHelp()
                os.Exit(1)
            }

            if len(os.Args) < 3 || os.Args[2] == "help" {
                if cmd.help_fun != nil {
                    cmd.help_fun()
                } else {
                    fmt.Println("No help available for this command")
                }
                os.Exit(1)
            }

            if len(os.Args) < 2+cmd.min_args_exec {
                fmt.Printf("Expected at least %d arguments\n", cmd.min_args_exec)
                os.Exit(1)
            }

            if cmd.exec_fun == nil {
                fmt.Println("This command is not implemented")
                os.Exit(1)
            }

            cmd.exec_fun(os.Args[2:])

            handled = true
            break
        }
    }

    if !handled {
        println("Unknown command: " + os.Args[1])
        printErrExit()
    }
}
