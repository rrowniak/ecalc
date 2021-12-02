package logic

import "math"

var epsilon float64

func init() {
    epsilon = math.Nextafter(0.0, 1.0)
}

func Zero(v float64) bool {
    return math.Abs(v) <= epsilon
}
