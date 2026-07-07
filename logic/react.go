package logic

import (
    "errors"
    "fmt"
    "math"
    "os"
)

func React_help() {
    fmt.Println("Electrical reactance calculator.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " react VAL1 VAL2")
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

func xc(f, c float64) float64 {
    return 1.0 / (2 * math.Pi * f * c)
}

func xl(f, l float64) float64 {
    return 2 * math.Pi * f * l
}

func React_exec(args []string) {
    if len(args) != 2 {
        fmt.Println("Two arguments expected")
        return
    }

    v1, e := ParseQuantity(args[0])
    if e != nil {
        fmt.Println(e)
        return
    }

    v2, e := ParseQuantity(args[1])
    if e != nil {
        fmt.Println(e)
        return
    }

    a, b := v1, v2
    if a.u == b.u {
        fmt.Println(errors.New("two different quantities expected"))
        return
    }

    pick := func(u UnitSymbol) (Val, bool) {
        if a.u == u {
            return a, true
        }
        if b.u == u {
            return b, true
        }
        return Val{}, false
    }

    hasHz := a.u == U_Hz || b.u == U_Hz
    hasF := a.u == U_F || b.u == U_F
    hasH := a.u == U_H || b.u == U_H
    hasR := a.u == U_Ohm || b.u == U_Ohm

    switch {
    case hasHz && hasH:
        f, _ := pick(U_Hz)
        l, _ := pick(U_H)
        x := Val{v: xl(f.v, l.v), u: U_Ohm}
        fmt.Printf("Inductive reactance: %s (%f %s)\n", x.ToString(), x.v, x.u.ToString())

    case hasHz && hasF:
        f, _ := pick(U_Hz)
        c, _ := pick(U_F)
        x := Val{v: xc(f.v, c.v), u: U_Ohm}
        fmt.Printf("Capacitive reactance: %s (%f %s)\n", x.ToString(), x.v, x.u.ToString())

    case hasHz && hasR:
        f, _ := pick(U_Hz)
        x, _ := pick(U_Ohm)
        l := Val{v: x.v / (2 * math.Pi * f.v), u: U_H}
        c := Val{v: 1.0 / (2 * math.Pi * f.v * x.v), u: U_F}
        fmt.Printf("Inductance: %s (%f %s)\n", l.ToString(), l.v, l.u.ToString())
        fmt.Printf("Capacitance: %s (%f %s)\n", c.ToString(), c.v, c.u.ToString())

    case hasH && hasF:
        l, _ := pick(U_H)
        c, _ := pick(U_F)
        f := Val{v: 1.0 / (2 * math.Pi * math.Sqrt(l.v*c.v)), u: U_Hz}
        fmt.Printf("Resonant frequency: %s (%f %s)\n", f.ToString(), f.v, f.u.ToString())

    case hasH && hasR:
        l, _ := pick(U_H)
        x, _ := pick(U_Ohm)
        f := Val{v: x.v / (2 * math.Pi * l.v), u: U_Hz}
        fmt.Printf("Frequency: %s (%f %s)\n", f.ToString(), f.v, f.u.ToString())

    case hasF && hasR:
        c, _ := pick(U_F)
        x, _ := pick(U_Ohm)
        f := Val{v: 1.0 / (2 * math.Pi * c.v * x.v), u: U_Hz}
        fmt.Printf("Frequency: %s (%f %s)\n", f.ToString(), f.v, f.u.ToString())

    default:
        fmt.Println("Unsupported unit combination")
    }
}
