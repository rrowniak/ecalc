package logic

import (
    "math"
    "testing"
)

func comp_f(v1, v2, perc_err float64) bool {
    diff := math.Abs(v1 - v2)

    return (100 * diff / v1) <= perc_err
}

func check_expected(t *testing.T, expected, got Val, err_exp bool, e error,
    perc_error int) {
    if !comp_f(expected.v, got.v, float64(perc_error)) {
        t.Errorf("Expected != got (%f != %f)", expected.v, got.v)
    }

    if e != nil && !err_exp {
        t.Error("Got unexpected error")
    } else if e == nil && err_exp {
        t.Error("Expected error")
    }
}

func p(s string) Val {
    v, _ := ParseQuantity(s)
    return v
}

func TestCap(t *testing.T) {
    f := p("1MHz")
    l := p("33uH")
    c, e := lc_calc_cap(l, f)
    check_expected(t, p("768pF"), c, false, e, 1)

    f = p("125kHz")
    l = p("2uH")
    c, e = lc_calc_cap(l, f)
    check_expected(t, p("810.6nF"), c, false, e, 1)
}

func TestInd(t *testing.T) {
    f := p("1MHz")
    c := p("768pF")
    l, e := lc_calc_ind(c, f)
    check_expected(t, p("33uH"), l, false, e, 1)

    f = p("125kHz")
    c = p("810.6nF")
    l, e = lc_calc_ind(c, f)
    check_expected(t, p("2uH"), l, false, e, 1)
}

func TestFreq(t *testing.T) {

    c := p("100pF")
    l := p("10uH")
    f, e := lc_calc_freq(l, c)
    check_expected(t, p("5MHz"), f, false, e, 1)

    c = p("33nF")
    l = p("10mH")
    f, e = lc_calc_freq(l, c)
    check_expected(t, p("8.8kHz"), f, false, e, 1)
}
