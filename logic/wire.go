package logic

import (
    "errors"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

const copperRho = 1.724e-8

type lengthUnit struct {
    suffix     string
    multiplier float64
}

var lengthUnits = []lengthUnit{
    {"mm", 1e-3},
    {"cm", 1e-2},
    {"km", 1e3},
    {"ft", 0.3048},
    {"in", 0.0254},
    {"yd", 0.9144},
    {"m", 1.0},
}

type parsedArg int

const (
    paVoltage parsedArg = iota
    paAWG
    paCrossSection
    paLength
)

type parsedValue struct {
    typ   parsedArg
    value float64
}

func Wire_help() {
    fmt.Println("Wire resistance and voltage drop calculator.")
    fmt.Println()
    fmt.Println("usage: " + os.Args[0] + " wire VOLTAGE WIRE_SPEC LENGTH")
    fmt.Println()
    fmt.Println("Arguments may be in any order. WIRE_SPEC can be:")
    fmt.Println("  diameter (e.g. 1.5mm, 0.8mm)")
    fmt.Println("  cross-section (e.g. 2.5mm2, 1.5)")
    fmt.Println("  AWG gauge (e.g. 14AWG, 20AWG)")
    fmt.Println("LENGTH examples: 60cm, 10m, 2km, 100ft")
    fmt.Println()
    fmt.Println("Voltage examples: 12V, 230V, 5V")
    fmt.Println()
    fmt.Println("Practical ampacity estimates:")
    fmt.Println("  power wiring   ~4 A/mm²  (bundled or in conduit, poor heat dissipation)")
    fmt.Println("  chassis wiring ~10 A/mm² (single wire in open air, good heat dissipation)")
    fmt.Println()
    fmt.Println("Example: " + os.Args[0] + " wire 12V 1.5mm 60cm")
    fmt.Println("Example: " + os.Args[0] + " wire 230V 2.5 100m")
    fmt.Println("Example: " + os.Args[0] + " wire 12V 14AWG 100ft")
}

func awgToMm2(awg float64) float64 {
    dMm := 0.127 * math.Pow(92, (36-awg)/39)
    return math.Pi * (dMm / 2) * (dMm / 2)
}

func diameterToMm2(diameterM float64) float64 {
    dMm := diameterM * 1000
    return math.Pi * (dMm / 2) * (dMm / 2)
}

func mm2ToDiameterMm(aMm2 float64) float64 {
    return 2 * math.Sqrt(aMm2/math.Pi)
}

func resistanceOhm(crossSectionMm2, lengthM float64) float64 {
    return copperRho * lengthM / (crossSectionMm2 * 1e-6)
}

func calcDrop(voltageV, resistanceOhm float64, dropPerc float64) DropInfo {
    v := voltageV * (dropPerc / 100.0)
    i := v / resistanceOhm
    p := v * i
    return DropInfo{VDropV: v, I: i, P: p}
}

func parseWireArg(s string) (parsedValue, error) {
    if len(s) == 0 {
        return parsedValue{}, errors.New("empty string")
    }

    v, err := ParseQuantity(s)
    if err == nil && v.u == U_V {
        return parsedValue{typ: paVoltage, value: v.v}, nil
    }

    upper := strings.ToUpper(s)
    lower := strings.ToLower(s)

    if strings.HasSuffix(upper, "AWG") {
        gaugeStr := strings.TrimSpace(s[:len(s)-3])
        gauge, err := strconv.Atoi(gaugeStr)
        if err != nil {
            return parsedValue{}, errors.New("invalid AWG value: " + gaugeStr)
        }
        if gauge < 0 || gauge > 56 {
            return parsedValue{}, errors.New("AWG value out of range (0-56)")
        }
        return parsedValue{typ: paAWG, value: float64(gauge)}, nil
    }

    if strings.HasSuffix(lower, "mm2") {
        numStr := strings.TrimSpace(s[:len(s)-3])
        val, err := strconv.ParseFloat(numStr, 64)
        if err != nil {
            return parsedValue{}, errors.New("invalid cross-section: " + numStr)
        }
        if val <= 0 {
            return parsedValue{}, errors.New("cross-section must be positive")
        }
        return parsedValue{typ: paCrossSection, value: val}, nil
    }

    for _, lu := range lengthUnits {
        if strings.HasSuffix(lower, lu.suffix) {
            numStr := strings.TrimSpace(s[:len(s)-len(lu.suffix)])
            val, err := strconv.ParseFloat(numStr, 64)
            if err != nil {
                return parsedValue{}, errors.New("invalid " + lu.suffix + " value: " + numStr)
            }
            if val <= 0 {
                return parsedValue{}, errors.New(lu.suffix + " value must be positive")
            }
            return parsedValue{typ: paLength, value: val * lu.multiplier}, nil
        }
    }

    val, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return parsedValue{}, errors.New("unable to parse: " + s)
    }
    if val <= 0 {
        return parsedValue{}, errors.New("value must be positive")
    }
    return parsedValue{typ: paCrossSection, value: val}, nil
}

