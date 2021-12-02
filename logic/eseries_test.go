package logic

import "testing"

func expect(t *testing.T, m float64, n int,
	m_expected float64, n_expected int) {
	if !Zero(m - m_expected) {
		t.Errorf("m != m_expected (%f != %f)\n", m, m_expected)
	}

	if n != n_expected {
		t.Errorf("n != n_expected (%d != %d)\n", n, n_expected)
	}
}

func TestSciNorm(t *testing.T) {
	m, n := sci_norm(1.0)
	expect(t, m, n, 1.0, 0)

	// > 1

	m, n = sci_norm(10.0)
	expect(t, m, n, 1.0, 1)

	m, n = sci_norm(100.0)
	expect(t, m, n, 1.0, 2)

	m, n = sci_norm(100000.0)
	expect(t, m, n, 1.0, 5)

	// < 1

	m, n = sci_norm(0.1)
	expect(t, m, n, 1.0, -1)

	m, n = sci_norm(0.01)
	expect(t, m, n, 1.0, -2)

	m, n = sci_norm(0.00001)
	expect(t, m, n, 1.0, -5)

	// different values

	m, n = sci_norm(0.0000347)
	expect(t, m, n, 3.47, -5)

	m, n = sci_norm(4700.0)
	expect(t, m, n, 4.7, 3)
}

func expectSer(t *testing.T, l_m float64, l_n int, u_m float64, u_n int, e error,
	l_m_e float64, l_n_e int, u_m_e float64, u_n_e int, e_e bool) {
	if e != nil && e_e == false {
		t.Errorf("Got unexpected error: %s\n", e)
	} else if e == nil && e_e == true {
		t.Errorf("Error expected but got nil\n")
	}

	if !Zero(l_m - l_m_e) {
		t.Errorf("l_m != l_m_expected (%f != %f)\n", l_m, l_m_e)
	}

	if !Zero(u_m - u_m_e) {
		t.Errorf("u_m != u_m_expected (%f != %f)\n", u_m, u_m_e)
	}

	if l_n != l_n_e {
		t.Errorf("l_n != l_n_expected (%d != %d)\n", l_n, l_n_e)
	}

	if u_n != u_n_e {
		t.Errorf("u_n != u_n_expected (%d != %d)\n", u_n, u_n_e)
	}
}

func TestEseriesStruct(t *testing.T) {
	s := ESeries{name: "E12", series: e12, tolerance_perc: 10.0}

	l_m, l_n, u_m, u_n, e := s.Calc(1.0)
	expectSer(t, l_m, l_n, u_m, u_n, e, 1.0, 0, 1.0, 0, false)

	l_m, l_n, u_m, u_n, e = s.Calc(1.1)
	expectSer(t, l_m, l_n, u_m, u_n, e, 1.0, 0, 1.2, 0, false)

}
