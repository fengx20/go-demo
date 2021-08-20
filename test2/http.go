// 同一个目录下面不能有个多 package main
// 标记当前文件为 main 包，main 包也是 Go 程序的入口包
package main

// 导入 net/http 包，这个包的作用是 HTTP 的基础封装和访问
import (
	"net/http"
)

/**
HTTP文件服务器
*/
// 程序执行的入口函数 main()
func main() {

	// 使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 默认的 HTTP 服务侦听在本机 8080 端口
	_ = http.ListenAndServe(":8080", nil)
}

// 在浏览器里输入http://127.0.0.1:8080即可浏览文件，这些文件正是当前目录在HTTP服务器上的映射目录。
