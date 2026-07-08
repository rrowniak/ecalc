package logic

import "testing"

func TestParseSeries(t *testing.T) {
    s, ok := parseSeries("E24")
    if !ok {
        t.Fatal("expected E24 to be found")
    }
    if len(s) != 24 {
        t.Errorf("expected 24 values, got %d", len(s))
    }
}

func TestParseSeriesLowercase(t *testing.T) {
    s, ok := parseSeries("e96")
    if !ok {
        t.Fatal("expected e96 to be found")
    }
    if len(s) != 96 {
        t.Errorf("expected 96 values, got %d", len(s))
    }
}

func TestParseSeriesUnknown(t *testing.T) {
    _, ok := parseSeries("E99")
    if ok {
        t.Error("expected E99 to not be found")
    }
}

func TestGenerateValuesRange(t *testing.T) {
    vals := generateValues(e24, 10000)
    if len(vals) == 0 {
        t.Fatal("expected non-empty values")
    }
    if vals[0] < 0.1 {
        t.Errorf("expected min >= 0.1, got %f", vals[0])
    }
}

func TestCloseIndex(t *testing.T) {
    vals := []float64{1, 2, 3, 4, 5}
    if idx := closestIndex(vals, 2.4); idx != 1 {
        t.Errorf("expected index 1 for 2.4, got %d", idx)
    }
    if idx := closestIndex(vals, 2.6); idx != 2 {
        t.Errorf("expected index 2 for 2.6, got %d", idx)
    }
    if idx := closestIndex(vals, 0.5); idx != 0 {
        t.Errorf("expected index 0 for 0.5, got %d", idx)
    }
    if idx := closestIndex(vals, 10); idx != 4 {
        t.Errorf("expected index 4 for 10, got %d", idx)
    }
}

func TestSearchCombinationsSingle(t *testing.T) {
    result := SearchCombinations(10000, e24)
    if len(result) == 0 {
        t.Fatal("expected at least one result")
    }
    if result[0].topo != TSingle || !almostEqual(result[0].r1, 10000, 1) {
        t.Errorf("expected single 10k, got topo=%d r1=%f", result[0].topo, result[0].r1)
    }
}

func TestSearchCombinationsSorted(t *testing.T) {
    result := SearchCombinations(12300, e24)
    if len(result) < 2 {
        t.Fatal("expected at least 2 results")
    }
    sorted := true
    for i := 1; i < len(result); i++ {
        if result[i].errPct < result[i-1].errPct-1e-9 {
            sorted = false
            break
        }
    }
    if !sorted {
        t.Error("results not sorted by error ascending")
    }
}
