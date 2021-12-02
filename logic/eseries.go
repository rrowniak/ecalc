package logic

import (
    "errors"
    "fmt"
    "math"
    "os"
)

var e3 = []float64{1.0, 2.2, 4.7}

var e6 = []float64{1.0, 1.5, 2.2, 3.3, 4.7, 6.8}

var e12 = []float64{1.0, 1.2, 1.5, 1.8, 2.2, 2.7, 3.3, 3.9, 4.7, 5.6, 6.8, 8.2}

var e24 = []float64{1.0, 1.1, 1.2, 1.3, 1.5, 1.6, 1.8, 2.0, 2.2, 2.4, 2.7, 3.0,
    3.3, 3.6, 3.9, 4.3, 4.7, 5.1, 5.6, 6.2, 6.8, 7.5, 8.2, 9.1}

var e48 = []float64{1.00, 1.05, 1.10, 1.15, 1.21, 1.27, 1.33, 1.40, 1.47, 1.54,
    1.62, 1.69, 1.78, 1.87, 1.96, 2.05, 2.15, 2.26, 2.37, 2.49, 2.61, 2.74, 2.87,
    3.01, 3.16, 3.32, 3.48, 3.65, 3.83, 4.02, 4.22, 4.42, 4.64, 4.87, 5.11, 5.36,
    5.62, 5.90, 6.19, 6.49, 6.81, 7.15, 7.50, 7.87, 8.25, 8.66, 9.09, 9.53}

var e96 = []float64{1.00, 1.02, 1.05, 1.07, 1.10, 1.13, 1.15, 1.18, 1.21, 1.24,
    1.27, 1.30, 1.33, 1.37, 1.40, 1.43, 1.47, 1.50, 1.54, 1.58, 1.62, 1.65, 1.69,
    1.74, 1.78, 1.82, 1.87, 1.91, 1.96, 2.00, 2.05, 2.10, 2.15, 2.21, 2.26, 2.32,
    2.37, 2.43, 2.49, 2.55, 2.61, 2.67, 2.74, 2.80, 2.87, 2.94, 3.01, 3.09, 3.16,
    3.24, 3.32, 3.40, 3.48, 3.57, 3.65, 3.74, 3.83, 3.92, 4.02, 4.12, 4.22, 4.32,
    4.42, 4.53, 4.64, 4.75, 4.87, 4.99, 5.11, 5.23, 5.36, 5.49, 5.62, 5.76, 5.90,
    6.04, 6.19, 6.34, 6.49, 6.65, 6.81, 6.98, 7.15, 7.32, 7.50, 7.68, 7.87, 8.06,
    8.25, 8.45, 8.66, 8.87, 9.09, 9.31, 9.53, 9.76}

var e128 = []float64{1.00, 1.01, 1.02, 1.04, 1.05, 1.06, 1.07, 1.09, 1.10, 1.11,
    1.13, 1.14, 1.15, 1.17, 1.18, 1.20, 1.21, 1.23, 1.24, 1.26, 1.27, 1.29, 1.30,
    1.32, 1.33, 1.35, 1.37, 1.38, 1.40, 1.42, 1.43, 1.45, 1.47, 1.49, 1.50, 1.52,
    1.54, 1.56, 1.58, 1.60, 1.62, 1.64, 1.65, 1.67, 1.69, 1.72, 1.74, 1.76, 1.78,
    1.80, 1.82, 1.84, 1.87, 1.89, 1.91, 1.93, 1.96, 1.98, 2.00, 2.03, 2.05, 2.08,
    2.10, 2.13, 2.15, 2.18, 2.21, 2.23, 2.26, 2.29, 2.32, 2.34, 2.37, 2.40, 2.43,
    2.46, 2.49, 2.52, 2.55, 2.58, 2.61, 2.64, 2.67, 2.71, 2.74, 2.77, 2.80, 2.84,
    2.87, 2.91, 2.94, 2.98, 3.01, 3.05, 3.09, 3.12, 3.16, 3.20, 3.24, 3.28, 3.32,
    3.36, 3.40, 3.44, 3.48, 3.52, 3.57, 3.61, 3.65, 3.70, 3.74, 3.79, 3.83, 3.88,
    3.92, 3.97, 4.02, 4.07, 4.12, 4.17, 4.22, 4.27, 4.32, 4.37, 4.42, 4.48, 4.53,
    4.59, 4.64, 4.70, 4.75, 4.81, 4.87, 4.93, 4.99, 5.05, 5.11, 5.17, 5.23, 5.30,
    5.36, 5.42, 5.49, 5.56, 5.62, 5.69, 5.76, 5.83, 5.90, 5.97, 6.04, 6.12, 6.19,
    6.26, 6.34, 6.42, 6.49, 6.57, 6.65, 6.73, 6.81, 6.90, 6.98, 7.06, 7.15, 7.23,
    7.32, 7.41, 7.50, 7.59, 7.68, 7.77, 7.87, 7.96, 8.06, 8.16, 8.25, 8.35, 8.45,
    8.56, 8.66, 8.76, 8.87, 8.98, 9.09, 9.20, 9.31, 9.42, 9.53, 9.65, 9.76, 9.88}

