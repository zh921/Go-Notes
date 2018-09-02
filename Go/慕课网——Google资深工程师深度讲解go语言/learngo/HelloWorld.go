package main

import "fmt"

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

func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDedection()
	variableShorter()
	fmt.Println(aa, ss, bb)
}
