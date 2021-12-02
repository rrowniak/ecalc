package logic

import (
    "fmt"
    "os"
)

func React_help() {
    fmt.Println("Electrical reactance calculator.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " react VAL1 VAl2")
    fmt.Println()
    fmt.Println("VALs may represent the following combinations:")
    fmt.Printf("\t* frequency [Hz] AND capacitance [F], inductance [H],")
    fmt.Printf(" capacitive reactance [Ω], or inductive reactance [Ω]\n")
    fmt.Printf("\t* capacitance [F], inductance [H] AND ")
    fmt.Printf(" capacitive reactance [Ω], or inductive reactance [Ω]\n")
    fmt.Println("For two provided quantities the remaining will be calculated.")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " react 1MHz 1kR")
    fmt.Println("Example: " + os.Args[0] + " react 5uH 225kHz")
    fmt.Println("Example: " + os.Args[0] + " react 10uH 33nF")
}

func React_exec(args []string) {

}
