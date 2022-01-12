package main

/**
函数
函数可以没有参数或接受多个参数。
在本例中，add 接受两个 int 类型的参数。
注意类型在变量名 之后。
*/

import "fmt"

func add1(x int, y int) int {
	return x + y
}

/**
当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。
在本例中，
x int, y int
被缩写为
x, y int
*/

func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add2(1, 4))
}