type DropInfo struct {
    VDropV float64
    I      float64
    P      float64
}

type WireCalcResult struct {
    ResistanceOhm   float64
    CrossSectionMm2 float64
    DiameterMm      float64
    D1              DropInfo
    D5              DropInfo
    D10             DropInfo
    D20             DropInfo
}

func WireCalc(args []string) (WireCalcResult, error) {
    var result WireCalcResult

    if len(args) < 3 {
        return result, errors.New("three arguments expected: voltage, wire spec, and length")
    }

    var voltage, crossSectionMm2, lengthM float64
    var hasVoltage, hasWireSpec, hasLength bool
    var lengthValues []float64

    for _, arg := range args {
        pv, err := parseWireArg(arg)
        if err != nil {
            return result, err
        }

        switch pv.typ {
        case paVoltage:
            if hasVoltage {
                return result, errors.New("multiple voltage values provided")
            }
            voltage = pv.value
            hasVoltage = true
        case paAWG:
            if hasWireSpec {
                return result, errors.New("multiple wire specifications provided")
            }
            crossSectionMm2 = awgToMm2(pv.value)
            hasWireSpec = true
        case paCrossSection:
            if hasWireSpec {
                return result, errors.New("multiple wire specifications provided")
            }
            crossSectionMm2 = pv.value
            hasWireSpec = true
        case paLength:
            lengthValues = append(lengthValues, pv.value)
        }
    }

    switch len(lengthValues) {
    case 0:
        return result, errors.New("no length provided")
    case 1:
        lengthM = lengthValues[0]
        hasLength = true
    case 2:
        if hasWireSpec {
            return result, errors.New("ambiguous: wire spec combined with two length-like values")
        }
        if lengthValues[0] < lengthValues[1] {
            crossSectionMm2 = diameterToMm2(lengthValues[0])
            lengthM = lengthValues[1]
        } else {
            crossSectionMm2 = diameterToMm2(lengthValues[1])
            lengthM = lengthValues[0]
        }
        hasWireSpec = true
        hasLength = true
    default:
        return result, errors.New("too many length-like values")
    }

    if !hasVoltage {
        return result, errors.New("no voltage provided")
    }
    if !hasWireSpec {
        return result, errors.New("no wire specification provided (diameter, cross-section, or AWG)")
    }
    if !hasLength {
        return result, errors.New("no length provided")
    }

    r := resistanceOhm(crossSectionMm2, lengthM)

    result.ResistanceOhm = r
    result.CrossSectionMm2 = crossSectionMm2
    result.DiameterMm = mm2ToDiameterMm(crossSectionMm2)
    result.D1 = calcDrop(voltage, r, 1)
    result.D5 = calcDrop(voltage, r, 5)
    result.D10 = calcDrop(voltage, r, 10)
    result.D20 = calcDrop(voltage, r, 20)

    return result, nil
}

func ampacityPower(aMm2 float64) float64   { return 4 * aMm2 }
func ampacityChassis(aMm2 float64) float64 { return 10 * aMm2 }

func Wire_exec(args []string) {
    result, err := WireCalc(args)
    if err != nil {
        fmt.Println(err)
        return
    }

    rVal := Val{v: result.ResistanceOhm, u: U_Ohm}
    fmt.Printf("Resistance: %s (%f %s)\n", rVal.ToString(), result.ResistanceOhm, rVal.u.ToString())
    fmt.Printf("Cross-section: %.2f mm², diameter: %.2f mm\n",
        result.CrossSectionMm2, result.DiameterMm)

    fmt.Printf("Practical ampacity (power wiring): ~%s\n",
        Val{v: ampacityPower(result.CrossSectionMm2), u: U_A}.ToString())
    fmt.Printf("Practical ampacity (chassis wiring): ~%s\n",
        Val{v: ampacityChassis(result.CrossSectionMm2), u: U_A}.ToString())
    fmt.Println()

    printDropLine(1, result.D1)
    printDropLine(5, result.D5)
    printDropLine(10, result.D10)
    printDropLine(20, result.D20)
}

func printDropLine(dropPerc int, d DropInfo) {
    cVal := Val{v: d.I, u: U_A}
    vVal := Val{v: d.VDropV, u: U_V}
    pVal := Val{v: d.P, u: U_W}
    fmt.Printf("  @ %d%% drop: %s (V_drop %s, %s)\n", dropPerc,
        cVal.ToString(), vVal.ToString(), pVal.ToString())
}
