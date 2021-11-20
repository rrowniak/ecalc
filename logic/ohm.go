package logic

import (
	"errors"
	"fmt"
	"os"
)

func Ohm_help() {
	fmt.Println("Ohm's law calculator.")
	fmt.Println("Uasge: " + os.Args[0] + " ohm val1 val2")
	fmt.Println("Example: " + os.Args[0] + " ohm 10.6V 60mA")
	fmt.Println("Example: " + os.Args[0] + " ohm 2k5 81.60uA")
	fmt.Println("Example: " + os.Args[0] + " ohm 200R 5V")
}

func Ohm_exec(a, b string) {
	args := []string{a, b}

	result, w, e := calc(args)

	if e == nil {
		print_result(result, w)
	} else {
		fmt.Println(e)
	}
}

func calc(args []string) (Val, Val, error) {
	var I, V, R, W Val
	W.u = U_W
	var i, r bool

	for _, arg := range args {
		val, e := ParseQuantity(arg)

		if e != nil {
			return I, W, e
		}

		switch val.u {
		case U_A:
			I = val
			i = true
		case U_V:
			V = val
		case U_Ohm:
			R = val
			r = true
		case U_undefined:
			R = val
			r = true
		}
	}

	if !r {
		if Zero(I.v) {
			return I, W, errors.New("current cannot be zero")
		}
		R.v = V.v / I.v
		R.u = U_Ohm
		W.v = V.v * I.v
		return R, W, nil
	} else if !i {
		if Zero(R.v) {
			return I, W, errors.New("resistance cannot be zero")
		}
		I.v = V.v / R.v
		I.u = U_A
		W.v = V.v * I.v
		return I, W, nil
	} else {
		V.v = I.v * R.v
		V.u = U_V
		W.v = V.v * I.v
		return V, W, nil
	}
}

func print_result(v, w Val) {
	fmt.Printf("%s (%f %s), power %s\n", v.ToString(),
		v.v, v.u.ToString(), w.ToString())
}
