package main

import (
	"Go-test/src/simplemath"
	"fmt"
	"os"
	"strconv"
	// "google.golang.org/grpc"
)

// 定义一个用于打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

// 程序入口函数
func main() {
	args := os.Args

	// 除程序名本身外，至少需要传入两个其他参数，否则退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}

		ret := simplemath.Add(v1, v2)

		fmt.Println("Result: ", ret)

	case "sqrt":
		// 至少需要包含三个参数
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 获取待计算平方根的数值，并将类型转化为整型
		v, err := strconv.Atoi(args[2])
		// 获取参数出错，则退出
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		// 从 simplemath 包引入 Sqrt 方法进行平方根计算
		ret := simplemath.Sqrt(v)
		// 打印计算结果
		fmt.Println("Result: ", ret)
		// 如果计算方法不支持，打印程序使用指南
	default:
		Usage()
	}

	// grpc.NewServer()
}
