package problems

import (
	"math"
	"reflect"
)

const TOLERANCE = 1e-8

func MyPowTest() (bool, string) {
	var (
		x float64
		n int
	)

	x = 2.0000
	n = 10

	if !reflect.DeepEqual(myPow(x, n), 1024.00000) {
		return false, "case 1"
	}

	x = 2.10000
	n = 3

	if math.Abs(myPow(x, n)-9.26100) > TOLERANCE {
		return false, "case 2"
	}

	x = 2.0000
	n = -2
	if !reflect.DeepEqual(myPow(x, n), 0.25) {
		return false, "case 3"
	}

	return true, "good"
}
