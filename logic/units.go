package logic

import (
    "errors"
    "fmt"
    "math"
    "strconv"
)

type UnitSymbol int

const (
    U_undefined UnitSymbol = iota
    U_unknown
    U_Ohm
    U_V
    U_A
    U_W
    U_Db
    U_Hz
    U_H
    U_F
)

func (u UnitSymbol) ToString() string {
    switch u {
    case U_undefined:
        return ""
    case U_Ohm:
        return "Ω"
    case U_A:
        return "A"
    case U_V:
        return "V"
    case U_W:
        return "W"
    case U_Db:
        return "dB"
    case U_Hz:
        return "Hz"
    case U_H:
        return "H"
    case U_F:
        return "F"
    }
    return "unknown"
}

func (u UnitSymbol) ToStringLong() string {
    switch u {
    case U_undefined:
        return "undefined"
    case U_Ohm:
        return "Ohm"
    case U_A:
        return "Ampere"
    case U_V:
        return "Volt"
    case U_W:
        return "Watt"
    case U_Db:
        return "Decibel"
    case U_Hz:
        return "Hertz"
    case U_H:
        return "Henr"
    case U_F:
        return "Farad"
    }
    return "unknown"
}

func ParseUnitSymbol(s string) UnitSymbol {
    switch s {
    case "":
        return U_undefined
    case "R", "Ohm", "Ω":
        return U_Ohm
    case "A":
        return U_A
    case "V":
        return U_V
    case "W":
        return U_W
    case "dB":
        return U_Db
    case "Hz":
        return U_Hz
    case "H":
        return U_H
    case "F":
        return U_F
    }
    return U_unknown
}

var prefixes = map[string]float64{
    "T": 1e12,
    "G": 1e9,
    "M": 1e6,
    "k": 1e3,
    "m": 1e-3,
    "u": 1e-6,
    "n": 1e-9,
    "p": 1e-12,
    "f": 1e-15,
}

func isPrefix(s string) bool {
    _, e := prefixes[s]
    return e
}

type Val struct {
    v float64
    u UnitSymbol
}

func (quantity Val) ToString() string {
    // find proper prefix, a brute force method
    var pref_s string
    var pref_v float64
    for k, v := range prefixes {
        if quantity.v >= v && quantity.v < v*1000.0 {
            pref_s = k
            pref_v = v
            break
        }
    }

    if len(pref_s) > 0 {
        v := quantity.v / pref_v
        return fmt.Sprintf("%.2f %s%s", v, pref_s, quantity.u.ToString())
    }

    if quantity.u == U_undefined {
        return fmt.Sprintf("%.2f", quantity.v)
    }

    return fmt.Sprintf("%.2f %s", quantity.v, quantity.u.ToString())
}

const (
    parsing_sign = iota
    parsing_integer_part
    parsing_fractional_part
    parsing_literal
)

func ParseQuantity(s string) (Val, error) {
    var v Val

    if len(s) == 0 {
        return v, errors.New("empty string to be parsed")
    }

    state := parsing_sign
    sign := 1
    integer_part := ""
    fractional_part := ""
    literal := ""
    prefix := ""

    for _, c := range s {
        switch {
        case c == '-' || c == '+':
            if state != parsing_sign {
                return v, errors.New("unexpected sign symbol")
            }
            if c == '-' {
                sign = -1
            }
        case c >= '0' && c <= '9':
            if state == parsing_sign {
                state = parsing_integer_part
            }
            if state == parsing_literal {
                state = parsing_fractional_part
            }
            if state == parsing_integer_part {
                integer_part += string(c)
            } else if state == parsing_fractional_part {
                fractional_part += string(c)
            }
        case c == '.' || c == ',':
            if state == parsing_sign || state == parsing_integer_part {
                state = parsing_fractional_part
            }
        case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
            state = parsing_literal
            literal += string(c)
            if isPrefix(literal) {
                prefix = literal
                literal = ""
            }
        }
    }

    i := 0
    if len(integer_part) > 0 {
        ii, e := strconv.Atoi(integer_part)
        if e != nil {
            return v, e
        }
        i = ii
    }

    f := 0
    if len(fractional_part) > 0 {
        ff, e := strconv.Atoi(fractional_part)
        if e != nil {
            return v, e
        }
        f = ff
    }

    n := -len(fractional_part)
    dec := float64(i) + float64(f)*math.Pow10(n)
    dec *= float64(sign)

    if len(prefix) > 0 {
        dec *= prefixes[prefix]
    }

    us := U_undefined
    if len(literal) > 0 {
        us = ParseUnitSymbol(literal)
    }

    v.v = dec
    v.u = us

    return v, nil
}
