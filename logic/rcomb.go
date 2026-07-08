package logic

import (
    "fmt"
    "math"
    "os"
    "sort"
    "strings"
)

type Topology int

const (
    TSingle Topology = iota
    T2Series
    T2Parallel
    T3Series
    T3Parallel
    TParallelSeries // (R1||R2) + R3
    TSeriesParallel // (R1+R2) || R3
)

func (t Topology) String() string {
    switch t {
    case TSingle:
        return "single resistor"
    case T2Series:
        return "2 resistors in series"
    case T2Parallel:
        return "2 resistors in parallel"
    case T3Series:
        return "3 resistors in series"
    case T3Parallel:
        return "3 resistors in parallel"
    case TParallelSeries:
        return "(R1||R2)+R3"
    case TSeriesParallel:
        return "(R1+R2)||R3"
    }
    return ""
}

type CombResult struct {
    topo          Topology
    r1, r2, r3   float64
    value        float64
    diff         float64
    errPct       float64
}

type rcombSeries struct {
    name   string
    series []float64
}

var rcombSeriesList = []rcombSeries{
    {"E3", e3},
    {"E6", e6},
    {"E12", e12},
    {"E24", e24},
    {"E48", e48},
    {"E96", e96},
    {"E128", e128},
}

func Rcomb_help() {
    fmt.Println("Find the best resistor combination to approximate a target resistance.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " rcomb VALUE [SERIES]")
    fmt.Println()
    fmt.Println("VALUE can be any resistance (e.g. 10k, 4k7, 100R, 12.3k)")
    fmt.Println("SERIES (optional): E3, E6, E12, E24 (default), E48, E96, E128")
    fmt.Println()
    fmt.Println("Evaluates up to 3 resistors in the following topologies:")
    fmt.Println("  single, series, parallel, 3-series, 3-parallel,")
    fmt.Println("  (R1||R2)+R3, (R1+R2)||R3")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " rcomb 10k")
    fmt.Println("Example: " + os.Args[0] + " rcomb 12.3k E96")
    fmt.Println("Example: " + os.Args[0] + " rcomb 4k7 E12")
}

func parseSeries(name string) ([]float64, bool) {
    upper := strings.ToUpper(name)
    for _, s := range rcombSeriesList {
        if s.name == upper {
            return s.series, true
        }
    }
    return nil, false
}

func generateValues(series []float64, target float64) []float64 {
    _, n := sci_norm(target)
    minExp := n - 3
    maxExp := n + 3
    if minExp < -1 {
        minExp = -1
    }
    if maxExp > 7 {
        maxExp = 7
    }
    var values []float64
    for exp := minExp; exp <= maxExp; exp++ {
        mul := math.Pow10(exp)
        for _, base := range series {
            values = append(values, base*mul)
        }
    }
    sort.Float64s(values)
    return values
}

func closestIndex(values []float64, target float64) int {
    i := sort.SearchFloat64s(values, target)
    if i == 0 {
        return 0
    }
    if i == len(values) {
        return len(values) - 1
    }
    if target-values[i-1] <= values[i]-target {
        return i - 1
    }
    return i
}

