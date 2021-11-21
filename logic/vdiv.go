package logic

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func Vdiv_help() {
	fmt.Println("A simple voltage divider calculator.")
	fmt.Println()
	fmt.Println("usage: " + os.Args[0] + " <ARGs>...")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("\t-vin=val\tInput voltage")
	fmt.Println("\t-vout=val\tOutput voltage")
	fmt.Println("\t-r1=val\t\tRezistor R1 (Z1)")
	fmt.Println("\t-r2=val\t\tRezistor R2 (Z2)")
	fmt.Println("For these three provided quantities the forth one will be calculated.")
	fmt.Println()
	fmt.Println("Example: " + os.Args[0] + " vdiv -vin=10.0V -vout=1V -r1=1kR")
}

func Vdiv_exec(args []string) {
	if len(args) != 3 {
		fmt.Printf("Expected 3 arguments. Provided %d\n", len(args))
		os.Exit(1)
	}

	// parse arguments
	vdivCmd := flag.NewFlagSet("vdiv", flag.ExitOnError)

	vin_s := vdivCmd.String("vin", "", "input voltage")
	vout_s := vdivCmd.String("vout", "", "output voltage")
	r1_s := vdivCmd.String("r1", "", "Resistor R1")
	r2_s := vdivCmd.String("r2", "", "Resistor R2")

	vdivCmd.Parse(args)

	var vin, vout, r1, r2 Val

	// validate
	cnt := 0
	parse_quant(*vin_s, &vin, &cnt, U_V)
	parse_quant(*vout_s, &vout, &cnt, U_V)
	parse_quant(*r1_s, &r1, &cnt, U_Ohm)
	parse_quant(*r2_s, &r2, &cnt, U_Ohm)

	if cnt != 3 {
		fmt.Printf("Expected three quantities. Provided %d\n", cnt)
		os.Exit(1)
	}
	var e error
	if len(*vin_s) == 0 {
		vin, e = calc_vin(vout, r1, r2)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		print_results("vin", vin)
	} else if len(*vout_s) == 0 {
		vout, e = calc_vout(vin, r1, r2)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		print_results("vout", vout)
	} else if len(*r1_s) == 0 {
		r1, e = calc_r1(vin, vout, r2)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		print_results("r1", r1)
	} else if len(*r2_s) == 0 {
		r2, e = calc_r2(vin, vout, r1)
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}
		print_results("r2", r2)
	}
}

func parse_quant(s string, v *Val, cnt *int, u UnitSymbol) {
	if len(s) == 0 {
		return
	}
	var e error
	*v, e = ParseQuantity(s)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	if v.u == U_undefined {
		v.u = u
	} else if v.u != u {
		fmt.Printf("Unexpected unit in '%s'. Expected: %s unit\n", s, u.ToString())
		os.Exit(1)
	}

	if v.u == U_Ohm && (v.v <= 0 || Zero(v.v)) {
		fmt.Printf("Zero or negative resistance '%s' does't make any sense\n", s)
		os.Exit(1)
	}

	*cnt++
}

func print_results(pref string, v Val) {
	fmt.Printf("%s = %s (%f %s)\n", pref, v.ToString(), v.v, v.u.ToString())
}

func calc_vin(vout, r1, r2 Val) (Val, error) {
	vin := Val{u: U_V}

	vin.v = vout.v * (r1.v + r2.v) / r2.v

	return vin, nil
}

func calc_vout(vin, r1, r2 Val) (Val, error) {
	vout := Val{u: U_V}

	vout.v = vin.v * r2.v / (r1.v + r2.v)

	return vout, nil
}

func calc_r1(vin, vout, r2 Val) (Val, error) {
	r1 := Val{u: U_Ohm}

	if Zero(vin.v - vout.v) {
		return r1, errors.New("the same input and output voltage doesn't make any sense")
	}

	r1.v = r2.v * (vin.v - vout.v) / vout.v

	return r1, nil
}

func calc_r2(vin, vout, r1 Val) (Val, error) {
	r2 := Val{u: U_Ohm}

	if Zero(vin.v - vout.v) {
		return r2, errors.New("the same input and output voltage doesn't make any sense")
	}

	r2.v = r1.v * vout.v / (vin.v - vout.v)

	return r2, nil
}
