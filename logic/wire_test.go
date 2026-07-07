package logic

import (
    "math"
    "testing"
)

func almostEqual(a, b, eps float64) bool {
    return math.Abs(a-b) <= eps
}

func TestWireParseAWG(t *testing.T) {
    pv, err := parseWireArg("14AWG")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paAWG || pv.value != 14.0 {
        t.Errorf("expected AWG 14, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseAWGLowercase(t *testing.T) {
    pv, err := parseWireArg("20awg")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paAWG || pv.value != 20.0 {
        t.Errorf("expected AWG 20, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseCrossSectionMm2(t *testing.T) {
    pv, err := parseWireArg("2.5mm2")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paCrossSection || pv.value != 2.5 {
        t.Errorf("expected cross-section 2.5mm2, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseCrossSectionBare(t *testing.T) {
    pv, err := parseWireArg("2.5")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paCrossSection || pv.value != 2.5 {
        t.Errorf("expected cross-section 2.5, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseDiameter(t *testing.T) {
    pv, err := parseWireArg("1.5mm")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || !almostEqual(pv.value, 0.0015, 1e-10) {
        t.Errorf("expected diameter 0.0015m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthCm(t *testing.T) {
    pv, err := parseWireArg("60cm")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || pv.value != 0.6 {
        t.Errorf("expected 0.6m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthM(t *testing.T) {
    pv, err := parseWireArg("100m")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || pv.value != 100.0 {
        t.Errorf("expected 100m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthFt(t *testing.T) {
    pv, err := parseWireArg("100ft")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || !almostEqual(pv.value, 30.48, 1e-10) {
        t.Errorf("expected 30.48m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthKm(t *testing.T) {
    pv, err := parseWireArg("2km")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || pv.value != 2000.0 {
        t.Errorf("expected 2000m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthIn(t *testing.T) {
    pv, err := parseWireArg("12in")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || !almostEqual(pv.value, 0.3048, 1e-10) {
        t.Errorf("expected 0.3048m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseLengthYd(t *testing.T) {
    pv, err := parseWireArg("3yd")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paLength || !almostEqual(pv.value, 2.7432, 1e-10) {
        t.Errorf("expected 2.7432m, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseEmpty(t *testing.T) {
    _, err := parseWireArg("")
    if err == nil {
        t.Error("expected error for empty string")
    }
}

func TestWireParseInvalidAWG(t *testing.T) {
    _, err := parseWireArg("100AWG")
    if err == nil {
        t.Error("expected error for out-of-range AWG")
    }
}

func TestWireParseNonNumericAWG(t *testing.T) {
    _, err := parseWireArg("abcAWG")
    if err == nil {
        t.Error("expected error for non-numeric AWG")
    }
}

func TestWireParseCurrent(t *testing.T) {
    pv, err := parseWireArg("5A")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paCurrent || pv.value != 5.0 {
        t.Errorf("expected current 5A, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseCurrentDecimal(t *testing.T) {
    pv, err := parseWireArg("2.5A")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paCurrent || pv.value != 2.5 {
        t.Errorf("expected current 2.5A, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseCurrentNegative(t *testing.T) {
    _, err := parseWireArg("-5A")
    if err == nil {
        t.Error("expected error for negative current")
    }
}

func TestWireParseCurrentNonNumeric(t *testing.T) {
    _, err := parseWireArg("abcA")
    if err == nil {
        t.Error("expected error for non-numeric current")
    }
}

func TestAWG14ToMm2(t *testing.T) {
    mm2 := awgToMm2(14)
    if !almostEqual(mm2, 2.08, 0.01) {
        t.Errorf("AWG 14 expected ~2.08 mm\u00b2, got %f", mm2)
    }
}

func TestAWG20ToMm2(t *testing.T) {
    mm2 := awgToMm2(20)
    if !almostEqual(mm2, 0.518, 0.01) {
        t.Errorf("AWG 20 expected ~0.518 mm\u00b2, got %f", mm2)
    }
}

func TestDiameterToMm2(t *testing.T) {
    mm2 := diameterToMm2(0.0015)
    if !almostEqual(mm2, 1.767, 0.001) {
        t.Errorf("1.5mm diameter expected ~1.767 mm\u00b2, got %f", mm2)
    }
}

func TestResistanceOhm(t *testing.T) {
    r := resistanceOhm(2.5, 100)
    if !almostEqual(r, 0.6896, 0.001) {
        t.Errorf("2.5mm\u00b2 100m expected ~0.6896 \u03a9, got %f", r)
    }
}

func TestWireCalcCrossSection(t *testing.T) {
    result, err := WireCalc([]string{"2.5", "100m"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    expectedR := 0.6896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }

    if !almostEqual(result.CrossSectionMm2, 2.5, 0.001) {
        t.Errorf("cross-section expected 2.5, got %f", result.CrossSectionMm2)
    }
}

func TestWireCalcMm2WithUnit(t *testing.T) {
    result, err := WireCalc([]string{"2.5mm2", "10m"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    expectedR := 0.06896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcDiameter(t *testing.T) {
    result, err := WireCalc([]string{"1.5mm", "60cm"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    cs := diameterToMm2(0.0015)
    expectedR := resistanceOhm(cs, 0.6)
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcDiameterReversedOrder(t *testing.T) {
    result, err := WireCalc([]string{"60cm", "1.5mm"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    cs := diameterToMm2(0.0015)
    expectedR := resistanceOhm(cs, 0.6)
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcAWG(t *testing.T) {
    result, err := WireCalc([]string{"14AWG", "100ft"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    cs := awgToMm2(14)
    expectedR := resistanceOhm(cs, 30.48)
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcMissingArgs(t *testing.T) {
    _, err := WireCalc([]string{"1.5mm"})
    if err == nil {
        t.Error("expected error for missing args")
    }
}

func TestWireCalcNoWireSpec(t *testing.T) {
    _, err := WireCalc([]string{"60cm"})
    if err == nil {
        t.Error("expected error for missing wire spec")
    }
}

func TestWireCalcNoLength(t *testing.T) {
    _, err := WireCalc([]string{"2.5mm2"})
    if err == nil {
        t.Error("expected error for missing length")
    }
}

func TestWireCalcMultipleWireSpec(t *testing.T) {
    _, err := WireCalc([]string{"2.5mm2", "1.5mm", "10m"})
    if err == nil {
        t.Error("expected error for multiple wire specs")
    }
}

func TestWireCalcWithCurrent(t *testing.T) {
    result, err := WireCalc([]string{"2.5mm2", "10m", "5A"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if !almostEqual(result.CurrentA, 5.0, 0.001) {
        t.Errorf("current expected 5A, got %f", result.CurrentA)
    }
    expectedR := 0.06896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcWithCurrentReversed(t *testing.T) {
    result, err := WireCalc([]string{"5A", "10m", "2.5mm2"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if !almostEqual(result.CurrentA, 5.0, 0.001) {
        t.Errorf("current expected 5A, got %f", result.CurrentA)
    }
    expectedR := 0.06896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcMultipleCurrent(t *testing.T) {
    _, err := WireCalc([]string{"2.5mm2", "10m", "5A", "2A"})
    if err == nil {
        t.Error("expected error for multiple current values")
    }
}

func TestWireCalcNegativeCrossSection(t *testing.T) {
    _, err := WireCalc([]string{"-2.5", "10m"})
    if err == nil {
        t.Error("expected error for negative cross-section")
    }
}

func TestMm2ToDiameterMm(t *testing.T) {
    d := mm2ToDiameterMm(1.767)
    if !almostEqual(d, 1.5, 0.01) {
        t.Errorf("1.767 mm\u00b2 expected ~1.5 mm diameter, got %f", d)
    }
}
