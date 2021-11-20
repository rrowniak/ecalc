package logic

import "testing"

func TestOhm(t *testing.T) {
	r, w, e := calc([]string{"1A", "1V"})

	if e != nil {
		t.Errorf("Error is not expected")
	}

	if r.ToString() != "1.00 Ω" {
		t.Errorf("Value is wrong")
	}

	if w.ToString() != "1.00 W" {
		t.Errorf("Value is wrong")
	}
}
