package main

/**
多值返回
函数可以返回任意数量的返回值。
swap 函数返回了两个字符串。

注意：此程序 360 报毒。。。。。

*/

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "go")
	fmt.Println(a, b)
}
