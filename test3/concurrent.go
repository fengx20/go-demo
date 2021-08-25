// 同一个目录下面不能有个多 package main
package main

// 导入格式化（fmt）、随机数（math/rand）、时间（time）包参与编译
import (
	"fmt"
	"math/rand"
	"time"
)

// 生产数据的函数，传入一个标记类型的字符串及一个只能写入的通道
// 数据生产者
func producer(header string, channel chan<- string) {
	// for{} 构成一个无限循环
	// 无限循环，不停地生产数据
	for {
		// 使用 rand.Int31() 生成一个随机数，使用 fmt.Sprintf() 函数将 header 和随机数格式化为字符串
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 使用 time.Sleep() 函数暂停 1 秒再执行这个函数。如果在 goroutine 中执行时，暂停不会影响其他 goroutine 的执行
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 消费数据的函数，传入一个只能写入的通道
// 数据消费者
func customer(channel <-chan string) {
	// 不停的获取数据
	for {
		// 从通道中获取数据，此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}

// 程序的入口函数，总是在程序开始时执行
func main() {

	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 并发执行一个生产者函数，两行分别创建了这个函数搭配不同参数的两个 goroutine
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	// 执行消费者函数通过通道进行数据消费
	// 数据消费函数
	customer(channel)
}

/**
整段代码中，没有线程创建，没有线程池也没有加锁，仅仅通过关键字 go 实现 goroutine，和通道实现数据交换。

编译成exe文件：go build go-test/test3/concurrent-1.go
*/
