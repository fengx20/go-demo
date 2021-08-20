package main

// 引入其他包
import (
	"fmt"
	"go-demo/simplemath"
	"os"
	"strconv"
)

// 定义一个用于打印程序使用指南的函数
var Usage = func() {
	fmt.Println("USAGE：calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\t计算两个数值相加\n\tsqrt\t计算一个非负数的平方根")
}

// 程序入口函数
func main() {
	/*
	 * 用于获取命令行参数，注意程序名本身是第一个参数，
	 * 比如 calc add 1 2 这条指令，第一个参数是 calc
	 */
	// := 声明并初始化变量，不需要通过 var 声明该变量
	args := os.Args
	// 除程序名本身外，至少需要传入两个其他参数，否则退出
	if args == nil || len(args) < 3 {
		Usage()
		return
	}
	// 第二个参数表示计算方法
	switch args[1] {
	// 加法
	case "add":
		// 至少需要4个参数
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 获取待相加的类型，并将类型转换为整型
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		// 获取参数出错，则退出
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 从 simplemath 包引入 Add 方法
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	// 求方根
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)

	default:
		Usage()
	}
}