func SearchCombinations(target float64, series []float64) []CombResult {
    values := generateValues(series, target)

    var results []CombResult

    addResult := func(topo Topology, r1, r2, r3, combValue float64) {
        diff := math.Abs(combValue - target)
        errPct := 100.0 * diff / target
        results = append(results, CombResult{
            topo: topo, r1: r1, r2: r2, r3: r3,
            value: combValue, diff: diff, errPct: errPct,
        })
    }

    // ---- 1R ----
    for _, r := range values {
        addResult(TSingle, r, 0, 0, r)
    }

    // ---- 2S ----
    for _, r1 := range values {
        if r1 > target {
            break
        }
        r2need := target - r1
        idx := closestIndex(values, r2need)
        r2 := values[idx]
        if r2 < r1 {
            continue
        }
        addResult(T2Series, r1, r2, 0, r1+r2)
    }

    // ---- 2P ----
    for _, r1 := range values {
        if r1 <= target {
            continue
        }
        r2need := target * r1 / (r1 - target)
        idx := closestIndex(values, r2need)
        r2 := values[idx]
        if r2 < r1 {
            continue
        }
        par := r1 * r2 / (r1 + r2)
        addResult(T2Parallel, r1, r2, 0, par)
    }

    // ---- Precompute pairs ----
    type sumPair struct{ r1, r2, sum float64 }
    var sumPairs []sumPair
    for i := 0; i < len(values); i++ {
        for j := i; j < len(values); j++ {
            sumPairs = append(sumPairs, sumPair{values[i], values[j], values[i] + values[j]})
        }
    }

    type parPair struct{ r1, r2, par float64 }
    var parPairs []parPair
    for i := 0; i < len(values); i++ {
        for j := i; j < len(values); j++ {
            r1, r2 := values[i], values[j]
            parPairs = append(parPairs, parPair{r1, r2, r1 * r2 / (r1 + r2)})
        }
    }

    // ---- 3S ----
    for _, sp := range sumPairs {
        if sp.sum >= target {
            continue
        }
        r3need := target - sp.sum
        idx := closestIndex(values, r3need)
        r3 := values[idx]
        rs := []float64{sp.r1, sp.r2, r3}
        sort.Float64s(rs)
        addResult(T3Series, rs[0], rs[1], rs[2], sp.sum+r3)
    }

    // ---- 3P ----
    for _, pp := range parPairs {
        if pp.par <= target {
            continue
        }
        r3need := target * pp.par / (pp.par - target)
        idx := closestIndex(values, r3need)
        r3 := values[idx]
        par3 := 1.0 / (1.0/pp.r1 + 1.0/pp.r2 + 1.0/r3)
        rs := []float64{pp.r1, pp.r2, r3}
        sort.Float64s(rs)
        addResult(T3Parallel, rs[0], rs[1], rs[2], par3)
    }

    // ---- (R1||R2) + R3 ----
    for _, pp := range parPairs {
        if pp.par >= target {
            continue
        }
        r3need := target - pp.par
        idx := closestIndex(values, r3need)
        r3 := values[idx]
        addResult(TParallelSeries, pp.r1, pp.r2, r3, pp.par+r3)
    }

    // ---- (R1+R2) || R3 ----
    for _, sp := range sumPairs {
        if sp.sum <= target {
            continue
        }
        r3need := target * sp.sum / (sp.sum - target)
        idx := closestIndex(values, r3need)
        r3 := values[idx]
        comb := sp.sum * r3 / (sp.sum + r3)
        addResult(TSeriesParallel, sp.r1, sp.r2, r3, comb)
    }

    // Sort by error ascending
    sort.Slice(results, func(i, j int) bool {
        if results[i].errPct != results[j].errPct {
            return results[i].errPct < results[j].errPct
        }
        return results[i].topo < results[j].topo
    })

    // Deduplicate
    type comboKey struct {
        topo Topology
        vals string
    }
    seen := make(map[comboKey]bool)
    var deduped []CombResult
    for _, r := range results {
        key := comboKey{topo: r.topo}
        vals := []float64{r.r1, r.r2, r.r3}
        sort.Float64s(vals)
        key.vals = fmt.Sprintf("%.10f|%.10f|%.10f", vals[0], vals[1], vals[2])
        if seen[key] {
            continue
        }
        seen[key] = true
        deduped = append(deduped, r)
    }

    return deduped
}

func Rcomb_exec(args []string) {
    if len(args) < 1 {
        fmt.Println("Expected: resistance value [optional series]")
        return
    }

    val, err := ParseQuantity(args[0])
    if err != nil {
        fmt.Println("Error parsing resistance:", err)
        return
    }
    target := val.v

    if target <= 0 || Zero(target) {
        fmt.Println("Resistance must be positive")
        return
    }

    val.u = U_Ohm

    series := e24
    seriesName := "E24"
    if len(args) >= 2 {
        s, ok := parseSeries(args[1])
        if !ok {
            fmt.Printf("Unknown series: %s (valid: E3, E6, E12, E24, E48, E96, E128)\n", args[1])
            return
        }
        series = s
        seriesName = strings.ToUpper(args[1])
    }

    deduped := SearchCombinations(target, series)

    n := 3
    if len(deduped) < n {
        n = len(deduped)
    }
    if n == 0 {
        fmt.Println("No valid combinations found")
        return
    }

    fmt.Printf("Top %d combinations for %s (%s):\n", n, Val{v: target, u: val.u}.ToString(), seriesName)
    for i := 0; i < n; i++ {
        r := deduped[i]
        var line string
        switch r.topo {
        case TSingle:
            line = fmt.Sprintf("  %d. single: %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case T2Series:
            line = fmt.Sprintf("  %d. series: %s + %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case T2Parallel:
            line = fmt.Sprintf("  %d. parallel: %s ∥ %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case T3Series:
            line = fmt.Sprintf("  %d. series: %s + %s + %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.r3, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case T3Parallel:
            line = fmt.Sprintf("  %d. parallel: %s ∥ %s ∥ %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.r3, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case TParallelSeries:
            line = fmt.Sprintf("  %d. (R1∥R2)+R3: (%s ∥ %s) + %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.r3, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        case TSeriesParallel:
            line = fmt.Sprintf("  %d. (R1+R2)∥R3: (%s + %s) ∥ %s = %s", i+1,
                Val{v: r.r1, u: val.u}.ToString(),
                Val{v: r.r2, u: val.u}.ToString(),
                Val{v: r.r3, u: val.u}.ToString(),
                Val{v: r.value, u: val.u}.ToString())
        }
        line += fmt.Sprintf(", diff %s, error %.2f%%",
            Val{v: r.diff, u: val.u}.ToString(), r.errPct)
        fmt.Println(line)
    }
}
