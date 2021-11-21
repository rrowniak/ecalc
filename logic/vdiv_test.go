package logic

import "testing"

func TestVdivCalcVout(t *testing.T) {
	vin := Val{v: 10, u: U_V}
	r1 := Val{v: 1.0, u: U_Ohm}
	r2 := Val{v: 3.0, u: U_Ohm}

	vout, e := calc_vout(vin, r1, r2)

	if !Zero(vout.v - 7.50) {
		t.Errorf("Expected 7.5V, got: %s", vout.ToString())
	}

	if e != nil {
		t.Errorf("Expected no errors, got: %s", e)
	}
}

func TestVdivCalcVin(t *testing.T) {
	vout := Val{v: 3, u: U_V}
	r1 := Val{v: 3.0, u: U_Ohm}
	r2 := Val{v: 1.0, u: U_Ohm}

	vin, e := calc_vin(vout, r1, r2)

	if !Zero(vin.v - 12.0) {
		t.Errorf("Expected 12V, got: %s", vout.ToString())
	}

	if e != nil {
		t.Errorf("Expected no errors, got: %s", e)
	}
}

func TestVdivCalcR1(t *testing.T) {
	vout := Val{v: 1, u: U_V}
	vin := Val{v: 10, u: U_V}
	r2 := Val{v: 1.0, u: U_Ohm}

	r1, e := calc_r1(vin, vout, r2)

	if !Zero(r1.v - 9.0) {
		t.Errorf("Expected 9 Ohms, got: %s", vout.ToString())
	}

	if e != nil {
		t.Errorf("Expected no errors, got: %s", e)
	}
}

func TestVdivCalcR2(t *testing.T) {
	vout := Val{v: 1, u: U_V}
	vin := Val{v: 10, u: U_V}
	r1 := Val{v: 9.0, u: U_Ohm}

	r2, e := calc_r2(vin, vout, r1)

	if !Zero(r2.v - 1.0) {
		t.Errorf("Expected 1 Ohm, got: %s", vout.ToString())
	}

	if e != nil {
		t.Errorf("Expected no errors, got: %s", e)
	}
}
