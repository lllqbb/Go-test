package algorithm

import (
	"fmt"
	"time"
)

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		fi := a
		a, b = b, a+b // 1,1
		// 1,2
		//2,3
		return fi
	}
}

func GetFibonacci() {
	s := time.Now()
	f := fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Println(f())
	}
	e := time.Since(s)
	fmt.Println(e)
}