type ESeries struct {
    name           string
    series         []float64
    tolerance_perc float64
}

func (s ESeries) Calc(v float64) (float64, int, float64, int, error) {

    if v < 0 || Zero(v) {
        return 0, 0, 0, 0, errors.New("value can't be negative nor zero")
    }

    m, n := sci_norm(v)

    for i, sval := range s.series {
        if Zero(sval - m) {
            return sval, n, sval, n, nil
        }

        if i == len(s.series)-1 {
            // last element
            return sval, n, s.series[0], n + 1, nil
        }

        if m >= sval && m < s.series[i+1] {
            return sval, n, s.series[i+1], n, nil
        }
    }

    return 0, 0, 0, 0, errors.New("eseries.calc: we shouldn't be there")
}

func sci_norm(v float64) (float64, int) {
    const (
        upper_limit float64 = 10.0
        lower_limit float64 = 1.0
    )
    n := 0

    for v >= upper_limit {
        v /= 10
        n++
    }

    for v < lower_limit {
        v *= 10
        n--
    }

    return v, n
}

var series = [...]ESeries{
    {name: "E3", series: e3, tolerance_perc: 40.0},
    {name: "E6", series: e6, tolerance_perc: 20.0},
    {name: "E12", series: e12, tolerance_perc: 10.0},
    {name: "E24", series: e24, tolerance_perc: 5.0},
    {name: "E48", series: e48, tolerance_perc: 2.0},
    {name: "E96", series: e96, tolerance_perc: 1.0},
    {name: "E128", series: e128, tolerance_perc: 0.5},
}

func Eseries_help() {
    fmt.Printf("Calculate closest match in E-series: ")
    first := true
    for _, s := range series {
        if first {
            fmt.Printf("%s", s.name)
            first = false
        } else {
            fmt.Printf(", %s", s.name)
        }
    }
    fmt.Println()
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " eseries VAL")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " eseries 10.07k")
}

func eser_print_out(prefix string, v_orig Val, m_calc float64, n_calc int) {
    l := Val{v: m_calc * math.Pow10(n_calc), u: v_orig.u}
    diff := math.Abs(v_orig.v - l.v)
    diff_v := Val{v: diff, u: v_orig.u}
    err := 100 * diff / v_orig.v
    fmt.Printf("%s: %.2f (%s), error: %.1f%%, diff: %s\n", prefix, m_calc,
        l.ToString(), err, diff_v.ToString())
}

func Eseries_exec(args []string) {
    for _, arg := range args {
        v, e := ParseQuantity(arg)

        if e != nil {
            fmt.Printf("Error while parsing %s: %s\n", arg, e)
            continue
        }
        for _, es := range series {
            l_m, l_n, u_m, u_n, e := es.Calc(v.v)

            if e != nil {
                fmt.Printf("Problem with calculating %s series: %s\n", es.name, e)
                continue
            }

            // print result
            fmt.Printf("Closest match to %f in series %s (tolerance %.1f%%):\n", v.v,
                es.name, es.tolerance_perc)

            eser_print_out("\tlower boundary", v, l_m, l_n)
            eser_print_out("\tupper boundary", v, u_m, u_n)
        }
    }
}
