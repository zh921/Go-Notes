package main

import (
	"fmt"
	"time"
)

//func main() {
//	var a [10]int
//	for i := 0; i < 10; i++ {
//		go func(ii int) {
//			for {
//				a[ii]++
//				runtime.Gosched() // 手动交出协程的控制权
//			}
//		}(i) // 取外层的变量i传输函数参数ii
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}