package logic

import (
    "fmt"
    "math"
    "os"
)

func Lc_help() {
    fmt.Println("Calculate LC resonant frequency, or capacitance, or inductance.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " lc VAL1 VAL2")
    fmt.Println()
    fmt.Println("VALs may represent resonant frequency, capacitance or inductance in any order.")
    fmt.Println("For two provided quantities the third one will be calculated.")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " lc 1MHz 1nF")
    fmt.Println("Example: " + os.Args[0] + " lc 5uH 225kHz")
    fmt.Println("Example: " + os.Args[0] + " lc 10uH 33nF")
}

type selector struct {
    a, b UnitSymbol
}

type calcFun func(Val, Val) (Val, error)

var selectors map[selector]calcFun

func init() {
    selectors = map[selector]calcFun{
        {U_Hz, U_H}: lc_calc_cap,
        {U_H, U_Hz}: lc_calc_cap,
        {U_F, U_H}:  lc_calc_freq,
        {U_H, U_F}:  lc_calc_freq,
        {U_Hz, U_F}: lc_calc_ind,
        {U_F, U_Hz}: lc_calc_ind,
    }
}

func lc_calc_freq(v1, v2 Val) (Val, error) {
    // TODO: add unit check
    var freq Val
    freq.u = U_Hz
    freq.v = 1.0 / (2 * math.Pi * math.Sqrt(v1.v*v2.v))
    return freq, nil
}

func lc_calc_cap(v1, v2 Val) (Val, error) {
    var cap Val
    var l, f float64

    if v1.u == U_Hz {
        f, l = v1.v, v2.v
    } else {
        f, l = v2.v, v1.v
    }

    cap.u = U_F
    cap.v = 1.0 / (math.Pow(2*math.Pi*f, 2) * l)
    return cap, nil
}

func lc_calc_ind(v1, v2 Val) (Val, error) {
    var ind Val
    var c, f float64

    if v1.u == U_Hz {
        f, c = v1.v, v2.v
    } else {
        f, c = v2.v, v1.v
    }

    ind.u = U_H
    ind.v = 1.0 / (math.Pow(2*math.Pi*f, 2) * c)
    return ind, nil
}

func Lc_exec(args []string) {
    if len(args) != 2 {
        fmt.Println("Two arguments expected")
        return
    }

    var v1, v2 Val
    var e error
    v1, e = ParseQuantity(args[0])
    if e != nil {
        fmt.Println(e)
        return
    }

    v2, e = ParseQuantity(args[1])
    if e != nil {
        fmt.Println(e)
        return
    }

    // select a calc fun
    fun, exists := selectors[selector{a: v1.u, b: v2.u}]
    if !exists {
        fmt.Println("Wrong units provided")
        return
    }

    ret, e := fun(v1, v2)
    if e != nil {
        fmt.Println(e)
        return
    }

    fmt.Printf("%s (%f %s)\n", ret.ToString(), ret.v, ret.u.ToString())
}
