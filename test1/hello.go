package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Println("你好")
		time.Sleep(time.Millisecond * 100)
	}
}
