package logic

import "testing"

func TestDbBasic(t *testing.T) {
    pow, ampl, e := db_calc([]string{"100000", "1"})

    if e != nil {
        t.Error("Unexpected error")
    }

    if pow.u != U_Db {
        t.Error("Unexpected unit")
    }

    if pow.v != 50.0 {
        t.Error("Unexpected value")
    }

    if ampl.u != U_Db {
        t.Error("Unexpected unit")
    }

    if ampl.v != 100.0 {
        t.Error("Unexpected value")
    }
}
