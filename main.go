package main

import (
	"fmt"
	_ "math"
	_ "unsafe"
)

func GetNumber() int {
	return 100
}

// const num = GetNumber()

func main() {
	// 可以看到，我们的解决方案是一种近似判断，通过一个可以接受的最小误差值 p，约定如果两个浮点数的差值在此精度的误差范围之内，则判定这两个浮点数相等。这个解决方案也是其他语言判断浮点数相等所采用的通用方案，PHP 也是这么做的。
	// float_value_1 := 0.1
	// float_value_2 := 0.1
	// p := 0.00001
	// // 判断 float_vlalue_1 与 float_value_2 是否相等
	// if math.Dim(float64(float_value_1), float_value_2) < p {
	// 	fmt.Println("float_value_1 和 float_value_2 相等")
	// }

	// slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// slice4 := append(slice3[:0], slice3[3:]...) // 删除开头三个元素

	// fmt.Println("slice3 value:", slice3)
	// fmt.Println("slice4 value:", slice4)
	// slice5 := append(slice3[:1], slice3[4:]...) // 删除中间三个元素
	// fmt.Println("slice4 value:", slice4)

	// slice6 := append(slice3[:0], slice3[:7]...) // 删除最后三个元素

	// slice7 := slice3[:copy(slice3, slice3[3:])] // 删除开头前三个元素
	// fmt.Println("slice3 value:", slice3)
	// fmt.Println("slice4 value:", slice4)
	// fmt.Println("slice5 value:", slice5)
	// fmt.Println("slice6 value:", slice6)
	// fmt.Println("slice7 value:", slice7)

	// 元素交换
	/*a := []int{1, 2, 3, 4, 5, 6}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)*/
	algorithm.GetFibonacci()
}
