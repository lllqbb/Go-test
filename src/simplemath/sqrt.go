package simplemath

import "math"

func Sqrt(param1 int) int {
	ret := math.Sqrt(float64(param1))
	return int(ret)
}
