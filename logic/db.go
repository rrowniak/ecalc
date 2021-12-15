package logic

import (
    "errors"
    "fmt"
    "math"
    "os"
)

func Db_help() {
    fmt.Println("Ddecibel calculator that calculates the ratio of a value to a fixed reference value")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " db VAL REF_VAL")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " db 1 10")
}

func db_calc(args []string) (pow, ampl Val, e error) {
    var val, ref Val

    if len(args) < 2 {
        return pow, ampl, errors.New("two values expected")
    }

    val, e = ParseQuantity(args[0])

    if e != nil {
        return pow, ampl, e
    }

    ref, e = ParseQuantity(args[1])

    if e != nil {
        return pow, ampl, e
    }

    if Zero(ref.v) {
        return pow, ampl, errors.New("reference value cannot be zero")
    }

    log10 := math.Log10(val.v / ref.v)

    pow.u = U_Db
    pow.v = 10 * log10

    ampl.u = U_Db
    ampl.v = 20 * log10

    return pow, ampl, nil
}

func Db_exec(args []string) {
    pow, ampl, e := db_calc(args)

    if e != nil {
        fmt.Println(e)
        return
    }

    fmt.Printf("Power ratio:     %s (%f %s)\n", pow.ToString(), pow.v, pow.u.ToString())
    fmt.Printf("Amplitude ratio: %s (%f %s)\n", ampl.ToString(), ampl.v, ampl.u.ToString())
}
