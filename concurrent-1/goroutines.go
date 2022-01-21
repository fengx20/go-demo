package main

/**
Go 程
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
go f(x, y, z)
会启动一个新的 Go 程并执行
f(x, y, z)
f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。
Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
*/
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
