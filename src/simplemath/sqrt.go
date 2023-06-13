package simplemath

import "math"

// 计算平方根
func Sqrt(param1 int) int {
	ret := math.Sqrt(float64(param1))
	return int(ret)
}
