package logic

import "testing"

func TestUnitToString(t *testing.T) {
    u := U_V

    s := u.ToString()

    if s != "V" {
        t.Errorf("Can't convert U_V to \"V\"")
    }
}

func TestParseToUnitSymbol(t *testing.T) {
    u := ParseUnitSymbol("R")
    if u != U_Ohm {
        t.Errorf("Can't parse R to U_Ohm")
    }
}

func checkUnit(t *testing.T, msg string, v_recv Val, e_recv error,
    v float64, u UnitSymbol, e error) {
    if e_recv != e {
        t.Errorf("%s: errors doesn't match", msg)
    }

    if v_recv.v != v {
        t.Errorf("%s: values (received) %f != %f", msg, v_recv.v, v)
    }

    if v_recv.u != u {
        t.Errorf("%s: units (received) %s != %s", msg, v_recv.u.ToString(),
            u.ToString())
    }
}

func TestParseQuantity(t *testing.T) {
    s := "10V"
    u, e := ParseQuantity(s)
    checkUnit(t, "10V", u, e, 10.0, U_V, nil)

    s = "10.01A"
    u, e = ParseQuantity(s)
    checkUnit(t, "10.01A", u, e, 10.01, U_A, nil)

    s = "0.99A"
    u, e = ParseQuantity(s)
    checkUnit(t, "0.99A", u, e, 0.99, U_A, nil)

    s = "10kR"
    u, e = ParseQuantity(s)
    checkUnit(t, "10kR", u, e, 10000.0, U_Ohm, nil)

    s = "4k7"
    u, e = ParseQuantity(s)
    checkUnit(t, "4k7", u, e, 4700.0, U_undefined, nil)

    s = "13.7mV"
    u, e = ParseQuantity(s)
    checkUnit(t, "13.7mV", u, e, 0.0137, U_V, nil)

    s = "1300.7Hz"
    u, e = ParseQuantity(s)
    checkUnit(t, "1.3kHz", u, e, 1300.7, U_Hz, nil)

    s = "13.7H"
    u, e = ParseQuantity(s)
    checkUnit(t, "13.7H", u, e, 13.7, U_H, nil)

    s = "13.7dB"
    u, e = ParseQuantity(s)
    checkUnit(t, "13.7dB", u, e, 13.7, U_Db, nil)
}
