package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkK"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDedection() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def" // 第一次定义变量用  :=
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // (0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // (0+1.2246467991473515e-16i)
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1) // (0.000+0.000i)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) // 强制类型转换
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp        = iota // 自增    0
		_                 // 跳过这个值
		python            //        2
		golang            //        3
		javascript        //        4
	)

	// kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

// func main() {
// 	fmt.Println("Hello world!")
// 	variableZeroValue()
// 	variableInitialValue()
// 	variableTypeDedection()
// 	variableShorter()
// 	fmt.Println(aa, ss, bb)
// 	euler()
// 	triangle()
// 	consts()
// 	enums()
// }
