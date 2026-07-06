package logic

import (
    "math"
    "testing"
)

func almostEqual(a, b, eps float64) bool {
    return math.Abs(a-b) <= eps
}

func TestWireParseVoltage(t *testing.T) {
    pv, err := parseWireArg("12V")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paVoltage || pv.value != 12.0 {
        t.Errorf("expected voltage 12V, got type=%d val=%f", pv.typ, pv.value)
    }
}

func TestWireParseVoltageDecimal(t *testing.T) {
    pv, err := parseWireArg("5.5V")
    if err != nil {
        t.Fatal("unexpected error:", err)
    }
    if pv.typ != paVoltage || pv.value != 5.5 {
        t.Errorf("expected voltage 5.5V, got type=%d val=%f", pv.typ, pv.value)
    }
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

func TestAWG14ToMm2(t *testing.T) {
    mm2 := awgToMm2(14)
    if !almostEqual(mm2, 2.08, 0.01) {
        t.Errorf("AWG 14 expected ~2.08 mm², got %f", mm2)
    }
}

func TestAWG20ToMm2(t *testing.T) {
    mm2 := awgToMm2(20)
    if !almostEqual(mm2, 0.518, 0.01) {
        t.Errorf("AWG 20 expected ~0.518 mm², got %f", mm2)
    }
}

func TestDiameterToMm2(t *testing.T) {
    mm2 := diameterToMm2(0.0015)
    if !almostEqual(mm2, 1.767, 0.001) {
        t.Errorf("1.5mm diameter expected ~1.767 mm², got %f", mm2)
    }
}

func TestResistanceOhm(t *testing.T) {
    r := resistanceOhm(2.5, 100)
    if !almostEqual(r, 0.6896, 0.001) {
        t.Errorf("2.5mm² 100m expected ~0.6896 Ω, got %f", r)
    }
}

func TestWireCalcCrossSection(t *testing.T) {
    result, err := WireCalc([]string{"12V", "2.5", "100m"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    expectedR := 0.6896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }

    expectedI5 := 12.0 * 0.05 / expectedR
    if !almostEqual(result.D5.I, expectedI5, 0.01) {
        t.Errorf("current@5%% expected %f, got %f", expectedI5, result.D5.I)
    }

    expectedI10 := 12.0 * 0.10 / expectedR
    if !almostEqual(result.D10.I, expectedI10, 0.01) {
        t.Errorf("current@10%% expected %f, got %f", expectedI10, result.D10.I)
    }

    expectedI20 := 12.0 * 0.20 / expectedR
    if !almostEqual(result.D20.I, expectedI20, 0.01) {
        t.Errorf("current@20%% expected %f, got %f", expectedI20, result.D20.I)
    }

    expectedV5 := 12.0 * 0.05
    expectedP5 := expectedV5 * expectedI5
    if !almostEqual(result.D5.P, expectedP5, 0.01) {
        t.Errorf("power@5%% expected %f, got %f", expectedP5, result.D5.P)
    }
}

func TestWireCalcMm2WithUnit(t *testing.T) {
    result, err := WireCalc([]string{"2.5mm2", "12V", "10m"})
    if err != nil {
        t.Fatal("unexpected error:", err)
    }

    expectedR := 0.06896
    if !almostEqual(result.ResistanceOhm, expectedR, 0.001) {
        t.Errorf("resistance expected %f, got %f", expectedR, result.ResistanceOhm)
    }
}

func TestWireCalcDiameter(t *testing.T) {
    result, err := WireCalc([]string{"12V", "1.5mm", "60cm"})
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
    result, err := WireCalc([]string{"60cm", "12V", "1.5mm"})
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
    result, err := WireCalc([]string{"12V", "14AWG", "100ft"})
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
    _, err := WireCalc([]string{"12V", "1.5mm"})
    if err == nil {
        t.Error("expected error for missing args")
    }
}

func TestWireCalcNoVoltage(t *testing.T) {
    _, err := WireCalc([]string{"1.5mm", "60cm"})
    if err == nil {
        t.Error("expected error for missing voltage")
    }
}

func TestWireCalcNoWireSpec(t *testing.T) {
    _, err := WireCalc([]string{"12V", "60cm"})
    if err == nil {
        t.Error("expected error for missing wire spec")
    }
}

func TestWireCalcNoLength(t *testing.T) {
    _, err := WireCalc([]string{"12V", "2.5mm2"})
    if err == nil {
        t.Error("expected error for missing length")
    }
}

func TestWireCalcMultipleVoltage(t *testing.T) {
    _, err := WireCalc([]string{"12V", "5V", "2.5mm2", "10m"})
    if err == nil {
        t.Error("expected error for multiple voltages")
    }
}

func TestWireCalcMultipleWireSpec(t *testing.T) {
    _, err := WireCalc([]string{"12V", "2.5mm2", "1.5mm", "10m"})
    if err == nil {
        t.Error("expected error for multiple wire specs")
    }
}

func TestWireCalcNegativeCrossSection(t *testing.T) {
    _, err := WireCalc([]string{"12V", "-2.5", "10m"})
    if err == nil {
        t.Error("expected error for negative cross-section")
    }
}
